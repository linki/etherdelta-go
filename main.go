package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/alecthomas/kingpin"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/linki/etherdelta-go/contract"
)

const (
	version                = "v0.2.0"
	defaultEndpoint        = "https://mainnet.infura.io"
	defaultContractAddress = "0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"
	defaultGasPrice        = "1000000000"
	defaultGasLimit        = "100000"
)

var (
	endpoint        string
	contractAddress string
	tokenAddress    string
	keyStorePath    string
	passphrase      string
	skipWithdraw    bool
	gasPrice        int64
	gasLimit        int64
)

func init() {
	kingpin.Flag("endpoint", "Ethereum RPC endpoint (optional, default: https://mainnet.infura.io)").Default(defaultEndpoint).StringVar(&endpoint)
	kingpin.Flag("contract", "EtherDelta contract address (optional, default: 0x8d12A197cB00D4747a1fe03395095ce2A5CC6819)").Default(defaultContractAddress).StringVar(&contractAddress)
	kingpin.Flag("token", "The token's contract address (required)").Required().StringVar(&tokenAddress)
	kingpin.Flag("keystore-file", "The owner's Keystore file location (required)").Required().ExistingFileVar(&keyStorePath)
	kingpin.Flag("passphrase", "The Keystore file's passphrase (required)").Required().StringVar(&passphrase)
	kingpin.Flag("skip-withdraw", "Don't withdraw remaining tokens (optional, default: false)").BoolVar(&skipWithdraw)
	kingpin.Flag("gas-price", "The gas price in wei (optional, default: 1000000000, i.e. 1 Gwei)").Default(defaultGasPrice).Int64Var(&gasPrice)
	kingpin.Flag("gas-limit", "The gas limit; default: 100000").Default(defaultGasLimit).Int64Var(&gasLimit)
}

func main() {
	// Parse command line flags.
	kingpin.Version(version)
	kingpin.Parse()

	// Create an Ethereum client connecting to the provided RPC endpoint.
	c, err := rpc.DialHTTP(endpoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}
	conn := ethclient.NewClient(c)

	// Create an instance of the EtherDelta contract at the specific address.
	etherdelta, err := contract.NewEtherDelta(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate EtherDelta contract instance: %v", err)
	}

	// Open the provided Keystore file.
	keyStoreFile, err := os.OpenFile(keyStorePath, os.O_RDONLY, 0400)
	if err != nil {
		log.Fatalf("Failed to open Keystore file: %v", err)
	}
	defer keyStoreFile.Close()

	// Read the Keystore file's JSON content.
	keyStoreJSON, err := ioutil.ReadAll(keyStoreFile)
	if err != nil {
		log.Fatalf("Failed to parse Keystore file: %v", err)
	}

	// Parse and decrypt the content given a passphrase.
	ownerKey, err := keystore.DecryptKey(keyStoreJSON, passphrase)
	if err != nil {
		log.Fatalf("Failed to decrypt Keystore file: %v", err)
	}

	// Get the owner's as well as the token's addresses.
	owner := ownerKey.Address
	token := common.HexToAddress(tokenAddress)

	// Retrieve the owner's deposited balance of the token.
	balance, err := etherdelta.BalanceOf(nil, token, owner)
	if err != nil {
		log.Fatalf("Failed to retrieve balance: %v", err)
	}
	fmt.Println("Deposited tokens:", balance)

	// If withdrawal is disabled or there's nothing to withdraw, end the program.
	if skipWithdraw || balance.Cmp(common.Big0) == 0 {
		os.Exit(0)
	}

	// Otherwise withdraw the deposited tokens.
	fmt.Println("Withdrawing tokens:", balance)

	// The transaction needs additional metadata, e.g. the private key for
	// authorization as well as gas price and limit.
	opts := bind.NewKeyedTransactor(ownerKey.PrivateKey)
	opts.GasLimit = big.NewInt(gasLimit)
	opts.GasPrice = big.NewInt(gasPrice)

	// Execute the WithdrawToken function of the EtherDelta contract given the
	// required arguments: the token address and amount. Also provide the metadata
	// created above in order to successfully create and sign the transaction.
	tx, err := etherdelta.WithdrawToken(opts, token, balance)
	if err != nil {
		log.Fatalf("Failed to withdraw tokens: %v", err)
	}

	// Print the hash of the submitted transaction for tracking.
	fmt.Println("Transaction hash:", tx.Hash().String())
}
