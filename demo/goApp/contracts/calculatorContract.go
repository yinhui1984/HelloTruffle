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
)

// CalculatorMetaData contains all meta data concerning the Calculator contract.
var CalculatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"subtract\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a5f3c23b": "add(int256,int256)",
		"b93ea812": "subtract(int256,int256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610180806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063a5f3c23b1461003b578063b93ea81214610060575b600080fd5b61004e610049366004610092565b610073565b60405190815260200160405180910390f35b61004e61006e366004610092565b610086565b600061007f82846100ca565b9392505050565b600061007f828461010b565b600080604083850312156100a557600080fd5b50508035926020909101359150565b634e487b7160e01b600052601160045260246000fd5b600080821280156001600160ff1b03849003851316156100ec576100ec6100b4565b600160ff1b8390038412811615610105576101056100b4565b50500190565b60008083128015600160ff1b850184121615610129576101296100b4565b6001600160ff1b0384018313811615610144576101446100b4565b5050039056fea2646970667358221220a843658ed5041d0eb9dacfc47d69d772a8f49bbeaf9828e1455c767aa0c493fd64736f6c634300080d0033",
}

// CalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use CalculatorMetaData.ABI instead.
var CalculatorABI = CalculatorMetaData.ABI

// Deprecated: Use CalculatorMetaData.Sigs instead.
// CalculatorFuncSigs maps the 4-byte function signature to its string representation.
var CalculatorFuncSigs = CalculatorMetaData.Sigs

// CalculatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CalculatorMetaData.Bin instead.
var CalculatorBin = CalculatorMetaData.Bin

// DeployCalculator deploys a new Ethereum contract, binding an instance of Calculator to it.
func DeployCalculator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Calculator, error) {
	parsed, err := CalculatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CalculatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Calculator{CalculatorCaller: CalculatorCaller{contract: contract}, CalculatorTransactor: CalculatorTransactor{contract: contract}, CalculatorFilterer: CalculatorFilterer{contract: contract}}, nil
}

// Calculator is an auto generated Go binding around an Ethereum contract.
type Calculator struct {
	CalculatorCaller     // Read-only binding to the contract
	CalculatorTransactor // Write-only binding to the contract
	CalculatorFilterer   // Log filterer for contract events
}

// CalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CalculatorSession struct {
	Contract     *Calculator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CalculatorCallerSession struct {
	Contract *CalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CalculatorTransactorSession struct {
	Contract     *CalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CalculatorRaw struct {
	Contract *Calculator // Generic contract binding to access the raw methods on
}

// CalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CalculatorCallerRaw struct {
	Contract *CalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// CalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CalculatorTransactorRaw struct {
	Contract *CalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCalculator creates a new instance of Calculator, bound to a specific deployed contract.
func NewCalculator(address common.Address, backend bind.ContractBackend) (*Calculator, error) {
	contract, err := bindCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Calculator{CalculatorCaller: CalculatorCaller{contract: contract}, CalculatorTransactor: CalculatorTransactor{contract: contract}, CalculatorFilterer: CalculatorFilterer{contract: contract}}, nil
}

// NewCalculatorCaller creates a new read-only instance of Calculator, bound to a specific deployed contract.
func NewCalculatorCaller(address common.Address, caller bind.ContractCaller) (*CalculatorCaller, error) {
	contract, err := bindCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CalculatorCaller{contract: contract}, nil
}

// NewCalculatorTransactor creates a new write-only instance of Calculator, bound to a specific deployed contract.
func NewCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*CalculatorTransactor, error) {
	contract, err := bindCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CalculatorTransactor{contract: contract}, nil
}

// NewCalculatorFilterer creates a new log filterer instance of Calculator, bound to a specific deployed contract.
func NewCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*CalculatorFilterer, error) {
	contract, err := bindCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CalculatorFilterer{contract: contract}, nil
}

// bindCalculator binds a generic wrapper to an already deployed contract.
func bindCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CalculatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Calculator *CalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Calculator.Contract.CalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Calculator *CalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Calculator.Contract.CalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Calculator *CalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Calculator.Contract.CalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Calculator *CalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Calculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Calculator *CalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Calculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Calculator *CalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Calculator.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0xa5f3c23b.
//
// Solidity: function add(int256 a, int256 b) pure returns(int256)
func (_Calculator *CalculatorCaller) Add(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Calculator.contract.Call(opts, &out, "add", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Add is a free data retrieval call binding the contract method 0xa5f3c23b.
//
// Solidity: function add(int256 a, int256 b) pure returns(int256)
func (_Calculator *CalculatorSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Calculator.Contract.Add(&_Calculator.CallOpts, a, b)
}

// Add is a free data retrieval call binding the contract method 0xa5f3c23b.
//
// Solidity: function add(int256 a, int256 b) pure returns(int256)
func (_Calculator *CalculatorCallerSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Calculator.Contract.Add(&_Calculator.CallOpts, a, b)
}

// Subtract is a free data retrieval call binding the contract method 0xb93ea812.
//
// Solidity: function subtract(int256 a, int256 b) pure returns(int256)
func (_Calculator *CalculatorCaller) Subtract(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Calculator.contract.Call(opts, &out, "subtract", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Subtract is a free data retrieval call binding the contract method 0xb93ea812.
//
// Solidity: function subtract(int256 a, int256 b) pure returns(int256)
func (_Calculator *CalculatorSession) Subtract(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Calculator.Contract.Subtract(&_Calculator.CallOpts, a, b)
}

// Subtract is a free data retrieval call binding the contract method 0xb93ea812.
//
// Solidity: function subtract(int256 a, int256 b) pure returns(int256)
func (_Calculator *CalculatorCallerSession) Subtract(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Calculator.Contract.Subtract(&_Calculator.CallOpts, a, b)
}
