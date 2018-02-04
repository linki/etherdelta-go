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

// AccountLevelsABI is the input ABI used to generate the binding from.
const AccountLevelsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"accountLevel\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AccountLevelsBin is the compiled bytecode used for deploying new contracts.
const AccountLevelsBin = `0x6060604052341561000f57600080fd5b60b08061001d6000396000f300606060405260043610603e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631cbd051981146043575b600080fd5b3415604d57600080fd5b606c73ffffffffffffffffffffffffffffffffffffffff60043516607e565b60405190815260200160405180910390f35b506000905600a165627a7a72305820028d9da2c5e80e6004244767020c20de3da52955c046c67a60dea503881ca5230029`

// DeployAccountLevels deploys a new Ethereum contract, binding an instance of AccountLevels to it.
func DeployAccountLevels(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccountLevels, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountLevelsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccountLevelsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountLevels{AccountLevelsCaller: AccountLevelsCaller{contract: contract}, AccountLevelsTransactor: AccountLevelsTransactor{contract: contract}}, nil
}

// AccountLevels is an auto generated Go binding around an Ethereum contract.
type AccountLevels struct {
	AccountLevelsCaller     // Read-only binding to the contract
	AccountLevelsTransactor // Write-only binding to the contract
}

// AccountLevelsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountLevelsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountLevelsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountLevelsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountLevelsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountLevelsSession struct {
	Contract     *AccountLevels    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountLevelsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountLevelsCallerSession struct {
	Contract *AccountLevelsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccountLevelsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountLevelsTransactorSession struct {
	Contract     *AccountLevelsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccountLevelsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountLevelsRaw struct {
	Contract *AccountLevels // Generic contract binding to access the raw methods on
}

// AccountLevelsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountLevelsCallerRaw struct {
	Contract *AccountLevelsCaller // Generic read-only contract binding to access the raw methods on
}

// AccountLevelsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountLevelsTransactorRaw struct {
	Contract *AccountLevelsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountLevels creates a new instance of AccountLevels, bound to a specific deployed contract.
func NewAccountLevels(address common.Address, backend bind.ContractBackend) (*AccountLevels, error) {
	contract, err := bindAccountLevels(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountLevels{AccountLevelsCaller: AccountLevelsCaller{contract: contract}, AccountLevelsTransactor: AccountLevelsTransactor{contract: contract}}, nil
}

// NewAccountLevelsCaller creates a new read-only instance of AccountLevels, bound to a specific deployed contract.
func NewAccountLevelsCaller(address common.Address, caller bind.ContractCaller) (*AccountLevelsCaller, error) {
	contract, err := bindAccountLevels(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AccountLevelsCaller{contract: contract}, nil
}

// NewAccountLevelsTransactor creates a new write-only instance of AccountLevels, bound to a specific deployed contract.
func NewAccountLevelsTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountLevelsTransactor, error) {
	contract, err := bindAccountLevels(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AccountLevelsTransactor{contract: contract}, nil
}

// bindAccountLevels binds a generic wrapper to an already deployed contract.
func bindAccountLevels(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountLevelsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountLevels *AccountLevelsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountLevels.Contract.AccountLevelsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountLevels *AccountLevelsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountLevels.Contract.AccountLevelsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountLevels *AccountLevelsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountLevels.Contract.AccountLevelsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountLevels *AccountLevelsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountLevels.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountLevels *AccountLevelsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountLevels.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountLevels *AccountLevelsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountLevels.Contract.contract.Transact(opts, method, params...)
}

// AccountLevel is a free data retrieval call binding the contract method 0x1cbd0519.
//
// Solidity: function accountLevel(user address) constant returns(uint256)
func (_AccountLevels *AccountLevelsCaller) AccountLevel(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccountLevels.contract.Call(opts, out, "accountLevel", user)
	return *ret0, err
}

// AccountLevel is a free data retrieval call binding the contract method 0x1cbd0519.
//
// Solidity: function accountLevel(user address) constant returns(uint256)
func (_AccountLevels *AccountLevelsSession) AccountLevel(user common.Address) (*big.Int, error) {
	return _AccountLevels.Contract.AccountLevel(&_AccountLevels.CallOpts, user)
}

// AccountLevel is a free data retrieval call binding the contract method 0x1cbd0519.
//
// Solidity: function accountLevel(user address) constant returns(uint256)
func (_AccountLevels *AccountLevelsCallerSession) AccountLevel(user common.Address) (*big.Int, error) {
	return _AccountLevels.Contract.AccountLevel(&_AccountLevels.CallOpts, user)
}

// AccountLevelsTestABI is the input ABI used to generate the binding from.
const AccountLevelsTestABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountLevels\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"accountLevel\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"level\",\"type\":\"uint256\"}],\"name\":\"setAccountLevel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AccountLevelsTestBin is the compiled bytecode used for deploying new contracts.
const AccountLevelsTestBin = `0x6060604052341561000f57600080fd5b6101858061001e6000396000f3006060604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166314577c55811461005b5780631cbd0519146100995780638abadb6b146100c5575b600080fd5b341561006657600080fd5b61008773ffffffffffffffffffffffffffffffffffffffff600435166100f6565b60405190815260200160405180910390f35b34156100a457600080fd5b61008773ffffffffffffffffffffffffffffffffffffffff60043516610108565b34156100d057600080fd5b6100f473ffffffffffffffffffffffffffffffffffffffff60043516602435610130565b005b60006020819052908152604090205481565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b73ffffffffffffffffffffffffffffffffffffffff9091166000908152602081905260409020555600a165627a7a723058207075bc6c613eb4e9f77a308c396f917693498a9e6aa65a39b6d5e3fba548b26e0029`

// DeployAccountLevelsTest deploys a new Ethereum contract, binding an instance of AccountLevelsTest to it.
func DeployAccountLevelsTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccountLevelsTest, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountLevelsTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccountLevelsTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountLevelsTest{AccountLevelsTestCaller: AccountLevelsTestCaller{contract: contract}, AccountLevelsTestTransactor: AccountLevelsTestTransactor{contract: contract}}, nil
}

// AccountLevelsTest is an auto generated Go binding around an Ethereum contract.
type AccountLevelsTest struct {
	AccountLevelsTestCaller     // Read-only binding to the contract
	AccountLevelsTestTransactor // Write-only binding to the contract
}

// AccountLevelsTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountLevelsTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountLevelsTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountLevelsTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountLevelsTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountLevelsTestSession struct {
	Contract     *AccountLevelsTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AccountLevelsTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountLevelsTestCallerSession struct {
	Contract *AccountLevelsTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AccountLevelsTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountLevelsTestTransactorSession struct {
	Contract     *AccountLevelsTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AccountLevelsTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountLevelsTestRaw struct {
	Contract *AccountLevelsTest // Generic contract binding to access the raw methods on
}

// AccountLevelsTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountLevelsTestCallerRaw struct {
	Contract *AccountLevelsTestCaller // Generic read-only contract binding to access the raw methods on
}

// AccountLevelsTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountLevelsTestTransactorRaw struct {
	Contract *AccountLevelsTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountLevelsTest creates a new instance of AccountLevelsTest, bound to a specific deployed contract.
func NewAccountLevelsTest(address common.Address, backend bind.ContractBackend) (*AccountLevelsTest, error) {
	contract, err := bindAccountLevelsTest(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountLevelsTest{AccountLevelsTestCaller: AccountLevelsTestCaller{contract: contract}, AccountLevelsTestTransactor: AccountLevelsTestTransactor{contract: contract}}, nil
}

// NewAccountLevelsTestCaller creates a new read-only instance of AccountLevelsTest, bound to a specific deployed contract.
func NewAccountLevelsTestCaller(address common.Address, caller bind.ContractCaller) (*AccountLevelsTestCaller, error) {
	contract, err := bindAccountLevelsTest(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AccountLevelsTestCaller{contract: contract}, nil
}

// NewAccountLevelsTestTransactor creates a new write-only instance of AccountLevelsTest, bound to a specific deployed contract.
func NewAccountLevelsTestTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountLevelsTestTransactor, error) {
	contract, err := bindAccountLevelsTest(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AccountLevelsTestTransactor{contract: contract}, nil
}

// bindAccountLevelsTest binds a generic wrapper to an already deployed contract.
func bindAccountLevelsTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountLevelsTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountLevelsTest *AccountLevelsTestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountLevelsTest.Contract.AccountLevelsTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountLevelsTest *AccountLevelsTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountLevelsTest.Contract.AccountLevelsTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountLevelsTest *AccountLevelsTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountLevelsTest.Contract.AccountLevelsTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountLevelsTest *AccountLevelsTestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountLevelsTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountLevelsTest *AccountLevelsTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountLevelsTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountLevelsTest *AccountLevelsTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountLevelsTest.Contract.contract.Transact(opts, method, params...)
}

// AccountLevel is a free data retrieval call binding the contract method 0x1cbd0519.
//
// Solidity: function accountLevel(user address) constant returns(uint256)
func (_AccountLevelsTest *AccountLevelsTestCaller) AccountLevel(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccountLevelsTest.contract.Call(opts, out, "accountLevel", user)
	return *ret0, err
}

// AccountLevel is a free data retrieval call binding the contract method 0x1cbd0519.
//
// Solidity: function accountLevel(user address) constant returns(uint256)
func (_AccountLevelsTest *AccountLevelsTestSession) AccountLevel(user common.Address) (*big.Int, error) {
	return _AccountLevelsTest.Contract.AccountLevel(&_AccountLevelsTest.CallOpts, user)
}

// AccountLevel is a free data retrieval call binding the contract method 0x1cbd0519.
//
// Solidity: function accountLevel(user address) constant returns(uint256)
func (_AccountLevelsTest *AccountLevelsTestCallerSession) AccountLevel(user common.Address) (*big.Int, error) {
	return _AccountLevelsTest.Contract.AccountLevel(&_AccountLevelsTest.CallOpts, user)
}

// AccountLevels is a free data retrieval call binding the contract method 0x14577c55.
//
// Solidity: function accountLevels( address) constant returns(uint256)
func (_AccountLevelsTest *AccountLevelsTestCaller) AccountLevels(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccountLevelsTest.contract.Call(opts, out, "accountLevels", arg0)
	return *ret0, err
}

// AccountLevels is a free data retrieval call binding the contract method 0x14577c55.
//
// Solidity: function accountLevels( address) constant returns(uint256)
func (_AccountLevelsTest *AccountLevelsTestSession) AccountLevels(arg0 common.Address) (*big.Int, error) {
	return _AccountLevelsTest.Contract.AccountLevels(&_AccountLevelsTest.CallOpts, arg0)
}

// AccountLevels is a free data retrieval call binding the contract method 0x14577c55.
//
// Solidity: function accountLevels( address) constant returns(uint256)
func (_AccountLevelsTest *AccountLevelsTestCallerSession) AccountLevels(arg0 common.Address) (*big.Int, error) {
	return _AccountLevelsTest.Contract.AccountLevels(&_AccountLevelsTest.CallOpts, arg0)
}

// SetAccountLevel is a paid mutator transaction binding the contract method 0x8abadb6b.
//
// Solidity: function setAccountLevel(user address, level uint256) returns()
func (_AccountLevelsTest *AccountLevelsTestTransactor) SetAccountLevel(opts *bind.TransactOpts, user common.Address, level *big.Int) (*types.Transaction, error) {
	return _AccountLevelsTest.contract.Transact(opts, "setAccountLevel", user, level)
}

// SetAccountLevel is a paid mutator transaction binding the contract method 0x8abadb6b.
//
// Solidity: function setAccountLevel(user address, level uint256) returns()
func (_AccountLevelsTest *AccountLevelsTestSession) SetAccountLevel(user common.Address, level *big.Int) (*types.Transaction, error) {
	return _AccountLevelsTest.Contract.SetAccountLevel(&_AccountLevelsTest.TransactOpts, user, level)
}

// SetAccountLevel is a paid mutator transaction binding the contract method 0x8abadb6b.
//
// Solidity: function setAccountLevel(user address, level uint256) returns()
func (_AccountLevelsTest *AccountLevelsTestTransactorSession) SetAccountLevel(user common.Address, level *big.Int) (*types.Transaction, error) {
	return _AccountLevelsTest.Contract.SetAccountLevel(&_AccountLevelsTest.TransactOpts, user, level)
}

// EtherDeltaABI is the input ABI used to generate the binding from.
const EtherDeltaABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"order\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderFills\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"amountFilled\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"feeMake_\",\"type\":\"uint256\"}],\"name\":\"changeFeeMake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeMake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"feeRebate_\",\"type\":\"uint256\"}],\"name\":\"changeFeeRebate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"testTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"feeAccount_\",\"type\":\"address\"}],\"name\":\"changeFeeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeRebate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"feeTake_\",\"type\":\"uint256\"}],\"name\":\"changeFeeTake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orders\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"accountLevelsAddr_\",\"type\":\"address\"}],\"name\":\"changeAccountLevelsAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountLevelsAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"availableVolume\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\"},{\"name\":\"feeAccount_\",\"type\":\"address\"},{\"name\":\"accountLevelsAddr_\",\"type\":\"address\"},{\"name\":\"feeMake_\",\"type\":\"uint256\"},{\"name\":\"feeTake_\",\"type\":\"uint256\"},{\"name\":\"feeRebate_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenGet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGet\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenGive\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGive\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expires\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Order\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenGet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGet\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenGive\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGive\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expires\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"v\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"r\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"Cancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenGet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGet\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenGive\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGive\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"get\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"give\",\"type\":\"address\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// EtherDeltaBin is the compiled bytecode used for deploying new contracts.
const EtherDeltaBin = `0x6060604052341561000f57600080fd5b60405160c080611a3d83398101604052808051919060200180519190602001805191906020018051919060200180519190602001805160008054600160a060020a0319908116600160a060020a039a8b16178255600180548216998b1699909917909855600280549098169690981695909517909555506003919091556004556005555061199a9081906100a390396000f30060606040526004361061013a5763ffffffff60e060020a6000350416630a19b14a811461014a5780630b9276661461019957806319774d43146101cb578063278b8c0e146101ff5780632e1a7d4d1461023e578063338b5dea1461025457806346be96c314610276578063508493bc146102bf57806354d03b5c146102e457806357786394146102fa5780635e1d7ae41461030d57806365e17c9d146103235780636c86888b1461035257806371ffcb16146103be578063731c2f81146103dd5780638823a9c0146103f05780638f283970146104065780639e281a9814610425578063bb5f462914610447578063c281309e14610469578063d0e30db01461047c578063e8f6bc2e14610484578063f3412942146104a3578063f7888aec146104b6578063f851a440146104db578063fb6e155f146104ee575b341561014557600080fd5b600080fd5b341561015557600080fd5b610197600160a060020a0360043581169060243590604435811690606435906084359060a4359060c4351660ff60e43516610104356101243561014435610537565b005b34156101a457600080fd5b610197600160a060020a03600435811690602435906044351660643560843560a4356107ea565b34156101d657600080fd5b6101ed600160a060020a036004351660243561092e565b60405190815260200160405180910390f35b341561020a57600080fd5b610197600160a060020a03600435811690602435906044351660643560843560a43560ff60c4351660e4356101043561094b565b341561024957600080fd5b610197600435610ba2565b341561025f57600080fd5b610197600160a060020a0360043516602435610cd7565b341561028157600080fd5b6101ed600160a060020a0360043581169060243590604435811690606435906084359060a4359060c4351660ff60e435166101043561012435610e37565b34156102ca57600080fd5b6101ed600160a060020a0360043581169060243516610efe565b34156102ef57600080fd5b610197600435610f1b565b341561030557600080fd5b6101ed610f4a565b341561031857600080fd5b610197600435610f50565b341561032e57600080fd5b610336610f8b565b604051600160a060020a03909116815260200160405180910390f35b341561035d57600080fd5b6103aa600160a060020a0360043581169060243590604435811690606435906084359060a4359060c43581169060ff60e43516906101043590610124359061014435906101643516610f9a565b604051901515815260200160405180910390f35b34156103c957600080fd5b610197600160a060020a0360043516611005565b34156103e857600080fd5b6101ed61104f565b34156103fb57600080fd5b610197600435611055565b341561041157600080fd5b610197600160a060020a0360043516611090565b341561043057600080fd5b610197600160a060020a03600435166024356110da565b341561045257600080fd5b6103aa600160a060020a0360043516602435611281565b341561047457600080fd5b6101ed6112a1565b6101976112a7565b341561048f57600080fd5b610197600160a060020a036004351661135c565b34156104ae57600080fd5b6103366113a6565b34156104c157600080fd5b6101ed600160a060020a03600435811690602435166113b5565b34156104e657600080fd5b6103366113e0565b34156104f957600080fd5b6101ed600160a060020a0360043581169060243590604435811690606435906084359060a4359060c4351660ff60e4351661010435610124356113ef565b60006002308d8d8d8d8d8d6000604051602001526040516c01000000000000000000000000600160a060020a0398891681028252968816870260148201526028810195909552929095169093026048830152605c820192909252607c810192909252609c82015260bc0160206040518083038160008661646e5a03f115156105be57600080fd5b50506040518051600160a060020a038816600090815260076020908152604080832084845290915290205490925060ff169050806106b3575085600160a060020a03166001826040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c810191909152603c0160405180910390208787876040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f1151561069f57600080fd5b505060206040510351600160a060020a0316145b80156106bf5750874311155b80156106f95750600160a060020a03861660009081526008602090815260408083208484529091529020548b906106f6908461161f565b11155b151561070457600080fd5b6107128c8c8c8c8a87611643565b600160a060020a0386166000908152600860209081526040808320848452909152902054610740908361161f565b600160a060020a03871660009081526008602090815260408083208584529091529020557f6effdda786735d5033bfad5f53e5131abcced9e52be6c507b62d639685fbed6d8c838c8e8d830281151561079557fe5b048a33604051600160a060020a0396871681526020810195909552928516604080860191909152606085019290925284166080840152921660a082015260c001905180910390a1505050505050505050505050565b60006002308888888888886000604051602001526040516c01000000000000000000000000600160a060020a0398891681028252968816870260148201526028810195909552929095169093026048830152605c820192909252607c810192909252609c82015260bc0160206040518083038160008661646e5a03f1151561087157600080fd5b50506040518051600160a060020a0333908116600090815260076020908152604080832085845290915290819020805460ff191660011790559193507f3f7f2eda73683c21a15f9435af1028c93185b5f1fa38270762dc32be606b3e8592508991899189918991899189919051600160a060020a03978816815260208101969096529386166040808701919091526060860193909352608085019190915260a0840152921660c082015260e001905180910390a150505050505050565b600860209081526000928352604080842090915290825290205481565b60006002308b8b8b8b8b8b6000604051602001526040516c01000000000000000000000000600160a060020a0398891681028252968816870260148201526028810195909552929095169093026048830152605c820192909252607c810192909252609c82015260bc0160206040518083038160008661646e5a03f115156109d257600080fd5b50506040518051600160a060020a033316600090815260076020908152604080832084845290915290205490925060ff16905080610ac7575033600160a060020a03166001826040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c810191909152603c0160405180910390208686866040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f11515610ab357600080fd5b505060206040510351600160a060020a0316145b1515610ad257600080fd5b33600160a060020a0381166000908152600860209081526040808320858452909152908190208b90557f1e0b760c386003e9cb9bcf4fcf3997886042859d9b6ed6320e804597fcdb28b0918c918c918c918c918c918c91908c908c908c9051600160a060020a039a8b16815260208101999099529689166040808a01919091526060890196909652608088019490945260a087019290925290951660c085015260ff90941660e084015261010083019390935261012082015261014001905180910390a150505050505050505050565b33600160a060020a0316600090815260008051602061194f833981519152602052604090205481901015610bd557600080fd5b33600160a060020a0316600090815260008051602061194f8339815191526020526040902054610c05908261190c565b33600160a060020a0316600081815260008051602061194f833981519152602052604090819020929092559082905160006040518083038185876187965a03f1925050501515610c5457600080fd5b33600160a060020a038116600090815260008051602061194f8339815191526020526040808220547ff341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567939185919051600160a060020a0394851681529290931660208301526040808301919091526060820192909252608001905180910390a150565b600160a060020a0382161515610cec57600080fd5b81600160a060020a03166323b872dd33308460006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610d5657600080fd5b6102c65a03f11515610d6757600080fd5b505050604051805190501515610d7c57600080fd5b600160a060020a0380831660009081526006602090815260408083203390941683529290522054610dad908261161f565b600160a060020a038381166000908152600660209081526040808320339485168452909152908190208390557fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d79285929185919051600160a060020a0394851681529290931660208301526040808301919091526060820192909252608001905180910390a15050565b6000806002308d8d8d8d8d8d6000604051602001526040516c01000000000000000000000000600160a060020a0398891681028252968816870260148201526028810195909552929095169093026048830152605c820192909252607c810192909252609c82015260bc0160206040518083038160008661646e5a03f11515610ebf57600080fd5b50506040518051600160a060020a03881660009081526008602090815260408083208484529091529020549350915050509a9950505050505050505050565b600660209081526000928352604080842090915290825290205481565b60005433600160a060020a03908116911614610f3657600080fd5b600354811115610f4557600080fd5b600355565b60035481565b60005433600160a060020a03908116911614610f6b57600080fd5b600554811080610f7c575060045481115b15610f8657600080fd5b600555565b600154600160a060020a031681565b600160a060020a03808d166000908152600660209081526040808320938516835292905290812054839010801590610fe3575082610fe08e8e8e8e8e8e8e8e8e8e6113ef565b10155b1515610ff157506000610ff5565b5060015b9c9b505050505050505050505050565b60005433600160a060020a0390811691161461102057600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60055481565b60005433600160a060020a0390811691161461107057600080fd5b600454811180611081575060055481105b1561108b57600080fd5b600455565b60005433600160a060020a039081169116146110ab57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a03821615156110ef57600080fd5b600160a060020a03808316600090815260066020908152604080832033909416835292905220548190101561112357600080fd5b600160a060020a0380831660009081526006602090815260408083203390941683529290522054611154908261190c565b600160a060020a03808416600081815260066020908152604080832033958616845290915280822094909455909263a9059cbb92918591516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156111d357600080fd5b6102c65a03f115156111e457600080fd5b5050506040518051905015156111f957600080fd5b600160a060020a03808316600090815260066020908152604080832033948516845290915290819020547ff341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb5679285929091859151600160a060020a0394851681529290931660208301526040808301919091526060820192909252608001905180910390a15050565b600760209081526000928352604080842090915290825290205460ff1681565b60045481565b33600160a060020a0316600090815260008051602061194f83398151915260205260409020546112d7903461161f565b33600160a060020a038116600090815260008051602061194f83398151915260205260408082208490557fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d793919291349151600160a060020a0394851681529290931660208301526040808301919091526060820192909252608001905180910390a1565b60005433600160a060020a0390811691161461137757600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a031681565b600160a060020a03918216600090815260066020908152604080832093909416825291909152205490565b600054600160a060020a031681565b6000806000806002308f8f8f8f8f8f6000604051602001526040516c01000000000000000000000000600160a060020a0398891681028252968816870260148201526028810195909552929095169093026048830152605c820192909252607c810192909252609c82015260bc0160206040518083038160008661646e5a03f1151561147a57600080fd5b50506040518051600160a060020a038a16600090815260076020908152604080832084845290915290205490945060ff1690508061156f575087600160a060020a03166001846040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c810191909152603c0160405180910390208989896040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f1151561155b57600080fd5b505060206040510351600160a060020a0316145b801561157b5750894311155b151561158a576000935061160e565b600160a060020a03881660009081526008602090815260408083208684529091529020546115b9908e9061190c565b600160a060020a03808e166000908152600660209081526040808320938d16835292905220549092508b906115ee908f611920565b8115156115f757fe5b0490508082101561160a5781935061160e565b8093505b5050509a9950505050505050505050565b600082820161163c8482108015906116375750838210155b61193f565b9392505050565b600080600080670de0b6b3a764000061165e86600354611920565b81151561166757fe5b049350670de0b6b3a764000061167f86600454611920565b81151561168857fe5b600254919004935060009250600160a060020a03161561174f57600254600160a060020a0316631cbd05198760006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156116fb57600080fd5b6102c65a03f1151561170c57600080fd5b5050506040518051915050600181141561174257670de0b6b3a764000061173586600554611920565b81151561173e57fe5b0491505b806002141561174f578291505b600160a060020a03808b166000908152600660209081526040808320339094168352929052205461178990611784878661161f565b61190c565b600160a060020a038b81166000908152600660209081526040808320338516845290915280822093909355908816815220546117d7906117d26117cc888661161f565b8761190c565b61161f565b600160a060020a038b811660009081526006602090815260408083208b851684529091528082209390935560015490911681522054611823906117d261181d878761161f565b8561190c565b600160a060020a03808c166000908152600660208181526040808420600154861685528252808420959095558c84168352908152838220928a168252919091522054611883908a6118748a89611920565b81151561187d57fe5b0461190c565b600160a060020a0389811660009081526006602090815260408083208b851684529091528082209390935533909116815220546118d4908a6118c58a89611920565b8115156118ce57fe5b0461161f565b600160a060020a03988916600090815260066020908152604080832033909c1683529a90529890982097909755505050505050505050565b600061191a8383111561193f565b50900390565b600082820261163c841580611637575083858381151561193c57fe5b04145b80151561194b57600080fd5b50560054cdd369e4e8a8515e52ca72ec816c2101831ad1f18bf44102ed171459c9b4f8a165627a7a72305820fcf29a2298c8abdf30a8d710540ac47d0bd33eb3c22d65b469a08379f7e935fe0029`

// DeployEtherDelta deploys a new Ethereum contract, binding an instance of EtherDelta to it.
func DeployEtherDelta(auth *bind.TransactOpts, backend bind.ContractBackend, admin_ common.Address, feeAccount_ common.Address, accountLevelsAddr_ common.Address, feeMake_ *big.Int, feeTake_ *big.Int, feeRebate_ *big.Int) (common.Address, *types.Transaction, *EtherDelta, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherDeltaABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EtherDeltaBin), backend, admin_, feeAccount_, accountLevelsAddr_, feeMake_, feeTake_, feeRebate_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtherDelta{EtherDeltaCaller: EtherDeltaCaller{contract: contract}, EtherDeltaTransactor: EtherDeltaTransactor{contract: contract}}, nil
}

// EtherDelta is an auto generated Go binding around an Ethereum contract.
type EtherDelta struct {
	EtherDeltaCaller     // Read-only binding to the contract
	EtherDeltaTransactor // Write-only binding to the contract
}

// EtherDeltaCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherDeltaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherDeltaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherDeltaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherDeltaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherDeltaSession struct {
	Contract     *EtherDelta       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherDeltaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherDeltaCallerSession struct {
	Contract *EtherDeltaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EtherDeltaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherDeltaTransactorSession struct {
	Contract     *EtherDeltaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EtherDeltaRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherDeltaRaw struct {
	Contract *EtherDelta // Generic contract binding to access the raw methods on
}

// EtherDeltaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherDeltaCallerRaw struct {
	Contract *EtherDeltaCaller // Generic read-only contract binding to access the raw methods on
}

// EtherDeltaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherDeltaTransactorRaw struct {
	Contract *EtherDeltaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherDelta creates a new instance of EtherDelta, bound to a specific deployed contract.
func NewEtherDelta(address common.Address, backend bind.ContractBackend) (*EtherDelta, error) {
	contract, err := bindEtherDelta(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherDelta{EtherDeltaCaller: EtherDeltaCaller{contract: contract}, EtherDeltaTransactor: EtherDeltaTransactor{contract: contract}}, nil
}

// NewEtherDeltaCaller creates a new read-only instance of EtherDelta, bound to a specific deployed contract.
func NewEtherDeltaCaller(address common.Address, caller bind.ContractCaller) (*EtherDeltaCaller, error) {
	contract, err := bindEtherDelta(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &EtherDeltaCaller{contract: contract}, nil
}

// NewEtherDeltaTransactor creates a new write-only instance of EtherDelta, bound to a specific deployed contract.
func NewEtherDeltaTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherDeltaTransactor, error) {
	contract, err := bindEtherDelta(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &EtherDeltaTransactor{contract: contract}, nil
}

// bindEtherDelta binds a generic wrapper to an already deployed contract.
func bindEtherDelta(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherDeltaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherDelta *EtherDeltaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtherDelta.Contract.EtherDeltaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherDelta *EtherDeltaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherDelta.Contract.EtherDeltaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherDelta *EtherDeltaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherDelta.Contract.EtherDeltaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherDelta *EtherDeltaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtherDelta.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherDelta *EtherDeltaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherDelta.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherDelta *EtherDeltaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherDelta.Contract.contract.Transact(opts, method, params...)
}

// AccountLevelsAddr is a free data retrieval call binding the contract method 0xf3412942.
//
// Solidity: function accountLevelsAddr() constant returns(address)
func (_EtherDelta *EtherDeltaCaller) AccountLevelsAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "accountLevelsAddr")
	return *ret0, err
}

// AccountLevelsAddr is a free data retrieval call binding the contract method 0xf3412942.
//
// Solidity: function accountLevelsAddr() constant returns(address)
func (_EtherDelta *EtherDeltaSession) AccountLevelsAddr() (common.Address, error) {
	return _EtherDelta.Contract.AccountLevelsAddr(&_EtherDelta.CallOpts)
}

// AccountLevelsAddr is a free data retrieval call binding the contract method 0xf3412942.
//
// Solidity: function accountLevelsAddr() constant returns(address)
func (_EtherDelta *EtherDeltaCallerSession) AccountLevelsAddr() (common.Address, error) {
	return _EtherDelta.Contract.AccountLevelsAddr(&_EtherDelta.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_EtherDelta *EtherDeltaCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_EtherDelta *EtherDeltaSession) Admin() (common.Address, error) {
	return _EtherDelta.Contract.Admin(&_EtherDelta.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_EtherDelta *EtherDeltaCallerSession) Admin() (common.Address, error) {
	return _EtherDelta.Contract.Admin(&_EtherDelta.CallOpts)
}

// AmountFilled is a free data retrieval call binding the contract method 0x46be96c3.
//
// Solidity: function amountFilled(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) AmountFilled(opts *bind.CallOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "amountFilled", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
	return *ret0, err
}

// AmountFilled is a free data retrieval call binding the contract method 0x46be96c3.
//
// Solidity: function amountFilled(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) AmountFilled(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.AmountFilled(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
}

// AmountFilled is a free data retrieval call binding the contract method 0x46be96c3.
//
// Solidity: function amountFilled(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) AmountFilled(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.AmountFilled(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
}

// AvailableVolume is a free data retrieval call binding the contract method 0xfb6e155f.
//
// Solidity: function availableVolume(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) AvailableVolume(opts *bind.CallOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "availableVolume", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
	return *ret0, err
}

// AvailableVolume is a free data retrieval call binding the contract method 0xfb6e155f.
//
// Solidity: function availableVolume(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) AvailableVolume(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.AvailableVolume(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
}

// AvailableVolume is a free data retrieval call binding the contract method 0xfb6e155f.
//
// Solidity: function availableVolume(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) AvailableVolume(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.AvailableVolume(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) BalanceOf(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "balanceOf", token, user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) BalanceOf(token common.Address, user common.Address) (*big.Int, error) {
	return _EtherDelta.Contract.BalanceOf(&_EtherDelta.CallOpts, token, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) BalanceOf(token common.Address, user common.Address) (*big.Int, error) {
	return _EtherDelta.Contract.BalanceOf(&_EtherDelta.CallOpts, token, user)
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() constant returns(address)
func (_EtherDelta *EtherDeltaCaller) FeeAccount(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "feeAccount")
	return *ret0, err
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() constant returns(address)
func (_EtherDelta *EtherDeltaSession) FeeAccount() (common.Address, error) {
	return _EtherDelta.Contract.FeeAccount(&_EtherDelta.CallOpts)
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() constant returns(address)
func (_EtherDelta *EtherDeltaCallerSession) FeeAccount() (common.Address, error) {
	return _EtherDelta.Contract.FeeAccount(&_EtherDelta.CallOpts)
}

// FeeMake is a free data retrieval call binding the contract method 0x57786394.
//
// Solidity: function feeMake() constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) FeeMake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "feeMake")
	return *ret0, err
}

// FeeMake is a free data retrieval call binding the contract method 0x57786394.
//
// Solidity: function feeMake() constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) FeeMake() (*big.Int, error) {
	return _EtherDelta.Contract.FeeMake(&_EtherDelta.CallOpts)
}

// FeeMake is a free data retrieval call binding the contract method 0x57786394.
//
// Solidity: function feeMake() constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) FeeMake() (*big.Int, error) {
	return _EtherDelta.Contract.FeeMake(&_EtherDelta.CallOpts)
}

