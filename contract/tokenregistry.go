// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// TokenRegistryABI is the input ABI used to generate the binding from.
const TokenRegistryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"removeToken\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getTokenAddressByName\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"getTokenAddressBySymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_swarmHash\",\"type\":\"bytes\"}],\"name\":\"setTokenSwarmHash\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getTokenMetaData\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"_ipfsHash\",\"type\":\"bytes\"},{\"name\":\"_swarmHash\",\"type\":\"bytes\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setTokenName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"ipfsHash\",\"type\":\"bytes\"},{\"name\":\"swarmHash\",\"type\":\"bytes\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getTokenByName\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTokenAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_ipfsHash\",\"type\":\"bytes\"}],\"name\":\"setTokenIpfsHash\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"getTokenBySymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"setTokenSymbol\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"ipfsHash\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"swarmHash\",\"type\":\"bytes\"}],\"name\":\"LogAddToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"ipfsHash\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"swarmHash\",\"type\":\"bytes\"}],\"name\":\"LogRemoveToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"LogTokenNameChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldSymbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"newSymbol\",\"type\":\"string\"}],\"name\":\"LogTokenSymbolChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldIpfsHash\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"newIpfsHash\",\"type\":\"bytes\"}],\"name\":\"LogTokenIpfsHashChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldSwarmHash\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"newSwarmHash\",\"type\":\"bytes\"}],\"name\":\"LogTokenSwarmHashChange\",\"type\":\"event\"}]"

// TokenRegistry is an auto generated Go binding around an Ethereum contract.
type TokenRegistry struct {
	TokenRegistryCaller     // Read-only binding to the contract
	TokenRegistryTransactor // Write-only binding to the contract
	TokenRegistryFilterer   // Log filterer for contract events
}

// TokenRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenRegistrySession struct {
	Contract     *TokenRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenRegistryCallerSession struct {
	Contract *TokenRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenRegistryTransactorSession struct {
	Contract     *TokenRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRegistryRaw struct {
	Contract *TokenRegistry // Generic contract binding to access the raw methods on
}

// TokenRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenRegistryCallerRaw struct {
	Contract *TokenRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenRegistryTransactorRaw struct {
	Contract *TokenRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenRegistry creates a new instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistry(address common.Address, backend bind.ContractBackend) (*TokenRegistry, error) {
	contract, err := bindTokenRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenRegistry{TokenRegistryCaller: TokenRegistryCaller{contract: contract}, TokenRegistryTransactor: TokenRegistryTransactor{contract: contract}, TokenRegistryFilterer: TokenRegistryFilterer{contract: contract}}, nil
}

// NewTokenRegistryCaller creates a new read-only instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryCaller(address common.Address, caller bind.ContractCaller) (*TokenRegistryCaller, error) {
	contract, err := bindTokenRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryCaller{contract: contract}, nil
}

// NewTokenRegistryTransactor creates a new write-only instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenRegistryTransactor, error) {
	contract, err := bindTokenRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryTransactor{contract: contract}, nil
}

// NewTokenRegistryFilterer creates a new log filterer instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenRegistryFilterer, error) {
	contract, err := bindTokenRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryFilterer{contract: contract}, nil
}

// bindTokenRegistry binds a generic wrapper to an already deployed contract.
func bindTokenRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRegistry *TokenRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenRegistry.Contract.TokenRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRegistry *TokenRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TokenRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRegistry *TokenRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TokenRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRegistry *TokenRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRegistry *TokenRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRegistry *TokenRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetTokenAddressByName is a free data retrieval call binding the contract method 0x2fbfeba9.
//
// Solidity: function getTokenAddressByName(_name string) constant returns(address)
func (_TokenRegistry *TokenRegistryCaller) GetTokenAddressByName(opts *bind.CallOpts, _name string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenRegistry.contract.Call(opts, out, "getTokenAddressByName", _name)
	return *ret0, err
}

// GetTokenAddressByName is a free data retrieval call binding the contract method 0x2fbfeba9.
//
// Solidity: function getTokenAddressByName(_name string) constant returns(address)
func (_TokenRegistry *TokenRegistrySession) GetTokenAddressByName(_name string) (common.Address, error) {
	return _TokenRegistry.Contract.GetTokenAddressByName(&_TokenRegistry.CallOpts, _name)
}

