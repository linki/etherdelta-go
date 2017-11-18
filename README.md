# etherdelta-go

This is an example how to interact with Ethereum contracts from Go through RPC.

It uses [abigen](https://github.com/ethereum/go-ethereum/tree/v1.7.2/cmd/abigen)
to generate Go language bindings given a contract's ABI definition.

To regenerate the code [install abigen](https://github.com/ethereum/go-ethereum/tree/v1.7.2#building-the-source), put the ABI definition into a file (e.g., [for the EtherDelta contract](https://etherscan.io/address/0x8d12a197cb00d4747a1fe03395095ce2a5cc6819#code)) and run:

```console
$ cd contract
$ abigen --abi etherdelta.abi --pkg contract --type EtherDelta --out etherdelta.go
```

This will create an [etherdelta.go](contract/etherdelta.go) file that you can import into your Go programs.

The [example code](main.go) contains a working interaction with the Ethereum blockchain and may even be useful for somebody.

```
$ ./etherdelta-go --endpoint "https://mainnet.infura.io" --token "0x27d...488" --owner "0x585...828C" --private-key "d4c...7b3"
```

When using [EtherDelta](https://etherdelta.com/) (a decentralized exchange based on smart contracts) a user deposits funds from her own address into a pool that can be accessed by EtherDelta for executing trades. Later the user withdraws any funds from the pool back to her own address. Often there's a tiny amount left in the pool which is tedious to withdraw manually.

`etherdelta-go` to the rescue; the above command will withdraw any remaining deposited funds back to the owner for a given token. It takes the following arguments:
* an RPC endpoint to interact with the Ethereum blockchain (you can run your own Ethereum node or connect to one provided by [Infura](https://infura.io/) for no charge)
* the address of the token you want to withdraw, e.g., [AirToken](https://etherscan.io/token/0x27dce1ec4d3f72c3e457cc50354f1f975ddef488)
* the address of the owner of the tokens and where the funds will be withdrawn to (this should be one of yours)
* the private key corresponding to the owner's address in order to sign the withdraw transaction (this proves that you are the owner of the above address)

The example code shows two things:

* how to read data from the blockchain without spending gas (e.g., printing the token balance for a particular address)
* how to change state in the blockchain by providing the correct private key and spending gas

The example code uses one Gwei for the gas cost and 100,000 for the gas limit but you can configure that via the `--gas-price` (in Gwei) and `--gas-limit` flags.

If the deposited token balance is zero then no transaction will be issued.

# Next steps

It would be nice to extend the tool to iterate over all tokens that have some deposited funds left in EtherDelta's pool and send it back to the owner.

# Disclaimer