// FeeRebate is a free data retrieval call binding the contract method 0x731c2f81.
//
// Solidity: function feeRebate() constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) FeeRebate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "feeRebate")
	return *ret0, err
}

// FeeRebate is a free data retrieval call binding the contract method 0x731c2f81.
//
// Solidity: function feeRebate() constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) FeeRebate() (*big.Int, error) {
	return _EtherDelta.Contract.FeeRebate(&_EtherDelta.CallOpts)
}

// FeeRebate is a free data retrieval call binding the contract method 0x731c2f81.
//
// Solidity: function feeRebate() constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) FeeRebate() (*big.Int, error) {
	return _EtherDelta.Contract.FeeRebate(&_EtherDelta.CallOpts)
}

// FeeTake is a free data retrieval call binding the contract method 0xc281309e.
//
// Solidity: function feeTake() constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) FeeTake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "feeTake")
	return *ret0, err
}

// FeeTake is a free data retrieval call binding the contract method 0xc281309e.
//
// Solidity: function feeTake() constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) FeeTake() (*big.Int, error) {
	return _EtherDelta.Contract.FeeTake(&_EtherDelta.CallOpts)
}

// FeeTake is a free data retrieval call binding the contract method 0xc281309e.
//
// Solidity: function feeTake() constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) FeeTake() (*big.Int, error) {
	return _EtherDelta.Contract.FeeTake(&_EtherDelta.CallOpts)
}