// GetTokenAddressByName is a free data retrieval call binding the contract method 0x2fbfeba9.
//
// Solidity: function getTokenAddressByName(_name string) constant returns(address)
func (_TokenRegistry *TokenRegistryCallerSession) GetTokenAddressByName(_name string) (common.Address, error) {
	return _TokenRegistry.Contract.GetTokenAddressByName(&_TokenRegistry.CallOpts, _name)
}

// GetTokenAddressBySymbol is a free data retrieval call binding the contract method 0x3550b6d9.
//
// Solidity: function getTokenAddressBySymbol(_symbol string) constant returns(address)
func (_TokenRegistry *TokenRegistryCaller) GetTokenAddressBySymbol(opts *bind.CallOpts, _symbol string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenRegistry.contract.Call(opts, out, "getTokenAddressBySymbol", _symbol)
	return *ret0, err
}

// GetTokenAddressBySymbol is a free data retrieval call binding the contract method 0x3550b6d9.
//
// Solidity: function getTokenAddressBySymbol(_symbol string) constant returns(address)
func (_TokenRegistry *TokenRegistrySession) GetTokenAddressBySymbol(_symbol string) (common.Address, error) {
	return _TokenRegistry.Contract.GetTokenAddressBySymbol(&_TokenRegistry.CallOpts, _symbol)
}

// GetTokenAddressBySymbol is a free data retrieval call binding the contract method 0x3550b6d9.
//
// Solidity: function getTokenAddressBySymbol(_symbol string) constant returns(address)
func (_TokenRegistry *TokenRegistryCallerSession) GetTokenAddressBySymbol(_symbol string) (common.Address, error) {
	return _TokenRegistry.Contract.GetTokenAddressBySymbol(&_TokenRegistry.CallOpts, _symbol)
}

// GetTokenAddresses is a free data retrieval call binding the contract method 0xee8c24b8.
//
// Solidity: function getTokenAddresses() constant returns(address[])
func (_TokenRegistry *TokenRegistryCaller) GetTokenAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TokenRegistry.contract.Call(opts, out, "getTokenAddresses")
	return *ret0, err
}

// GetTokenAddresses is a free data retrieval call binding the contract method 0xee8c24b8.
//
// Solidity: function getTokenAddresses() constant returns(address[])
func (_TokenRegistry *TokenRegistrySession) GetTokenAddresses() ([]common.Address, error) {
	return _TokenRegistry.Contract.GetTokenAddresses(&_TokenRegistry.CallOpts)
}

// GetTokenAddresses is a free data retrieval call binding the contract method 0xee8c24b8.
//
// Solidity: function getTokenAddresses() constant returns(address[])
func (_TokenRegistry *TokenRegistryCallerSession) GetTokenAddresses() ([]common.Address, error) {
	return _TokenRegistry.Contract.GetTokenAddresses(&_TokenRegistry.CallOpts)
}

