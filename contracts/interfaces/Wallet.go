// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"AllowanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MoneyReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MoneySent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"reduceAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b503380603357604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b603a81603f565b50608e565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6106f58061009b5f395ff3fe60806040526004361061009d575f3560e01c8063715018a611610062578063715018a61461017e5780638da5cb5b146101925780638fdf08df146101b85780639b96eece146101d7578063ee4ae2c9146101fe578063f2fde38b1461021d575f5ffd5b806312065fe0146100dd5780632b603c71146100ff578063310ec4a71461012a578063395093511461014b5780636f9fb98a1461016a575f5ffd5b366100d95760405134815233907f27b15ed4cf832749ed39f33a64e4707ed60a761485e41ffec7343ecaddc0c02a9060200160405180910390a2005b5f5ffd5b3480156100e8575f5ffd5b5033315b6040519081526020015b60405180910390f35b34801561010a575f5ffd5b506100ec610119366004610633565b60016020525f908152604090205481565b348015610135575f5ffd5b50610149610144366004610655565b61023c565b005b348015610156575f5ffd5b50610149610165366004610655565b6102b4565b348015610175575f5ffd5b506100ec610355565b348015610189575f5ffd5b50610149610363565b34801561019d575f5ffd5b505f546040516001600160a01b0390911681526020016100f6565b3480156101c3575f5ffd5b506101496101d2366004610655565b610376565b3480156101e2575f5ffd5b506100ec6101f1366004610633565b6001600160a01b03163190565b348015610209575f5ffd5b50610149610218366004610655565b6103a1565b348015610228575f5ffd5b50610149610237366004610633565b610545565b610244610582565b6001600160a01b0382165f818152600160209081526040918290205482519081529081018490523392917f3691d1a86d99355e52b689ca70a7bdf6d80763237a6aa06e5fa43964eac7244b910160405180910390a36001600160a01b039091165f90815260016020526040902055565b6102bc610582565b6001600160a01b0382165f908152600160205260408120546102df908390610693565b6001600160a01b0384165f8181526001602090815260409182902054825190815290810184905292935033927f3691d1a86d99355e52b689ca70a7bdf6d80763237a6aa06e5fa43964eac7244b910160405180910390a36001600160a01b039092165f9081526001602052604090209190915550565b5f61035e610582565b504790565b61036b610582565b6103745f6105ae565b565b61037e610582565b6001600160a01b0382165f908152600160205260408120546102df9083906106ac565b6103a96105fd565b6103fa5760405162461bcd60e51b815260206004820152601a60248201527f556e617574686f72697a656420746f2073656e64206d6f6e657900000000000060448201526064015b60405180910390fd5b8047101561044a5760405162461bcd60e51b815260206004820152601d60248201527f496e73756666696369656e7420636f6e74726163742062616c616e636500000060448201526064016103f1565b6001600160a01b0382165f908152600160205260409020548111156104c05760405162461bcd60e51b815260206004820152602660248201527f496e73756666696369656e7420616c6c6f77616e636520666f72207472616e7360448201526530b1ba34b7b760d11b60648201526084016103f1565b6104ca8282610376565b6040516001600160a01b0383169082156108fc029083905f818181858888f193505050501580156104fd573d5f5f3e3d5ffd5b50816001600160a01b03167f7f51d406915971d4ac1c91af96be5187ea6ab64753785aad519a533def80a41e8260405161053991815260200190565b60405180910390a25050565b61054d610582565b6001600160a01b03811661057657604051631e4fbdf760e01b81525f60048201526024016103f1565b61057f816105ae565b50565b5f546001600160a01b031633146103745760405163118cdaa760e01b81523360048201526024016103f1565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f336106105f546001600160a01b031690565b6001600160a01b031614905090565b6001600160a01b038116811461057f575f5ffd5b5f60208284031215610643575f5ffd5b813561064e8161061f565b9392505050565b5f5f60408385031215610666575f5ffd5b82356106718161061f565b946020939093013593505050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156106a6576106a661067f565b92915050565b818103818111156106a6576106a661067f56fea26469706673582212207780a1dba4d262d557a08f1c04bfd6fc8fa1205edf82d7def4cd3fcff1e83e5d64736f6c634300081c0033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// Allowances is a free data retrieval call binding the contract method 0x2b603c71.
//
// Solidity: function allowances(address ) view returns(uint256)
func (_Contracts *ContractsCaller) Allowances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allowances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x2b603c71.
//
// Solidity: function allowances(address ) view returns(uint256)
func (_Contracts *ContractsSession) Allowances(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowances(&_Contracts.CallOpts, arg0)
}