// OrderFills is a free data retrieval call binding the contract method 0x19774d43.
//
// Solidity: function orderFills( address,  bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) OrderFills(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "orderFills", arg0, arg1)
	return *ret0, err
}

// OrderFills is a free data retrieval call binding the contract method 0x19774d43.
//
// Solidity: function orderFills( address,  bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) OrderFills(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.OrderFills(&_EtherDelta.CallOpts, arg0, arg1)
}

// OrderFills is a free data retrieval call binding the contract method 0x19774d43.
//
// Solidity: function orderFills( address,  bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) OrderFills(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.OrderFills(&_EtherDelta.CallOpts, arg0, arg1)
}

// Orders is a free data retrieval call binding the contract method 0xbb5f4629.
//
// Solidity: function orders( address,  bytes32) constant returns(bool)
func (_EtherDelta *EtherDeltaCaller) Orders(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "orders", arg0, arg1)
	return *ret0, err
}

// Orders is a free data retrieval call binding the contract method 0xbb5f4629.
//
// Solidity: function orders( address,  bytes32) constant returns(bool)
func (_EtherDelta *EtherDeltaSession) Orders(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _EtherDelta.Contract.Orders(&_EtherDelta.CallOpts, arg0, arg1)
}

// Orders is a free data retrieval call binding the contract method 0xbb5f4629.
//
// Solidity: function orders( address,  bytes32) constant returns(bool)
func (_EtherDelta *EtherDeltaCallerSession) Orders(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _EtherDelta.Contract.Orders(&_EtherDelta.CallOpts, arg0, arg1)
}

