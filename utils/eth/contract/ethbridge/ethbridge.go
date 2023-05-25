// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridge

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

// EthbridgeMetaData contains all meta data concerning the Ethbridge contract.
var EthbridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ETH_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"externalAddr\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"externalAddr\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"safeMultisigContractAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EthbridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use EthbridgeMetaData.ABI instead.
var EthbridgeABI = EthbridgeMetaData.ABI

// Ethbridge is an auto generated Go binding around an Ethereum contract.
type Ethbridge struct {
	EthbridgeCaller     // Read-only binding to the contract
	EthbridgeTransactor // Write-only binding to the contract
	EthbridgeFilterer   // Log filterer for contract events
}

// EthbridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthbridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthbridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthbridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthbridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthbridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthbridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthbridgeSession struct {
	Contract     *Ethbridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthbridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthbridgeCallerSession struct {
	Contract *EthbridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EthbridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthbridgeTransactorSession struct {
	Contract     *EthbridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthbridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthbridgeRaw struct {
	Contract *Ethbridge // Generic contract binding to access the raw methods on
}

// EthbridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthbridgeCallerRaw struct {
	Contract *EthbridgeCaller // Generic read-only contract binding to access the raw methods on
}

// EthbridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthbridgeTransactorRaw struct {
	Contract *EthbridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthbridge creates a new instance of Ethbridge, bound to a specific deployed contract.
func NewEthbridge(address common.Address, backend bind.ContractBackend) (*Ethbridge, error) {
	contract, err := bindEthbridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethbridge{EthbridgeCaller: EthbridgeCaller{contract: contract}, EthbridgeTransactor: EthbridgeTransactor{contract: contract}, EthbridgeFilterer: EthbridgeFilterer{contract: contract}}, nil
}

// NewEthbridgeCaller creates a new read-only instance of Ethbridge, bound to a specific deployed contract.
func NewEthbridgeCaller(address common.Address, caller bind.ContractCaller) (*EthbridgeCaller, error) {
	contract, err := bindEthbridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthbridgeCaller{contract: contract}, nil
}

// NewEthbridgeTransactor creates a new write-only instance of Ethbridge, bound to a specific deployed contract.
func NewEthbridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*EthbridgeTransactor, error) {
	contract, err := bindEthbridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthbridgeTransactor{contract: contract}, nil
}

// NewEthbridgeFilterer creates a new log filterer instance of Ethbridge, bound to a specific deployed contract.
func NewEthbridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*EthbridgeFilterer, error) {
	contract, err := bindEthbridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthbridgeFilterer{contract: contract}, nil
}