// Allowances is a free data retrieval call binding the contract method 0x2b603c71.
//
// Solidity: function allowances(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) Allowances(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowances(&_Contracts.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contracts *ContractsCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contracts *ContractsSession) GetBalance() (*big.Int, error) {
	return _Contracts.Contract.GetBalance(&_Contracts.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contracts *ContractsCallerSession) GetBalance() (*big.Int, error) {
	return _Contracts.Contract.GetBalance(&_Contracts.CallOpts)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address _address) view returns(uint256)
func (_Contracts *ContractsCaller) GetBalanceOf(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getBalanceOf", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address _address) view returns(uint256)
func (_Contracts *ContractsSession) GetBalanceOf(_address common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetBalanceOf(&_Contracts.CallOpts, _address)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address _address) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetBalanceOf(_address common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetBalanceOf(&_Contracts.CallOpts, _address)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Contracts *ContractsCaller) GetContractBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getContractBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Contracts *ContractsSession) GetContractBalance() (*big.Int, error) {
	return _Contracts.Contract.GetContractBalance(&_Contracts.CallOpts)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Contracts *ContractsCallerSession) GetContractBalance() (*big.Int, error) {
	return _Contracts.Contract.GetContractBalance(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contracts *ContractsTransactor) IncreaseAllowance(opts *bind.TransactOpts, _beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "increaseAllowance", _beneficiary, _amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contracts *ContractsSession) IncreaseAllowance(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.IncreaseAllowance(&_Contracts.TransactOpts, _beneficiary, _amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contracts *ContractsTransactorSession) IncreaseAllowance(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.IncreaseAllowance(&_Contracts.TransactOpts, _beneficiary, _amount)
}

// ReduceAllowance is a paid mutator transaction binding the contract method 0x8fdf08df.
//
// Solidity: function reduceAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contracts *ContractsTransactor) ReduceAllowance(opts *bind.TransactOpts, _beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "reduceAllowance", _beneficiary, _amount)
}

// ReduceAllowance is a paid mutator transaction binding the contract method 0x8fdf08df.
//
// Solidity: function reduceAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contracts *ContractsSession) ReduceAllowance(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ReduceAllowance(&_Contracts.TransactOpts, _beneficiary, _amount)
}

// ReduceAllowance is a paid mutator transaction binding the contract method 0x8fdf08df.
//
// Solidity: function reduceAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contracts *ContractsTransactorSession) ReduceAllowance(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ReduceAllowance(&_Contracts.TransactOpts, _beneficiary, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// SendMoney is a paid mutator transaction binding the contract method 0xee4ae2c9.
//
// Solidity: function sendMoney(address _to, uint256 _amount) returns()
func (_Contracts *ContractsTransactor) SendMoney(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "sendMoney", _to, _amount)
}

// SendMoney is a paid mutator transaction binding the contract method 0xee4ae2c9.
//
// Solidity: function sendMoney(address _to, uint256 _amount) returns()
func (_Contracts *ContractsSession) SendMoney(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SendMoney(&_Contracts.TransactOpts, _to, _amount)
}

// SendMoney is a paid mutator transaction binding the contract method 0xee4ae2c9.
//
// Solidity: function sendMoney(address _to, uint256 _amount) returns()
func (_Contracts *ContractsTransactorSession) SendMoney(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SendMoney(&_Contracts.TransactOpts, _to, _amount)
}