// TestTrade is a free data retrieval call binding the contract method 0x6c86888b.
//
// Solidity: function testTrade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256, sender address) constant returns(bool)
func (_EtherDelta *EtherDeltaCaller) TestTrade(opts *bind.CallOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "testTrade", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount, sender)
	return *ret0, err
}

// TestTrade is a free data retrieval call binding the contract method 0x6c86888b.
//
// Solidity: function testTrade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256, sender address) constant returns(bool)
func (_EtherDelta *EtherDeltaSession) TestTrade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int, sender common.Address) (bool, error) {
	return _EtherDelta.Contract.TestTrade(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount, sender)
}

// TestTrade is a free data retrieval call binding the contract method 0x6c86888b.
//
// Solidity: function testTrade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256, sender address) constant returns(bool)
func (_EtherDelta *EtherDeltaCallerSession) TestTrade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int, sender common.Address) (bool, error) {
	return _EtherDelta.Contract.TestTrade(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount, sender)
}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens( address,  address) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) Tokens(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "tokens", arg0, arg1)
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens( address,  address) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) Tokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EtherDelta.Contract.Tokens(&_EtherDelta.CallOpts, arg0, arg1)
}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens( address,  address) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) Tokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EtherDelta.Contract.Tokens(&_EtherDelta.CallOpts, arg0, arg1)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x278b8c0e.
//
// Solidity: function cancelOrder(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) returns()
func (_EtherDelta *EtherDeltaTransactor) CancelOrder(opts *bind.TransactOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "cancelOrder", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x278b8c0e.
//
// Solidity: function cancelOrder(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) returns()
func (_EtherDelta *EtherDeltaSession) CancelOrder(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _EtherDelta.Contract.CancelOrder(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x278b8c0e.
//
// Solidity: function cancelOrder(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) returns()
func (_EtherDelta *EtherDeltaTransactorSession) CancelOrder(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _EtherDelta.Contract.CancelOrder(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
}

// ChangeAccountLevelsAddr is a paid mutator transaction binding the contract method 0xe8f6bc2e.
//
// Solidity: function changeAccountLevelsAddr(accountLevelsAddr_ address) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeAccountLevelsAddr(opts *bind.TransactOpts, accountLevelsAddr_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeAccountLevelsAddr", accountLevelsAddr_)
}

// ChangeAccountLevelsAddr is a paid mutator transaction binding the contract method 0xe8f6bc2e.
//
// Solidity: function changeAccountLevelsAddr(accountLevelsAddr_ address) returns()
func (_EtherDelta *EtherDeltaSession) ChangeAccountLevelsAddr(accountLevelsAddr_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeAccountLevelsAddr(&_EtherDelta.TransactOpts, accountLevelsAddr_)
}

// ChangeAccountLevelsAddr is a paid mutator transaction binding the contract method 0xe8f6bc2e.
//
// Solidity: function changeAccountLevelsAddr(accountLevelsAddr_ address) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeAccountLevelsAddr(accountLevelsAddr_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeAccountLevelsAddr(&_EtherDelta.TransactOpts, accountLevelsAddr_)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(admin_ address) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeAdmin(opts *bind.TransactOpts, admin_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeAdmin", admin_)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(admin_ address) returns()
func (_EtherDelta *EtherDeltaSession) ChangeAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeAdmin(&_EtherDelta.TransactOpts, admin_)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(admin_ address) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeAdmin(&_EtherDelta.TransactOpts, admin_)
}

// ChangeFeeAccount is a paid mutator transaction binding the contract method 0x71ffcb16.
//
// Solidity: function changeFeeAccount(feeAccount_ address) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeFeeAccount(opts *bind.TransactOpts, feeAccount_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeFeeAccount", feeAccount_)
}

// ChangeFeeAccount is a paid mutator transaction binding the contract method 0x71ffcb16.
//
// Solidity: function changeFeeAccount(feeAccount_ address) returns()
func (_EtherDelta *EtherDeltaSession) ChangeFeeAccount(feeAccount_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeAccount(&_EtherDelta.TransactOpts, feeAccount_)
}

// ChangeFeeAccount is a paid mutator transaction binding the contract method 0x71ffcb16.
//
// Solidity: function changeFeeAccount(feeAccount_ address) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeFeeAccount(feeAccount_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeAccount(&_EtherDelta.TransactOpts, feeAccount_)
}