// bindEthbridge binds a generic wrapper to an already deployed contract.
func bindEthbridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthbridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethbridge *EthbridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethbridge.Contract.EthbridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethbridge *EthbridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethbridge.Contract.EthbridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethbridge *EthbridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethbridge.Contract.EthbridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethbridge *EthbridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethbridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethbridge *EthbridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethbridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethbridge *EthbridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethbridge.Contract.contract.Transact(opts, method, params...)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Ethbridge *EthbridgeCaller) ETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethbridge.contract.Call(opts, &out, "ETH_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Ethbridge *EthbridgeSession) ETHTOKEN() (common.Address, error) {
	return _Ethbridge.Contract.ETHTOKEN(&_Ethbridge.CallOpts)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Ethbridge *EthbridgeCallerSession) ETHTOKEN() (common.Address, error) {
	return _Ethbridge.Contract.ETHTOKEN(&_Ethbridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethbridge *EthbridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethbridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethbridge *EthbridgeSession) Owner() (common.Address, error) {
	return _Ethbridge.Contract.Owner(&_Ethbridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethbridge *EthbridgeCallerSession) Owner() (common.Address, error) {
	return _Ethbridge.Contract.Owner(&_Ethbridge.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address externalAddr) payable returns()
func (_Ethbridge *EthbridgeTransactor) Deposit(opts *bind.TransactOpts, externalAddr common.Address) (*types.Transaction, error) {
	return _Ethbridge.contract.Transact(opts, "deposit", externalAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address externalAddr) payable returns()
func (_Ethbridge *EthbridgeSession) Deposit(externalAddr common.Address) (*types.Transaction, error) {
	return _Ethbridge.Contract.Deposit(&_Ethbridge.TransactOpts, externalAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address externalAddr) payable returns()
func (_Ethbridge *EthbridgeTransactorSession) Deposit(externalAddr common.Address) (*types.Transaction, error) {
	return _Ethbridge.Contract.Deposit(&_Ethbridge.TransactOpts, externalAddr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address token, uint256 amount, address externalAddr) returns()
func (_Ethbridge *EthbridgeTransactor) Deposit0(opts *bind.TransactOpts, token common.Address, amount *big.Int, externalAddr common.Address) (*types.Transaction, error) {
	return _Ethbridge.contract.Transact(opts, "deposit0", token, amount, externalAddr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address token, uint256 amount, address externalAddr) returns()
func (_Ethbridge *EthbridgeSession) Deposit0(token common.Address, amount *big.Int, externalAddr common.Address) (*types.Transaction, error) {
	return _Ethbridge.Contract.Deposit0(&_Ethbridge.TransactOpts, token, amount, externalAddr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address token, uint256 amount, address externalAddr) returns()
func (_Ethbridge *EthbridgeTransactorSession) Deposit0(token common.Address, amount *big.Int, externalAddr common.Address) (*types.Transaction, error) {
	return _Ethbridge.Contract.Deposit0(&_Ethbridge.TransactOpts, token, amount, externalAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address safeMultisigContractAddress) returns()
func (_Ethbridge *EthbridgeTransactor) Initialize(opts *bind.TransactOpts, safeMultisigContractAddress common.Address) (*types.Transaction, error) {
	return _Ethbridge.contract.Transact(opts, "initialize", safeMultisigContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address safeMultisigContractAddress) returns()
func (_Ethbridge *EthbridgeSession) Initialize(safeMultisigContractAddress common.Address) (*types.Transaction, error) {
	return _Ethbridge.Contract.Initialize(&_Ethbridge.TransactOpts, safeMultisigContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address safeMultisigContractAddress) returns()
func (_Ethbridge *EthbridgeTransactorSession) Initialize(safeMultisigContractAddress common.Address) (*types.Transaction, error) {
	return _Ethbridge.Contract.Initialize(&_Ethbridge.TransactOpts, safeMultisigContractAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethbridge *EthbridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethbridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethbridge *EthbridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ethbridge.Contract.RenounceOwnership(&_Ethbridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethbridge *EthbridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ethbridge.Contract.RenounceOwnership(&_Ethbridge.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethbridge *EthbridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ethbridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethbridge *EthbridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ethbridge.Contract.TransferOwnership(&_Ethbridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethbridge *EthbridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ethbridge.Contract.TransferOwnership(&_Ethbridge.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xedbd7668.
//
// Solidity: function withdraw(address token, address[] recipients, uint256[] amounts) returns()
func (_Ethbridge *EthbridgeTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Ethbridge.contract.Transact(opts, "withdraw", token, recipients, amounts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xedbd7668.
//
// Solidity: function withdraw(address token, address[] recipients, uint256[] amounts) returns()
func (_Ethbridge *EthbridgeSession) Withdraw(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Ethbridge.Contract.Withdraw(&_Ethbridge.TransactOpts, token, recipients, amounts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xedbd7668.
//
// Solidity: function withdraw(address token, address[] recipients, uint256[] amounts) returns()
func (_Ethbridge *EthbridgeTransactorSession) Withdraw(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Ethbridge.Contract.Withdraw(&_Ethbridge.TransactOpts, token, recipients, amounts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xf7ece0cf.
//
// Solidity: function withdraw(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_Ethbridge *EthbridgeTransactor) Withdraw0(opts *bind.TransactOpts, tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Ethbridge.contract.Transact(opts, "withdraw0", tokens, recipients, amounts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xf7ece0cf.
//
// Solidity: function withdraw(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_Ethbridge *EthbridgeSession) Withdraw0(tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Ethbridge.Contract.Withdraw0(&_Ethbridge.TransactOpts, tokens, recipients, amounts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xf7ece0cf.
//
// Solidity: function withdraw(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_Ethbridge *EthbridgeTransactorSession) Withdraw0(tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Ethbridge.Contract.Withdraw0(&_Ethbridge.TransactOpts, tokens, recipients, amounts)
}

// EthbridgeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Ethbridge contract.
type EthbridgeDepositIterator struct {
	Event *EthbridgeDeposit // Event containing the contract specifics and raw log

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
func (it *EthbridgeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthbridgeDeposit)
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
		it.Event = new(EthbridgeDeposit)
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
func (it *EthbridgeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthbridgeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthbridgeDeposit represents a Deposit event raised by the Ethbridge contract.
type EthbridgeDeposit struct {
	Token     common.Address
	Sender    common.Address
	Amount    *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x364bb76a44233df8584c690de6da7810626a5e77192f3ebc942c35bcb1add24f.
//
// Solidity: event Deposit(address token, address sender, uint256 amount, address recipient)
func (_Ethbridge *EthbridgeFilterer) FilterDeposit(opts *bind.FilterOpts) (*EthbridgeDepositIterator, error) {

	logs, sub, err := _Ethbridge.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &EthbridgeDepositIterator{contract: _Ethbridge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x364bb76a44233df8584c690de6da7810626a5e77192f3ebc942c35bcb1add24f.
//
// Solidity: event Deposit(address token, address sender, uint256 amount, address recipient)
func (_Ethbridge *EthbridgeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EthbridgeDeposit) (event.Subscription, error) {

	logs, sub, err := _Ethbridge.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthbridgeDeposit)
				if err := _Ethbridge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x364bb76a44233df8584c690de6da7810626a5e77192f3ebc942c35bcb1add24f.
//
// Solidity: event Deposit(address token, address sender, uint256 amount, address recipient)
func (_Ethbridge *EthbridgeFilterer) ParseDeposit(log types.Log) (*EthbridgeDeposit, error) {
	event := new(EthbridgeDeposit)
	if err := _Ethbridge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthbridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Ethbridge contract.
type EthbridgeInitializedIterator struct {
	Event *EthbridgeInitialized // Event containing the contract specifics and raw log

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
func (it *EthbridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthbridgeInitialized)
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
		it.Event = new(EthbridgeInitialized)
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
func (it *EthbridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthbridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthbridgeInitialized represents a Initialized event raised by the Ethbridge contract.
type EthbridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ethbridge *EthbridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*EthbridgeInitializedIterator, error) {

	logs, sub, err := _Ethbridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EthbridgeInitializedIterator{contract: _Ethbridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ethbridge *EthbridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EthbridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Ethbridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthbridgeInitialized)
				if err := _Ethbridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ethbridge *EthbridgeFilterer) ParseInitialized(log types.Log) (*EthbridgeInitialized, error) {
	event := new(EthbridgeInitialized)
	if err := _Ethbridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthbridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ethbridge contract.
type EthbridgeOwnershipTransferredIterator struct {
	Event *EthbridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EthbridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthbridgeOwnershipTransferred)
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
		it.Event = new(EthbridgeOwnershipTransferred)
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
func (it *EthbridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthbridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthbridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Ethbridge contract.
type EthbridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ethbridge *EthbridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthbridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ethbridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthbridgeOwnershipTransferredIterator{contract: _Ethbridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ethbridge *EthbridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthbridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ethbridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthbridgeOwnershipTransferred)
				if err := _Ethbridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ethbridge *EthbridgeFilterer) ParseOwnershipTransferred(log types.Log) (*EthbridgeOwnershipTransferred, error) {
	event := new(EthbridgeOwnershipTransferred)
	if err := _Ethbridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthbridgeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Ethbridge contract.
type EthbridgeWithdrawIterator struct {
	Event *EthbridgeWithdraw // Event containing the contract specifics and raw log

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
func (it *EthbridgeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthbridgeWithdraw)
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
		it.Event = new(EthbridgeWithdraw)
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
func (it *EthbridgeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthbridgeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthbridgeWithdraw represents a Withdraw event raised by the Ethbridge contract.
type EthbridgeWithdraw struct {
	Tokens     []common.Address
	Recipients []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xdf8b54622ff1f8c4514f6bd21ae42dd7c0f3111b74c909c7fd9ea7a8b9087b40.
//
// Solidity: event Withdraw(address[] tokens, address[] recipients, uint256[] amounts)
func (_Ethbridge *EthbridgeFilterer) FilterWithdraw(opts *bind.FilterOpts) (*EthbridgeWithdrawIterator, error) {

	logs, sub, err := _Ethbridge.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &EthbridgeWithdrawIterator{contract: _Ethbridge.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xdf8b54622ff1f8c4514f6bd21ae42dd7c0f3111b74c909c7fd9ea7a8b9087b40.
//
// Solidity: event Withdraw(address[] tokens, address[] recipients, uint256[] amounts)
func (_Ethbridge *EthbridgeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *EthbridgeWithdraw) (event.Subscription, error) {

	logs, sub, err := _Ethbridge.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthbridgeWithdraw)
				if err := _Ethbridge.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xdf8b54622ff1f8c4514f6bd21ae42dd7c0f3111b74c909c7fd9ea7a8b9087b40.
//
// Solidity: event Withdraw(address[] tokens, address[] recipients, uint256[] amounts)
func (_Ethbridge *EthbridgeFilterer) ParseWithdraw(log types.Log) (*EthbridgeWithdraw, error) {
	event := new(EthbridgeWithdraw)
	if err := _Ethbridge.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthbridgeWithdraw0Iterator is returned from FilterWithdraw0 and is used to iterate over the raw logs and unpacked data for Withdraw0 events raised by the Ethbridge contract.
type EthbridgeWithdraw0Iterator struct {
	Event *EthbridgeWithdraw0 // Event containing the contract specifics and raw log

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
func (it *EthbridgeWithdraw0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthbridgeWithdraw0)
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
		it.Event = new(EthbridgeWithdraw0)
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
func (it *EthbridgeWithdraw0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthbridgeWithdraw0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthbridgeWithdraw0 represents a Withdraw0 event raised by the Ethbridge contract.
type EthbridgeWithdraw0 struct {
	Token      common.Address
	Recipients []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdraw0 is a free log retrieval operation binding the contract event 0xa7bd5a2454bc906a642fda333278545850c8711673010966530dcabede413367.
//
// Solidity: event Withdraw(address token, address[] recipients, uint256[] amounts)
func (_Ethbridge *EthbridgeFilterer) FilterWithdraw0(opts *bind.FilterOpts) (*EthbridgeWithdraw0Iterator, error) {

	logs, sub, err := _Ethbridge.contract.FilterLogs(opts, "Withdraw0")
	if err != nil {
		return nil, err
	}
	return &EthbridgeWithdraw0Iterator{contract: _Ethbridge.contract, event: "Withdraw0", logs: logs, sub: sub}, nil
}

// WatchWithdraw0 is a free log subscription operation binding the contract event 0xa7bd5a2454bc906a642fda333278545850c8711673010966530dcabede413367.
//
// Solidity: event Withdraw(address token, address[] recipients, uint256[] amounts)
func (_Ethbridge *EthbridgeFilterer) WatchWithdraw0(opts *bind.WatchOpts, sink chan<- *EthbridgeWithdraw0) (event.Subscription, error) {

	logs, sub, err := _Ethbridge.contract.WatchLogs(opts, "Withdraw0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthbridgeWithdraw0)
				if err := _Ethbridge.contract.UnpackLog(event, "Withdraw0", log); err != nil {
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

// ParseWithdraw0 is a log parse operation binding the contract event 0xa7bd5a2454bc906a642fda333278545850c8711673010966530dcabede413367.
//
// Solidity: event Withdraw(address token, address[] recipients, uint256[] amounts)
func (_Ethbridge *EthbridgeFilterer) ParseWithdraw0(log types.Log) (*EthbridgeWithdraw0, error) {
	event := new(EthbridgeWithdraw0)
	if err := _Ethbridge.contract.UnpackLog(event, "Withdraw0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
