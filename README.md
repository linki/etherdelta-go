# etherdelta-go

This is an example of how to interact with an Ethereum contract from Go through RPC.

The example shows three things:
* how to generate Go language bindings given a contract's ABI definition,
* how to read data from the blockchain without spending gas, and
* how to change state on the blockchain by providing the correct private key and spending gas

# The contract library

We use [abigen](https://github.com/ethereum/go-ethereum/tree/v1.7.2/cmd/abigen)
to generate Go language bindings given a contract's ABI definition.

To regenerate the code [install abigen](https://github.com/ethereum/go-ethereum/tree/v1.7.2#building-the-source), put the ABI definition into a file (e.g., [for the EtherDelta contract](https://etherscan.io/address/0x8d12a197cb00d4747a1fe03395095ce2a5cc6819#code)) and run:

```console
$ cd contract
$ abigen --abi etherdelta.abi --pkg contract --type EtherDelta --out etherdelta.go
```

This will create an [etherdelta.go](contract/etherdelta.go) file that you can import into your Go programs.

# Example program

The [example code](main.go) contains a working interaction with the Ethereum blockchain and may even be useful for somebody.

When using [EtherDelta](https://etherdelta.com/) (a decentralized exchange based on smart contracts) a user deposits funds from her own address into a pool that can be accessed by EtherDelta for executing trades. Later the user withdraws any funds from the pool back to her own address. Often there's a tiny amount left in the pool which is tedious to withdraw manually.

`etherdelta-go` to the rescue; the following command will withdraw any remaining deposited funds back to the owner for a given token.

Build and install the binary:

```console
$ cd $GOPATH/src/github.com/linki/etherdelta-go
$ glide install --strip-vendor
$ go install
```

Make sure `$GOPATH/bin` is in your `$PATH`, then run it like that:

```console
$ etherdelta-go --token "0x27d...488" --keystore-file "UTC-...98c" --passphrase "my...pass"
Deposited tokens: 500000000
Withdrawing tokens: 500000000
Transaction hash: 0x1a6328...7ef38a

$ # Wait for transaction to be mined, then run it again.

$ etherdelta-go --token "0x27d...488" --keystore-file "UTC-...98c" --passphrase "my...pass"
Deposited tokens: 0
```

It takes the following arguments:
* `token`: the address of the [ERC20](https://theethereum.wiki/w/index.php/ERC20_Token_Standard) [Token](https://etherscan.io/tokens) you want to withdraw, e.g., [AirToken](https://etherscan.io/token/0x27dce1ec4d3f72c3e457cc50354f1f975ddef488)
* `keystore-file`: the location of an Ethereum [Keystore file](https://theethereum.wiki/w/index.php/Accounts,_Addresses,_Public_And_Private_Keys,_And_Tokens#UTC_JSON_Keystore_File) which describes where the funds will be withdrawn to as well as contains the encrypted private key for that address, e.g., `~/Library/Ethereum/keystore/UTC--2017-11-...61d3f9`
* `passphrase`: the passphrase unlocking the `keystore-file`. (This proves that you are the owner of the Keystore file)

If the deposited token balance is zero then no transaction will be issued.

There are some optional arguments as well:
* `endpoint`: an RPC endpoint to interact with the Ethereum blockchain. You can run your own Ethereum node or connect to one provided by [Infura](https://infura.io/) free of charge. Defaults to `https://mainnet.infura.io`.
* `contract`: The address of the EtherDelta contract to use. Since contracts are immutable each change to the contract requires a new contract to be deployed and results in a new contract address. Defaults to the most recent contract as of November 2017: [0x8d12A197cB00D4747a1fe03395095ce2A5CC6819](https://etherscan.io/address/0x8d12A197cB00D4747a1fe03395095ce2A5CC6819).
* `skip-withdraw`: doesn't send any withdraw transaction but merely shows the token balance. Note that `passphrase` must still be provided and correct in order to successfully process the Keystore file. Defaults to `false` (enable withdrawal).
* `gas-price`: The gas price in wei for the withdrawal transaction. Defaults to `1000000000` which is 1 Gwei.
* `gas-limit`: The maximum gas to use for the withdrawal transaction. Defaults to `100000`. During my testing for AirToken the final gas for successful transactions hovered around `50000`.

Keeping the defaults for `gas-price` and `gas-limit` will result in a maximum transaction fee of `0.0001` ETH.

# Next steps

It would be nice to extend the tool to iterate over all tokens that have some deposited funds left in EtherDelta's pool and send it back to the owner.

# Disclaimer

No warranty for any funds lost. See [LICENSE](LICENSE) for details.