// ChangeFeeMake is a paid mutator transaction binding the contract method 0x54d03b5c.
//
// Solidity: function changeFeeMake(feeMake_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeFeeMake(opts *bind.TransactOpts, feeMake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeFeeMake", feeMake_)
}

// ChangeFeeMake is a paid mutator transaction binding the contract method 0x54d03b5c.
//
// Solidity: function changeFeeMake(feeMake_ uint256) returns()
func (_EtherDelta *EtherDeltaSession) ChangeFeeMake(feeMake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeMake(&_EtherDelta.TransactOpts, feeMake_)
}

// ChangeFeeMake is a paid mutator transaction binding the contract method 0x54d03b5c.
//
// Solidity: function changeFeeMake(feeMake_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeFeeMake(feeMake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeMake(&_EtherDelta.TransactOpts, feeMake_)
}

// ChangeFeeRebate is a paid mutator transaction binding the contract method 0x5e1d7ae4.
//
// Solidity: function changeFeeRebate(feeRebate_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeFeeRebate(opts *bind.TransactOpts, feeRebate_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeFeeRebate", feeRebate_)
}

// ChangeFeeRebate is a paid mutator transaction binding the contract method 0x5e1d7ae4.
//
// Solidity: function changeFeeRebate(feeRebate_ uint256) returns()
func (_EtherDelta *EtherDeltaSession) ChangeFeeRebate(feeRebate_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeRebate(&_EtherDelta.TransactOpts, feeRebate_)
}

