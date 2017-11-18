package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/linki/etherdelta-go/contract"
)

const (
	version                = "v0.1.0"
	defaultContractAddress = "0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"
	defaultGasPrice        = "1"
	defaultGasLimit        = "100000"
)

var (
	endpoint        string
	contractAddress string
	tokenAddress    string
	ownerAddress    string
	ownerPrivateKey string
	gasPrice        int64
	gasLimit        int64
)

func init() {
	kingpin.Flag("endpoint", "Ethereum RPC endpoint").Required().StringVar(&endpoint)
	kingpin.Flag("contract", "EtherDelta contract address").Default(defaultContractAddress).StringVar(&contractAddress)
	kingpin.Flag("token", "The token's contract address").Required().StringVar(&tokenAddress)
	kingpin.Flag("owner", "The owner's address").Required().StringVar(&ownerAddress)
	kingpin.Flag("private-key", "The owner's address' private key").Required().StringVar(&ownerPrivateKey)
	kingpin.Flag("gas-price", "The gas price in Gwei").Default(defaultGasPrice).Int64Var(&gasPrice)
	kingpin.Flag("gas-limit", "The gas limit").Default(defaultGasLimit).Int64Var(&gasLimit)
}

func main() {
	kingpin.Version(version)
	kingpin.Parse()

	c, err := rpc.DialHTTP(endpoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}
	conn := ethclient.NewClient(c)

	etherdelta, err := contract.NewEtherDelta(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate EtherDelta contract instance: %v", err)
	}

	token := common.HexToAddress(tokenAddress)
	owner := common.HexToAddress(ownerAddress)

	balance, err := etherdelta.BalanceOf(nil, token, owner)
	if err != nil {
		log.Fatalf("Failed to retrieve balance: %v", err)
	}
	fmt.Println("Token balance:", balance)

	if balance.Cmp(common.Big0) == 0 {
		os.Exit(0)
	}

	fmt.Println("Withdrawing tokens:", balance)

	privateKey, err := crypto.ToECDSA(common.Hex2Bytes(ownerPrivateKey))
	if err != nil {
		log.Fatalf("Failed to read private key: %v", err)
	}

	opts := bind.NewKeyedTransactor(privateKey)
	opts.GasLimit = big.NewInt(gasLimit)
	opts.GasPrice = big.NewInt(gasPrice * 1000000000)

	tx, err := etherdelta.WithdrawToken(opts, token, balance)
	if err != nil {
		log.Fatalf("Failed to withdraw tokens: %v", err)
	}

	fmt.Println("Transaction hash:", tx.Hash().String())
}