// GetTokenByName is a free data retrieval call binding the contract method 0xe73fc0c3.
//
// Solidity: function getTokenByName(_name string) constant returns(address, string, string, uint8, bytes, bytes)
func (_TokenRegistry *TokenRegistryCaller) GetTokenByName(opts *bind.CallOpts, _name string) (common.Address, string, string, uint8, []byte, []byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new(uint8)
		ret4 = new([]byte)
		ret5 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _TokenRegistry.contract.Call(opts, out, "getTokenByName", _name)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetTokenByName is a free data retrieval call binding the contract method 0xe73fc0c3.
//
// Solidity: function getTokenByName(_name string) constant returns(address, string, string, uint8, bytes, bytes)
func (_TokenRegistry *TokenRegistrySession) GetTokenByName(_name string) (common.Address, string, string, uint8, []byte, []byte, error) {
	return _TokenRegistry.Contract.GetTokenByName(&_TokenRegistry.CallOpts, _name)
}

// GetTokenByName is a free data retrieval call binding the contract method 0xe73fc0c3.
//
// Solidity: function getTokenByName(_name string) constant returns(address, string, string, uint8, bytes, bytes)
func (_TokenRegistry *TokenRegistryCallerSession) GetTokenByName(_name string) (common.Address, string, string, uint8, []byte, []byte, error) {
	return _TokenRegistry.Contract.GetTokenByName(&_TokenRegistry.CallOpts, _name)
}

// GetTokenBySymbol is a free data retrieval call binding the contract method 0xefa74f1f.
//
// Solidity: function getTokenBySymbol(_symbol string) constant returns(address, string, string, uint8, bytes, bytes)
func (_TokenRegistry *TokenRegistryCaller) GetTokenBySymbol(opts *bind.CallOpts, _symbol string) (common.Address, string, string, uint8, []byte, []byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new(uint8)
		ret4 = new([]byte)
		ret5 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _TokenRegistry.contract.Call(opts, out, "getTokenBySymbol", _symbol)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetTokenBySymbol is a free data retrieval call binding the contract method 0xefa74f1f.
//
// Solidity: function getTokenBySymbol(_symbol string) constant returns(address, string, string, uint8, bytes, bytes)
func (_TokenRegistry *TokenRegistrySession) GetTokenBySymbol(_symbol string) (common.Address, string, string, uint8, []byte, []byte, error) {
	return _TokenRegistry.Contract.GetTokenBySymbol(&_TokenRegistry.CallOpts, _symbol)
}

// GetTokenBySymbol is a free data retrieval call binding the contract method 0xefa74f1f.
//
// Solidity: function getTokenBySymbol(_symbol string) constant returns(address, string, string, uint8, bytes, bytes)
func (_TokenRegistry *TokenRegistryCallerSession) GetTokenBySymbol(_symbol string) (common.Address, string, string, uint8, []byte, []byte, error) {
	return _TokenRegistry.Contract.GetTokenBySymbol(&_TokenRegistry.CallOpts, _symbol)
}

// GetTokenMetaData is a free data retrieval call binding the contract method 0x7abccac9.
//
// Solidity: function getTokenMetaData(_token address) constant returns(address, string, string, uint8, bytes, bytes)
func (_TokenRegistry *TokenRegistryCaller) GetTokenMetaData(opts *bind.CallOpts, _token common.Address) (common.Address, string, string, uint8, []byte, []byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new(uint8)
		ret4 = new([]byte)
		ret5 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _TokenRegistry.contract.Call(opts, out, "getTokenMetaData", _token)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetTokenMetaData is a free data retrieval call binding the contract method 0x7abccac9.
//
// Solidity: function getTokenMetaData(_token address) constant returns(address, string, string, uint8, bytes, bytes)
func (_TokenRegistry *TokenRegistrySession) GetTokenMetaData(_token common.Address) (common.Address, string, string, uint8, []byte, []byte, error) {
	return _TokenRegistry.Contract.GetTokenMetaData(&_TokenRegistry.CallOpts, _token)
}

// GetTokenMetaData is a free data retrieval call binding the contract method 0x7abccac9.
//
// Solidity: function getTokenMetaData(_token address) constant returns(address, string, string, uint8, bytes, bytes)
func (_TokenRegistry *TokenRegistryCallerSession) GetTokenMetaData(_token common.Address) (common.Address, string, string, uint8, []byte, []byte, error) {
	return _TokenRegistry.Contract.GetTokenMetaData(&_TokenRegistry.CallOpts, _token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenRegistry *TokenRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenRegistry *TokenRegistrySession) Owner() (common.Address, error) {
	return _TokenRegistry.Contract.Owner(&_TokenRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenRegistry *TokenRegistryCallerSession) Owner() (common.Address, error) {
	return _TokenRegistry.Contract.Owner(&_TokenRegistry.CallOpts)
}

// TokenAddresses is a free data retrieval call binding the contract method 0xe5df8b84.
//
// Solidity: function tokenAddresses( uint256) constant returns(address)
func (_TokenRegistry *TokenRegistryCaller) TokenAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenRegistry.contract.Call(opts, out, "tokenAddresses", arg0)
	return *ret0, err
}

// TokenAddresses is a free data retrieval call binding the contract method 0xe5df8b84.
//
// Solidity: function tokenAddresses( uint256) constant returns(address)
func (_TokenRegistry *TokenRegistrySession) TokenAddresses(arg0 *big.Int) (common.Address, error) {
	return _TokenRegistry.Contract.TokenAddresses(&_TokenRegistry.CallOpts, arg0)
}

// TokenAddresses is a free data retrieval call binding the contract method 0xe5df8b84.
//
// Solidity: function tokenAddresses( uint256) constant returns(address)
func (_TokenRegistry *TokenRegistryCallerSession) TokenAddresses(arg0 *big.Int) (common.Address, error) {
	return _TokenRegistry.Contract.TokenAddresses(&_TokenRegistry.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens( address) constant returns(token address, name string, symbol string, decimals uint8, ipfsHash bytes, swarmHash bytes)
func (_TokenRegistry *TokenRegistryCaller) Tokens(opts *bind.CallOpts, arg0 common.Address) (struct {
	Token     common.Address
	Name      string
	Symbol    string
	Decimals  uint8
	IpfsHash  []byte
	SwarmHash []byte
}, error) {
	ret := new(struct {
		Token     common.Address
		Name      string
		Symbol    string
		Decimals  uint8
		IpfsHash  []byte
		SwarmHash []byte
	})
	out := ret
	err := _TokenRegistry.contract.Call(opts, out, "tokens", arg0)
	return *ret, err
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens( address) constant returns(token address, name string, symbol string, decimals uint8, ipfsHash bytes, swarmHash bytes)
func (_TokenRegistry *TokenRegistrySession) Tokens(arg0 common.Address) (struct {
	Token     common.Address
	Name      string
	Symbol    string
	Decimals  uint8
	IpfsHash  []byte
	SwarmHash []byte
}, error) {
	return _TokenRegistry.Contract.Tokens(&_TokenRegistry.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens( address) constant returns(token address, name string, symbol string, decimals uint8, ipfsHash bytes, swarmHash bytes)
func (_TokenRegistry *TokenRegistryCallerSession) Tokens(arg0 common.Address) (struct {
	Token     common.Address
	Name      string
	Symbol    string
	Decimals  uint8
	IpfsHash  []byte
	SwarmHash []byte
}, error) {
	return _TokenRegistry.Contract.Tokens(&_TokenRegistry.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0xa880319d.
//
// Solidity: function addToken(_token address, _name string, _symbol string, _decimals uint8, _ipfsHash bytes, _swarmHash bytes) returns()
func (_TokenRegistry *TokenRegistryTransactor) AddToken(opts *bind.TransactOpts, _token common.Address, _name string, _symbol string, _decimals uint8, _ipfsHash []byte, _swarmHash []byte) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "addToken", _token, _name, _symbol, _decimals, _ipfsHash, _swarmHash)
}

// AddToken is a paid mutator transaction binding the contract method 0xa880319d.
//
// Solidity: function addToken(_token address, _name string, _symbol string, _decimals uint8, _ipfsHash bytes, _swarmHash bytes) returns()
func (_TokenRegistry *TokenRegistrySession) AddToken(_token common.Address, _name string, _symbol string, _decimals uint8, _ipfsHash []byte, _swarmHash []byte) (*types.Transaction, error) {
	return _TokenRegistry.Contract.AddToken(&_TokenRegistry.TransactOpts, _token, _name, _symbol, _decimals, _ipfsHash, _swarmHash)
}

// AddToken is a paid mutator transaction binding the contract method 0xa880319d.
//
// Solidity: function addToken(_token address, _name string, _symbol string, _decimals uint8, _ipfsHash bytes, _swarmHash bytes) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) AddToken(_token common.Address, _name string, _symbol string, _decimals uint8, _ipfsHash []byte, _swarmHash []byte) (*types.Transaction, error) {
	return _TokenRegistry.Contract.AddToken(&_TokenRegistry.TransactOpts, _token, _name, _symbol, _decimals, _ipfsHash, _swarmHash)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x13baf1e6.
//
// Solidity: function removeToken(_token address, _index uint256) returns()
func (_TokenRegistry *TokenRegistryTransactor) RemoveToken(opts *bind.TransactOpts, _token common.Address, _index *big.Int) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "removeToken", _token, _index)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x13baf1e6.
//
// Solidity: function removeToken(_token address, _index uint256) returns()
func (_TokenRegistry *TokenRegistrySession) RemoveToken(_token common.Address, _index *big.Int) (*types.Transaction, error) {
	return _TokenRegistry.Contract.RemoveToken(&_TokenRegistry.TransactOpts, _token, _index)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x13baf1e6.
//
// Solidity: function removeToken(_token address, _index uint256) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) RemoveToken(_token common.Address, _index *big.Int) (*types.Transaction, error) {
	return _TokenRegistry.Contract.RemoveToken(&_TokenRegistry.TransactOpts, _token, _index)
}

// SetTokenIpfsHash is a paid mutator transaction binding the contract method 0xeef05f65.
//
// Solidity: function setTokenIpfsHash(_token address, _ipfsHash bytes) returns()
func (_TokenRegistry *TokenRegistryTransactor) SetTokenIpfsHash(opts *bind.TransactOpts, _token common.Address, _ipfsHash []byte) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "setTokenIpfsHash", _token, _ipfsHash)
}

// SetTokenIpfsHash is a paid mutator transaction binding the contract method 0xeef05f65.
//
// Solidity: function setTokenIpfsHash(_token address, _ipfsHash bytes) returns()
func (_TokenRegistry *TokenRegistrySession) SetTokenIpfsHash(_token common.Address, _ipfsHash []byte) (*types.Transaction, error) {
	return _TokenRegistry.Contract.SetTokenIpfsHash(&_TokenRegistry.TransactOpts, _token, _ipfsHash)
}

// SetTokenIpfsHash is a paid mutator transaction binding the contract method 0xeef05f65.
//
// Solidity: function setTokenIpfsHash(_token address, _ipfsHash bytes) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) SetTokenIpfsHash(_token common.Address, _ipfsHash []byte) (*types.Transaction, error) {
	return _TokenRegistry.Contract.SetTokenIpfsHash(&_TokenRegistry.TransactOpts, _token, _ipfsHash)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xc370c86d.
//
// Solidity: function setTokenName(_token address, _name string) returns()
func (_TokenRegistry *TokenRegistryTransactor) SetTokenName(opts *bind.TransactOpts, _token common.Address, _name string) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "setTokenName", _token, _name)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xc370c86d.
//
// Solidity: function setTokenName(_token address, _name string) returns()
func (_TokenRegistry *TokenRegistrySession) SetTokenName(_token common.Address, _name string) (*types.Transaction, error) {
	return _TokenRegistry.Contract.SetTokenName(&_TokenRegistry.TransactOpts, _token, _name)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xc370c86d.
//
// Solidity: function setTokenName(_token address, _name string) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) SetTokenName(_token common.Address, _name string) (*types.Transaction, error) {
	return _TokenRegistry.Contract.SetTokenName(&_TokenRegistry.TransactOpts, _token, _name)
}

// SetTokenSwarmHash is a paid mutator transaction binding the contract method 0x56318820.
//
// Solidity: function setTokenSwarmHash(_token address, _swarmHash bytes) returns()
func (_TokenRegistry *TokenRegistryTransactor) SetTokenSwarmHash(opts *bind.TransactOpts, _token common.Address, _swarmHash []byte) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "setTokenSwarmHash", _token, _swarmHash)
}

// SetTokenSwarmHash is a paid mutator transaction binding the contract method 0x56318820.
//
// Solidity: function setTokenSwarmHash(_token address, _swarmHash bytes) returns()
func (_TokenRegistry *TokenRegistrySession) SetTokenSwarmHash(_token common.Address, _swarmHash []byte) (*types.Transaction, error) {
	return _TokenRegistry.Contract.SetTokenSwarmHash(&_TokenRegistry.TransactOpts, _token, _swarmHash)
}

// SetTokenSwarmHash is a paid mutator transaction binding the contract method 0x56318820.
//
// Solidity: function setTokenSwarmHash(_token address, _swarmHash bytes) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) SetTokenSwarmHash(_token common.Address, _swarmHash []byte) (*types.Transaction, error) {
	return _TokenRegistry.Contract.SetTokenSwarmHash(&_TokenRegistry.TransactOpts, _token, _swarmHash)
}

// SetTokenSymbol is a paid mutator transaction binding the contract method 0xf036417f.
//
// Solidity: function setTokenSymbol(_token address, _symbol string) returns()
func (_TokenRegistry *TokenRegistryTransactor) SetTokenSymbol(opts *bind.TransactOpts, _token common.Address, _symbol string) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "setTokenSymbol", _token, _symbol)
}