// ChangeFeeRebate is a paid mutator transaction binding the contract method 0x5e1d7ae4.
//
// Solidity: function changeFeeRebate(feeRebate_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeFeeRebate(feeRebate_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeRebate(&_EtherDelta.TransactOpts, feeRebate_)
}

// ChangeFeeTake is a paid mutator transaction binding the contract method 0x8823a9c0.
//
// Solidity: function changeFeeTake(feeTake_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeFeeTake(opts *bind.TransactOpts, feeTake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeFeeTake", feeTake_)
}

// ChangeFeeTake is a paid mutator transaction binding the contract method 0x8823a9c0.
//
// Solidity: function changeFeeTake(feeTake_ uint256) returns()
func (_EtherDelta *EtherDeltaSession) ChangeFeeTake(feeTake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeTake(&_EtherDelta.TransactOpts, feeTake_)
}

// ChangeFeeTake is a paid mutator transaction binding the contract method 0x8823a9c0.
//
// Solidity: function changeFeeTake(feeTake_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeFeeTake(feeTake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeTake(&_EtherDelta.TransactOpts, feeTake_)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_EtherDelta *EtherDeltaTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_EtherDelta *EtherDeltaSession) Deposit() (*types.Transaction, error) {
	return _EtherDelta.Contract.Deposit(&_EtherDelta.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_EtherDelta *EtherDeltaTransactorSession) Deposit() (*types.Transaction, error) {
	return _EtherDelta.Contract.Deposit(&_EtherDelta.TransactOpts)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) DepositToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "depositToken", token, amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaSession) DepositToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.DepositToken(&_EtherDelta.TransactOpts, token, amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) DepositToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.DepositToken(&_EtherDelta.TransactOpts, token, amount)
}