// SetAllowance is a paid mutator transaction binding the contract method 0x310ec4a7.
//
// Solidity: function setAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contracts *ContractsTransactor) SetAllowance(opts *bind.TransactOpts, _beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setAllowance", _beneficiary, _amount)
}

// SetAllowance is a paid mutator transaction binding the contract method 0x310ec4a7.
//
// Solidity: function setAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contracts *ContractsSession) SetAllowance(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetAllowance(&_Contracts.TransactOpts, _beneficiary, _amount)
}

// SetAllowance is a paid mutator transaction binding the contract method 0x310ec4a7.
//
// Solidity: function setAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contracts *ContractsTransactorSession) SetAllowance(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetAllowance(&_Contracts.TransactOpts, _beneficiary, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactorSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// ContractsAllowanceChangedIterator is returned from FilterAllowanceChanged and is used to iterate over the raw logs and unpacked data for AllowanceChanged events raised by the Contracts contract.
type ContractsAllowanceChangedIterator struct {
	Event *ContractsAllowanceChanged // Event containing the contract specifics and raw log

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
func (it *ContractsAllowanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAllowanceChanged)
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
		it.Event = new(ContractsAllowanceChanged)
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
func (it *ContractsAllowanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAllowanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAllowanceChanged represents a AllowanceChanged event raised by the Contracts contract.
type ContractsAllowanceChanged struct {
	Beneficiary common.Address
	Sender      common.Address
	PrevAmount  *big.Int
	NewAmount   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllowanceChanged is a free log retrieval operation binding the contract event 0x3691d1a86d99355e52b689ca70a7bdf6d80763237a6aa06e5fa43964eac7244b.
//
// Solidity: event AllowanceChanged(address indexed beneficiary, address indexed sender, uint256 prevAmount, uint256 newAmount)
func (_Contracts *ContractsFilterer) FilterAllowanceChanged(opts *bind.FilterOpts, beneficiary []common.Address, sender []common.Address) (*ContractsAllowanceChangedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AllowanceChanged", beneficiaryRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAllowanceChangedIterator{contract: _Contracts.contract, event: "AllowanceChanged", logs: logs, sub: sub}, nil
}

// WatchAllowanceChanged is a free log subscription operation binding the contract event 0x3691d1a86d99355e52b689ca70a7bdf6d80763237a6aa06e5fa43964eac7244b.
//
// Solidity: event AllowanceChanged(address indexed beneficiary, address indexed sender, uint256 prevAmount, uint256 newAmount)
func (_Contracts *ContractsFilterer) WatchAllowanceChanged(opts *bind.WatchOpts, sink chan<- *ContractsAllowanceChanged, beneficiary []common.Address, sender []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AllowanceChanged", beneficiaryRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAllowanceChanged)
				if err := _Contracts.contract.UnpackLog(event, "AllowanceChanged", log); err != nil {
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

// ParseAllowanceChanged is a log parse operation binding the contract event 0x3691d1a86d99355e52b689ca70a7bdf6d80763237a6aa06e5fa43964eac7244b.
//
// Solidity: event AllowanceChanged(address indexed beneficiary, address indexed sender, uint256 prevAmount, uint256 newAmount)
func (_Contracts *ContractsFilterer) ParseAllowanceChanged(log types.Log) (*ContractsAllowanceChanged, error) {
	event := new(ContractsAllowanceChanged)
	if err := _Contracts.contract.UnpackLog(event, "AllowanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMoneyReceivedIterator is returned from FilterMoneyReceived and is used to iterate over the raw logs and unpacked data for MoneyReceived events raised by the Contracts contract.
type ContractsMoneyReceivedIterator struct {
	Event *ContractsMoneyReceived // Event containing the contract specifics and raw log

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
func (it *ContractsMoneyReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMoneyReceived)
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
		it.Event = new(ContractsMoneyReceived)
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
func (it *ContractsMoneyReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMoneyReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMoneyReceived represents a MoneyReceived event raised by the Contracts contract.
type ContractsMoneyReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMoneyReceived is a free log retrieval operation binding the contract event 0x27b15ed4cf832749ed39f33a64e4707ed60a761485e41ffec7343ecaddc0c02a.
//
// Solidity: event MoneyReceived(address indexed sender, uint256 amount)
func (_Contracts *ContractsFilterer) FilterMoneyReceived(opts *bind.FilterOpts, sender []common.Address) (*ContractsMoneyReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MoneyReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMoneyReceivedIterator{contract: _Contracts.contract, event: "MoneyReceived", logs: logs, sub: sub}, nil
}

// WatchMoneyReceived is a free log subscription operation binding the contract event 0x27b15ed4cf832749ed39f33a64e4707ed60a761485e41ffec7343ecaddc0c02a.
//
// Solidity: event MoneyReceived(address indexed sender, uint256 amount)
func (_Contracts *ContractsFilterer) WatchMoneyReceived(opts *bind.WatchOpts, sink chan<- *ContractsMoneyReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MoneyReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMoneyReceived)
				if err := _Contracts.contract.UnpackLog(event, "MoneyReceived", log); err != nil {
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

// ParseMoneyReceived is a log parse operation binding the contract event 0x27b15ed4cf832749ed39f33a64e4707ed60a761485e41ffec7343ecaddc0c02a.
//
// Solidity: event MoneyReceived(address indexed sender, uint256 amount)
func (_Contracts *ContractsFilterer) ParseMoneyReceived(log types.Log) (*ContractsMoneyReceived, error) {
	event := new(ContractsMoneyReceived)
	if err := _Contracts.contract.UnpackLog(event, "MoneyReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMoneySentIterator is returned from FilterMoneySent and is used to iterate over the raw logs and unpacked data for MoneySent events raised by the Contracts contract.
type ContractsMoneySentIterator struct {
	Event *ContractsMoneySent // Event containing the contract specifics and raw log

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
func (it *ContractsMoneySentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMoneySent)
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
		it.Event = new(ContractsMoneySent)
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
func (it *ContractsMoneySentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMoneySentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMoneySent represents a MoneySent event raised by the Contracts contract.
type ContractsMoneySent struct {
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMoneySent is a free log retrieval operation binding the contract event 0x7f51d406915971d4ac1c91af96be5187ea6ab64753785aad519a533def80a41e.
//
// Solidity: event MoneySent(address indexed beneficiary, uint256 amount)
func (_Contracts *ContractsFilterer) FilterMoneySent(opts *bind.FilterOpts, beneficiary []common.Address) (*ContractsMoneySentIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MoneySent", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMoneySentIterator{contract: _Contracts.contract, event: "MoneySent", logs: logs, sub: sub}, nil
}

// WatchMoneySent is a free log subscription operation binding the contract event 0x7f51d406915971d4ac1c91af96be5187ea6ab64753785aad519a533def80a41e.
//
// Solidity: event MoneySent(address indexed beneficiary, uint256 amount)
func (_Contracts *ContractsFilterer) WatchMoneySent(opts *bind.WatchOpts, sink chan<- *ContractsMoneySent, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MoneySent", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMoneySent)
				if err := _Contracts.contract.UnpackLog(event, "MoneySent", log); err != nil {
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

// ParseMoneySent is a log parse operation binding the contract event 0x7f51d406915971d4ac1c91af96be5187ea6ab64753785aad519a533def80a41e.
//
// Solidity: event MoneySent(address indexed beneficiary, uint256 amount)
func (_Contracts *ContractsFilterer) ParseMoneySent(log types.Log) (*ContractsMoneySent, error) {
	event := new(ContractsMoneySent)
	if err := _Contracts.contract.UnpackLog(event, "MoneySent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
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
		it.Event = new(ContractsOwnershipTransferred)
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
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
