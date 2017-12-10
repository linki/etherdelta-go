// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TokenRegistryABI is the input ABI used to generate the binding from.
const TokenRegistryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"removeToken\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getTokenAddressByName\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"getTokenAddressBySymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_swarmHash\",\"type\":\"bytes\"}],\"name\":\"setTokenSwarmHash\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getTokenMetaData\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"_ipfsHash\",\"type\":\"bytes\"},{\"name\":\"_swarmHash\",\"type\":\"bytes\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setTokenName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"ipfsHash\",\"type\":\"bytes\"},{\"name\":\"swarmHash\",\"type\":\"bytes\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getTokenByName\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTokenAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_ipfsHash\",\"type\":\"bytes\"}],\"name\":\"setTokenIpfsHash\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"getTokenBySymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"setTokenSymbol\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"ipfsHash\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"swarmHash\",\"type\":\"bytes\"}],\"name\":\"LogAddToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"ipfsHash\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"swarmHash\",\"type\":\"bytes\"}],\"name\":\"LogRemoveToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"LogTokenNameChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldSymbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"newSymbol\",\"type\":\"string\"}],\"name\":\"LogTokenSymbolChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldIpfsHash\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"newIpfsHash\",\"type\":\"bytes\"}],\"name\":\"LogTokenIpfsHashChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldSwarmHash\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"newSwarmHash\",\"type\":\"bytes\"}],\"name\":\"LogTokenSwarmHashChange\",\"type\":\"event\"}]"

// TokenRegistry is an auto generated Go binding around an Ethereum contract.
type TokenRegistry struct {
	TokenRegistryCaller     // Read-only binding to the contract
	TokenRegistryTransactor // Write-only binding to the contract
}

// TokenRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenRegistryTransactor struct {
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
	contract, err := bindTokenRegistry(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenRegistry{TokenRegistryCaller: TokenRegistryCaller{contract: contract}, TokenRegistryTransactor: TokenRegistryTransactor{contract: contract}}, nil
}

// NewTokenRegistryCaller creates a new read-only instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryCaller(address common.Address, caller bind.ContractCaller) (*TokenRegistryCaller, error) {
	contract, err := bindTokenRegistry(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryCaller{contract: contract}, nil
}

// NewTokenRegistryTransactor creates a new write-only instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenRegistryTransactor, error) {
	contract, err := bindTokenRegistry(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryTransactor{contract: contract}, nil
}

// bindTokenRegistry binds a generic wrapper to an already deployed contract.
func bindTokenRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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