// Order is a paid mutator transaction binding the contract method 0x0b927666.
//
// Solidity: function order(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) Order(opts *bind.TransactOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "order", tokenGet, amountGet, tokenGive, amountGive, expires, nonce)
}

// Order is a paid mutator transaction binding the contract method 0x0b927666.
//
// Solidity: function order(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256) returns()
func (_EtherDelta *EtherDeltaSession) Order(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Order(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce)
}

// Order is a paid mutator transaction binding the contract method 0x0b927666.
//
// Solidity: function order(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) Order(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Order(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce)
}

// Trade is a paid mutator transaction binding the contract method 0x0a19b14a.
//
// Solidity: function trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) Trade(opts *bind.TransactOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "trade", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Trade is a paid mutator transaction binding the contract method 0x0a19b14a.
//
// Solidity: function trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_EtherDelta *EtherDeltaSession) Trade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Trade(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Trade is a paid mutator transaction binding the contract method 0x0a19b14a.
//
// Solidity: function trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) Trade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Trade(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_EtherDelta *EtherDeltaSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Withdraw(&_EtherDelta.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Withdraw(&_EtherDelta.TransactOpts, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "withdrawToken", token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.WithdrawToken(&_EtherDelta.TransactOpts, token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.WithdrawToken(&_EtherDelta.TransactOpts, token, amount)
}

// ReserveTokenABI is the input ABI used to generate the binding from.
const ReserveTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ReserveTokenBin is the compiled bytecode used for deploying new contracts.
const ReserveTokenBin = `0x6060604052341561000f57600080fd5b60058054600160a060020a03191633600160a060020a0316179055610748806100396000396000f3006060604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100b3578063075461721461013d578063095ea7b31461016c5780630ecaea73146101a257806318160ddd146101c657806323b872dd146101eb578063313ce5671461021357806370a0823114610226578063a24835d114610245578063a9059cbb14610267578063dd62ed3e14610289575b600080fd5b34156100be57600080fd5b6100c66102ae565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101025780820151838201526020016100ea565b50505050905090810190601f16801561012f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561014857600080fd5b61015061034c565b604051600160a060020a03909116815260200160405180910390f35b341561017757600080fd5b61018e600160a060020a036004351660243561035b565b604051901515815260200160405180910390f35b34156101ad57600080fd5b6101c4600160a060020a03600435166024356103c8565b005b34156101d157600080fd5b6101d9610433565b60405190815260200160405180910390f35b34156101f657600080fd5b61018e600160a060020a0360043581169060243516604435610439565b341561021e57600080fd5b6101d961054a565b341561023157600080fd5b6101d9600160a060020a0360043516610550565b341561025057600080fd5b6101c4600160a060020a036004351660243561056b565b341561027257600080fd5b61018e600160a060020a03600435166024356105f5565b341561029457600080fd5b6101d9600160a060020a03600435811690602435166106b1565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103445780601f1061031957610100808354040283529160200191610344565b820191906000526020600020905b81548152906001019060200180831161032757829003601f168201915b505050505081565b600554600160a060020a031681565b600160a060020a03338116600081815260036020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60055433600160a060020a039081169116146103e357600080fd5b600160a060020a03821660009081526002602052604090205461040690826106dc565b600160a060020a03831660009081526002602052604090205560045461042c90826106dc565b6004555050565b60045481565b600160a060020a0383166000908152600260205260408120548290108015906104895750600160a060020a0380851660009081526003602090815260408083203390941683529290522054829010155b80156104ae5750600160a060020a038316600090815260026020526040902054828101115b1561053f57600160a060020a03808416600081815260026020908152604080832080548801905588851680845281842080548990039055600383528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610543565b5060005b9392505050565b60005481565b600160a060020a031660009081526002602052604090205490565b60055433600160a060020a0390811691161461058657600080fd5b600160a060020a038216600090815260026020526040902054819010156105ac57600080fd5b600160a060020a0382166000908152600260205260409020546105cf90826106f9565b600160a060020a03831660009081526002602052604090205560045461042c90826106f9565b600160a060020a0333166000908152600260205260408120548290108015906106375750600160a060020a038316600090815260026020526040902054828101115b156106a957600160a060020a033381166000818152600260205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060016103c2565b5060006103c2565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b60008282016105438482108015906106f45750838210155b61070d565b60006107078383111561070d565b50900390565b80151561071957600080fd5b505600a165627a7a72305820487c81221773c58afec08cf98823302d78e12670ea6fc17ded99e5df1d6535ea0029`

// DeployReserveToken deploys a new Ethereum contract, binding an instance of ReserveToken to it.
func DeployReserveToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReserveToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReserveTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReserveToken{ReserveTokenCaller: ReserveTokenCaller{contract: contract}, ReserveTokenTransactor: ReserveTokenTransactor{contract: contract}}, nil
}

// ReserveToken is an auto generated Go binding around an Ethereum contract.
type ReserveToken struct {
	ReserveTokenCaller     // Read-only binding to the contract
	ReserveTokenTransactor // Write-only binding to the contract
}

// ReserveTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReserveTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReserveTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReserveTokenSession struct {
	Contract     *ReserveToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReserveTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReserveTokenCallerSession struct {
	Contract *ReserveTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ReserveTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReserveTokenTransactorSession struct {
	Contract     *ReserveTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ReserveTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReserveTokenRaw struct {
	Contract *ReserveToken // Generic contract binding to access the raw methods on
}

// ReserveTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReserveTokenCallerRaw struct {
	Contract *ReserveTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ReserveTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReserveTokenTransactorRaw struct {
	Contract *ReserveTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReserveToken creates a new instance of ReserveToken, bound to a specific deployed contract.
func NewReserveToken(address common.Address, backend bind.ContractBackend) (*ReserveToken, error) {
	contract, err := bindReserveToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReserveToken{ReserveTokenCaller: ReserveTokenCaller{contract: contract}, ReserveTokenTransactor: ReserveTokenTransactor{contract: contract}}, nil
}

// NewReserveTokenCaller creates a new read-only instance of ReserveToken, bound to a specific deployed contract.
func NewReserveTokenCaller(address common.Address, caller bind.ContractCaller) (*ReserveTokenCaller, error) {
	contract, err := bindReserveToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveTokenCaller{contract: contract}, nil
}

// NewReserveTokenTransactor creates a new write-only instance of ReserveToken, bound to a specific deployed contract.
func NewReserveTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ReserveTokenTransactor, error) {
	contract, err := bindReserveToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ReserveTokenTransactor{contract: contract}, nil
}

// bindReserveToken binds a generic wrapper to an already deployed contract.
func bindReserveToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReserveToken *ReserveTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReserveToken.Contract.ReserveTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReserveToken *ReserveTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveToken.Contract.ReserveTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReserveToken *ReserveTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReserveToken.Contract.ReserveTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReserveToken *ReserveTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReserveToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReserveToken *ReserveTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReserveToken *ReserveTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReserveToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ReserveToken *ReserveTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReserveToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ReserveToken *ReserveTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ReserveToken.Contract.Allowance(&_ReserveToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ReserveToken *ReserveTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ReserveToken.Contract.Allowance(&_ReserveToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ReserveToken *ReserveTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReserveToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ReserveToken *ReserveTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ReserveToken.Contract.BalanceOf(&_ReserveToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ReserveToken *ReserveTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ReserveToken.Contract.BalanceOf(&_ReserveToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_ReserveToken *ReserveTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReserveToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_ReserveToken *ReserveTokenSession) Decimals() (*big.Int, error) {
	return _ReserveToken.Contract.Decimals(&_ReserveToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_ReserveToken *ReserveTokenCallerSession) Decimals() (*big.Int, error) {
	return _ReserveToken.Contract.Decimals(&_ReserveToken.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() constant returns(address)
func (_ReserveToken *ReserveTokenCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReserveToken.contract.Call(opts, out, "minter")
	return *ret0, err
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() constant returns(address)
func (_ReserveToken *ReserveTokenSession) Minter() (common.Address, error) {
	return _ReserveToken.Contract.Minter(&_ReserveToken.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() constant returns(address)
func (_ReserveToken *ReserveTokenCallerSession) Minter() (common.Address, error) {
	return _ReserveToken.Contract.Minter(&_ReserveToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReserveToken *ReserveTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ReserveToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReserveToken *ReserveTokenSession) Name() (string, error) {
	return _ReserveToken.Contract.Name(&_ReserveToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReserveToken *ReserveTokenCallerSession) Name() (string, error) {
	return _ReserveToken.Contract.Name(&_ReserveToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReserveToken *ReserveTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReserveToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReserveToken *ReserveTokenSession) TotalSupply() (*big.Int, error) {
	return _ReserveToken.Contract.TotalSupply(&_ReserveToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReserveToken *ReserveTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ReserveToken.Contract.TotalSupply(&_ReserveToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Approve(&_ReserveToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Approve(&_ReserveToken.TransactOpts, _spender, _value)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(account address, amount uint256) returns()
func (_ReserveToken *ReserveTokenTransactor) Create(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveToken.contract.Transact(opts, "create", account, amount)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(account address, amount uint256) returns()
func (_ReserveToken *ReserveTokenSession) Create(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Create(&_ReserveToken.TransactOpts, account, amount)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(account address, amount uint256) returns()
func (_ReserveToken *ReserveTokenTransactorSession) Create(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Create(&_ReserveToken.TransactOpts, account, amount)
}

// Destroy is a paid mutator transaction binding the contract method 0xa24835d1.
//
// Solidity: function destroy(account address, amount uint256) returns()
func (_ReserveToken *ReserveTokenTransactor) Destroy(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveToken.contract.Transact(opts, "destroy", account, amount)
}

// Destroy is a paid mutator transaction binding the contract method 0xa24835d1.
//
// Solidity: function destroy(account address, amount uint256) returns()
func (_ReserveToken *ReserveTokenSession) Destroy(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Destroy(&_ReserveToken.TransactOpts, account, amount)
}

// Destroy is a paid mutator transaction binding the contract method 0xa24835d1.
//
// Solidity: function destroy(account address, amount uint256) returns()
func (_ReserveToken *ReserveTokenTransactorSession) Destroy(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Destroy(&_ReserveToken.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Transfer(&_ReserveToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Transfer(&_ReserveToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.TransferFrom(&_ReserveToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.TransferFrom(&_ReserveToken.TransactOpts, _from, _to, _value)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a7230582061f6ada94d042e481762cc6a282f68861a87e0e2e50b541da7053bc7ea5a2a330029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x6060604052341561000f57600080fd5b61056e8061001e6000396000f30060606040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610092578063095ea7b31461011c57806318160ddd1461015257806323b872dd14610177578063313ce5671461019f57806370a08231146101b2578063a9059cbb146101d1578063dd62ed3e146101f3575b600080fd5b341561009d57600080fd5b6100a5610218565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100e15780820151838201526020016100c9565b50505050905090810190601f16801561010e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561012757600080fd5b61013e600160a060020a03600435166024356102b6565b604051901515815260200160405180910390f35b341561015d57600080fd5b610165610323565b60405190815260200160405180910390f35b341561018257600080fd5b61013e600160a060020a0360043581169060243516604435610329565b34156101aa57600080fd5b61016561043a565b34156101bd57600080fd5b610165600160a060020a0360043516610440565b34156101dc57600080fd5b61013e600160a060020a036004351660243561045b565b34156101fe57600080fd5b610165600160a060020a0360043581169060243516610517565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102ae5780601f10610283576101008083540402835291602001916102ae565b820191906000526020600020905b81548152906001019060200180831161029157829003601f168201915b505050505081565b600160a060020a03338116600081815260036020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60045481565b600160a060020a0383166000908152600260205260408120548290108015906103795750600160a060020a0380851660009081526003602090815260408083203390941683529290522054829010155b801561039e5750600160a060020a038316600090815260026020526040902054828101115b1561042f57600160a060020a03808416600081815260026020908152604080832080548801905588851680845281842080548990039055600383528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610433565b5060005b9392505050565b60005481565b600160a060020a031660009081526002602052604090205490565b600160a060020a03331660009081526002602052604081205482901080159061049d5750600160a060020a038316600090815260026020526040902054828101115b1561050f57600160a060020a033381166000818152600260205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600161031d565b50600061031d565b600160a060020a039182166000908152600360209081526040808320939094168252919091522054905600a165627a7a72305820deaffa69935c61441d9429d69a5925561609407d2591fbf9888c5810a670c39c0029`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_StandardToken *StandardTokenSession) Decimals() (*big.Int, error) {
	return _StandardToken.Contract.Decimals(&_StandardToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) Decimals() (*big.Int, error) {
	return _StandardToken.Contract.Decimals(&_StandardToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StandardToken *StandardTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StandardToken *StandardTokenSession) Name() (string, error) {
	return _StandardToken.Contract.Name(&_StandardToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StandardToken *StandardTokenCallerSession) Name() (string, error) {
	return _StandardToken.Contract.Name(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x6060604052341561000f57600080fd5b6103168061001e6000396000f30060606040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610092578063095ea7b31461011c57806318160ddd1461015f57806323b872dd14610184578063313ce567146101b957806370a08231146101cc578063a9059cbb1461011c578063dd62ed3e146101f8575b600080fd5b341561009d57600080fd5b6100a561022a565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100e15780820151838201526020016100c9565b50505050905090810190601f16801561010e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561012757600080fd5b61014b73ffffffffffffffffffffffffffffffffffffffff600435166024356102c8565b604051901515815260200160405180910390f35b341561016a57600080fd5b6101726102d0565b60405190815260200160405180910390f35b341561018f57600080fd5b61014b73ffffffffffffffffffffffffffffffffffffffff600435811690602435166044356102d5565b34156101c457600080fd5b6101726102de565b34156101d757600080fd5b61017273ffffffffffffffffffffffffffffffffffffffff600435166102e4565b341561020357600080fd5b61017273ffffffffffffffffffffffffffffffffffffffff600435811690602435166102c8565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102c05780601f10610295576101008083540402835291602001916102c0565b820191906000526020600020905b8154815290600101906020018083116102a357829003601f168201915b505050505081565b600092915050565b600090565b60009392505050565b60005481565b506000905600a165627a7a72305820ef52f00124bd98860e8fc70ec57d957c1a1060697d0a86c80c5903fe10689fa60029`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Token *TokenSession) Decimals() (*big.Int, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Token *TokenCallerSession) Decimals() (*big.Int, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}