// SetTokenSymbol is a paid mutator transaction binding the contract method 0xf036417f.
//
// Solidity: function setTokenSymbol(_token address, _symbol string) returns()
func (_TokenRegistry *TokenRegistrySession) SetTokenSymbol(_token common.Address, _symbol string) (*types.Transaction, error) {
	return _TokenRegistry.Contract.SetTokenSymbol(&_TokenRegistry.TransactOpts, _token, _symbol)
}

// SetTokenSymbol is a paid mutator transaction binding the contract method 0xf036417f.
//
// Solidity: function setTokenSymbol(_token address, _symbol string) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) SetTokenSymbol(_token common.Address, _symbol string) (*types.Transaction, error) {
	return _TokenRegistry.Contract.SetTokenSymbol(&_TokenRegistry.TransactOpts, _token, _symbol)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenRegistry *TokenRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenRegistry *TokenRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TransferOwnership(&_TokenRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TransferOwnership(&_TokenRegistry.TransactOpts, newOwner)
}

// TokenRegistryLogAddTokenIterator is returned from FilterLogAddToken and is used to iterate over the raw logs and unpacked data for LogAddToken events raised by the TokenRegistry contract.
type TokenRegistryLogAddTokenIterator struct {
	Event *TokenRegistryLogAddToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenRegistryLogAddTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryLogAddToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenRegistryLogAddToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenRegistryLogAddTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryLogAddTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryLogAddToken represents a LogAddToken event raised by the TokenRegistry contract.
type TokenRegistryLogAddToken struct {
	Token     common.Address
	Name      string
	Symbol    string
	Decimals  uint8
	IpfsHash  []byte
	SwarmHash []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogAddToken is a free log retrieval operation binding the contract event 0xd8d928b0b50ca11d9dc273236b46f3526515b03602f71f3a6af4f45bd9fa9144.
//
// Solidity: event LogAddToken(token indexed address, name string, symbol string, decimals uint8, ipfsHash bytes, swarmHash bytes)
func (_TokenRegistry *TokenRegistryFilterer) FilterLogAddToken(opts *bind.FilterOpts, token []common.Address) (*TokenRegistryLogAddTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "LogAddToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryLogAddTokenIterator{contract: _TokenRegistry.contract, event: "LogAddToken", logs: logs, sub: sub}, nil
}

// WatchLogAddToken is a free log subscription operation binding the contract event 0xd8d928b0b50ca11d9dc273236b46f3526515b03602f71f3a6af4f45bd9fa9144.
//
// Solidity: event LogAddToken(token indexed address, name string, symbol string, decimals uint8, ipfsHash bytes, swarmHash bytes)
func (_TokenRegistry *TokenRegistryFilterer) WatchLogAddToken(opts *bind.WatchOpts, sink chan<- *TokenRegistryLogAddToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "LogAddToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryLogAddToken)
				if err := _TokenRegistry.contract.UnpackLog(event, "LogAddToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenRegistryLogRemoveTokenIterator is returned from FilterLogRemoveToken and is used to iterate over the raw logs and unpacked data for LogRemoveToken events raised by the TokenRegistry contract.
type TokenRegistryLogRemoveTokenIterator struct {
	Event *TokenRegistryLogRemoveToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenRegistryLogRemoveTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryLogRemoveToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenRegistryLogRemoveToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenRegistryLogRemoveTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryLogRemoveTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryLogRemoveToken represents a LogRemoveToken event raised by the TokenRegistry contract.
type TokenRegistryLogRemoveToken struct {
	Token     common.Address
	Name      string
	Symbol    string
	Decimals  uint8
	IpfsHash  []byte
	SwarmHash []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogRemoveToken is a free log retrieval operation binding the contract event 0x32c54f1e2ea75844ded7517e7dbcd3895da7cd0c28f9ab9f9cf6ecf5f83762c6.
//
// Solidity: event LogRemoveToken(token indexed address, name string, symbol string, decimals uint8, ipfsHash bytes, swarmHash bytes)
func (_TokenRegistry *TokenRegistryFilterer) FilterLogRemoveToken(opts *bind.FilterOpts, token []common.Address) (*TokenRegistryLogRemoveTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "LogRemoveToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryLogRemoveTokenIterator{contract: _TokenRegistry.contract, event: "LogRemoveToken", logs: logs, sub: sub}, nil
}

// WatchLogRemoveToken is a free log subscription operation binding the contract event 0x32c54f1e2ea75844ded7517e7dbcd3895da7cd0c28f9ab9f9cf6ecf5f83762c6.
//
// Solidity: event LogRemoveToken(token indexed address, name string, symbol string, decimals uint8, ipfsHash bytes, swarmHash bytes)
func (_TokenRegistry *TokenRegistryFilterer) WatchLogRemoveToken(opts *bind.WatchOpts, sink chan<- *TokenRegistryLogRemoveToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "LogRemoveToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryLogRemoveToken)
				if err := _TokenRegistry.contract.UnpackLog(event, "LogRemoveToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenRegistryLogTokenIpfsHashChangeIterator is returned from FilterLogTokenIpfsHashChange and is used to iterate over the raw logs and unpacked data for LogTokenIpfsHashChange events raised by the TokenRegistry contract.
type TokenRegistryLogTokenIpfsHashChangeIterator struct {
	Event *TokenRegistryLogTokenIpfsHashChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenRegistryLogTokenIpfsHashChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryLogTokenIpfsHashChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenRegistryLogTokenIpfsHashChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenRegistryLogTokenIpfsHashChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryLogTokenIpfsHashChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryLogTokenIpfsHashChange represents a LogTokenIpfsHashChange event raised by the TokenRegistry contract.
type TokenRegistryLogTokenIpfsHashChange struct {
	Token       common.Address
	OldIpfsHash []byte
	NewIpfsHash []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogTokenIpfsHashChange is a free log retrieval operation binding the contract event 0x5b19f79ac4e8cfa820815502e11615f1a449e28155dc289ec5cac1a11f908694.
//
// Solidity: event LogTokenIpfsHashChange(token indexed address, oldIpfsHash bytes, newIpfsHash bytes)
func (_TokenRegistry *TokenRegistryFilterer) FilterLogTokenIpfsHashChange(opts *bind.FilterOpts, token []common.Address) (*TokenRegistryLogTokenIpfsHashChangeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "LogTokenIpfsHashChange", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryLogTokenIpfsHashChangeIterator{contract: _TokenRegistry.contract, event: "LogTokenIpfsHashChange", logs: logs, sub: sub}, nil
}

// WatchLogTokenIpfsHashChange is a free log subscription operation binding the contract event 0x5b19f79ac4e8cfa820815502e11615f1a449e28155dc289ec5cac1a11f908694.
//
// Solidity: event LogTokenIpfsHashChange(token indexed address, oldIpfsHash bytes, newIpfsHash bytes)
func (_TokenRegistry *TokenRegistryFilterer) WatchLogTokenIpfsHashChange(opts *bind.WatchOpts, sink chan<- *TokenRegistryLogTokenIpfsHashChange, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "LogTokenIpfsHashChange", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryLogTokenIpfsHashChange)
				if err := _TokenRegistry.contract.UnpackLog(event, "LogTokenIpfsHashChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenRegistryLogTokenNameChangeIterator is returned from FilterLogTokenNameChange and is used to iterate over the raw logs and unpacked data for LogTokenNameChange events raised by the TokenRegistry contract.
type TokenRegistryLogTokenNameChangeIterator struct {
	Event *TokenRegistryLogTokenNameChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenRegistryLogTokenNameChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryLogTokenNameChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenRegistryLogTokenNameChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenRegistryLogTokenNameChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryLogTokenNameChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryLogTokenNameChange represents a LogTokenNameChange event raised by the TokenRegistry contract.
type TokenRegistryLogTokenNameChange struct {
	Token   common.Address
	OldName string
	NewName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogTokenNameChange is a free log retrieval operation binding the contract event 0x4a6dbfc867b179991dec22ff19960f0a94d8d9d891fc556f547764670340e8ae.
//
// Solidity: event LogTokenNameChange(token indexed address, oldName string, newName string)
func (_TokenRegistry *TokenRegistryFilterer) FilterLogTokenNameChange(opts *bind.FilterOpts, token []common.Address) (*TokenRegistryLogTokenNameChangeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "LogTokenNameChange", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryLogTokenNameChangeIterator{contract: _TokenRegistry.contract, event: "LogTokenNameChange", logs: logs, sub: sub}, nil
}

// WatchLogTokenNameChange is a free log subscription operation binding the contract event 0x4a6dbfc867b179991dec22ff19960f0a94d8d9d891fc556f547764670340e8ae.
//
// Solidity: event LogTokenNameChange(token indexed address, oldName string, newName string)
func (_TokenRegistry *TokenRegistryFilterer) WatchLogTokenNameChange(opts *bind.WatchOpts, sink chan<- *TokenRegistryLogTokenNameChange, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "LogTokenNameChange", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryLogTokenNameChange)
				if err := _TokenRegistry.contract.UnpackLog(event, "LogTokenNameChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenRegistryLogTokenSwarmHashChangeIterator is returned from FilterLogTokenSwarmHashChange and is used to iterate over the raw logs and unpacked data for LogTokenSwarmHashChange events raised by the TokenRegistry contract.
type TokenRegistryLogTokenSwarmHashChangeIterator struct {
	Event *TokenRegistryLogTokenSwarmHashChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenRegistryLogTokenSwarmHashChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryLogTokenSwarmHashChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenRegistryLogTokenSwarmHashChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenRegistryLogTokenSwarmHashChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryLogTokenSwarmHashChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryLogTokenSwarmHashChange represents a LogTokenSwarmHashChange event raised by the TokenRegistry contract.
type TokenRegistryLogTokenSwarmHashChange struct {
	Token        common.Address
	OldSwarmHash []byte
	NewSwarmHash []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogTokenSwarmHashChange is a free log retrieval operation binding the contract event 0xc3168fdc13112e44a031057dbf6c609b33353addb4d8037d24543e22cbfe2acd.
//
// Solidity: event LogTokenSwarmHashChange(token indexed address, oldSwarmHash bytes, newSwarmHash bytes)
func (_TokenRegistry *TokenRegistryFilterer) FilterLogTokenSwarmHashChange(opts *bind.FilterOpts, token []common.Address) (*TokenRegistryLogTokenSwarmHashChangeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "LogTokenSwarmHashChange", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryLogTokenSwarmHashChangeIterator{contract: _TokenRegistry.contract, event: "LogTokenSwarmHashChange", logs: logs, sub: sub}, nil
}

// WatchLogTokenSwarmHashChange is a free log subscription operation binding the contract event 0xc3168fdc13112e44a031057dbf6c609b33353addb4d8037d24543e22cbfe2acd.
//
// Solidity: event LogTokenSwarmHashChange(token indexed address, oldSwarmHash bytes, newSwarmHash bytes)
func (_TokenRegistry *TokenRegistryFilterer) WatchLogTokenSwarmHashChange(opts *bind.WatchOpts, sink chan<- *TokenRegistryLogTokenSwarmHashChange, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "LogTokenSwarmHashChange", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryLogTokenSwarmHashChange)
				if err := _TokenRegistry.contract.UnpackLog(event, "LogTokenSwarmHashChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenRegistryLogTokenSymbolChangeIterator is returned from FilterLogTokenSymbolChange and is used to iterate over the raw logs and unpacked data for LogTokenSymbolChange events raised by the TokenRegistry contract.
type TokenRegistryLogTokenSymbolChangeIterator struct {
	Event *TokenRegistryLogTokenSymbolChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenRegistryLogTokenSymbolChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryLogTokenSymbolChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenRegistryLogTokenSymbolChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenRegistryLogTokenSymbolChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryLogTokenSymbolChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryLogTokenSymbolChange represents a LogTokenSymbolChange event raised by the TokenRegistry contract.
type TokenRegistryLogTokenSymbolChange struct {
	Token     common.Address
	OldSymbol string
	NewSymbol string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogTokenSymbolChange is a free log retrieval operation binding the contract event 0x53d878a6530e56c9bc96548fa0a8cae4f1d1f49c86b0e934c086b992ebb6998f.
//
// Solidity: event LogTokenSymbolChange(token indexed address, oldSymbol string, newSymbol string)
func (_TokenRegistry *TokenRegistryFilterer) FilterLogTokenSymbolChange(opts *bind.FilterOpts, token []common.Address) (*TokenRegistryLogTokenSymbolChangeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "LogTokenSymbolChange", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryLogTokenSymbolChangeIterator{contract: _TokenRegistry.contract, event: "LogTokenSymbolChange", logs: logs, sub: sub}, nil
}

// WatchLogTokenSymbolChange is a free log subscription operation binding the contract event 0x53d878a6530e56c9bc96548fa0a8cae4f1d1f49c86b0e934c086b992ebb6998f.
//
// Solidity: event LogTokenSymbolChange(token indexed address, oldSymbol string, newSymbol string)
func (_TokenRegistry *TokenRegistryFilterer) WatchLogTokenSymbolChange(opts *bind.WatchOpts, sink chan<- *TokenRegistryLogTokenSymbolChange, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "LogTokenSymbolChange", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryLogTokenSymbolChange)
				if err := _TokenRegistry.contract.UnpackLog(event, "LogTokenSymbolChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
