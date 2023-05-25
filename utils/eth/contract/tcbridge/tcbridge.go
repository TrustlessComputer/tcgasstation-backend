// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tcbridge

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

// TcbridgeMetaData contains all meta data concerning the Tcbridge contract.
var TcbridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractWrappedToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcAddr\",\"type\":\"string\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractWrappedToken[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractWrappedToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractWrappedToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"externalAddr\",\"type\":\"string\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractWrappedToken[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractWrappedToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TcbridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use TcbridgeMetaData.ABI instead.
var TcbridgeABI = TcbridgeMetaData.ABI

// Tcbridge is an auto generated Go binding around an Ethereum contract.
type Tcbridge struct {
	TcbridgeCaller     // Read-only binding to the contract
	TcbridgeTransactor // Write-only binding to the contract
	TcbridgeFilterer   // Log filterer for contract events
}

// TcbridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TcbridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TcbridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TcbridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TcbridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TcbridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TcbridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TcbridgeSession struct {
	Contract     *Tcbridge         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TcbridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TcbridgeCallerSession struct {
	Contract *TcbridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TcbridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TcbridgeTransactorSession struct {
	Contract     *TcbridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TcbridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TcbridgeRaw struct {
	Contract *Tcbridge // Generic contract binding to access the raw methods on
}

// TcbridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TcbridgeCallerRaw struct {
	Contract *TcbridgeCaller // Generic read-only contract binding to access the raw methods on
}

// TcbridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TcbridgeTransactorRaw struct {
	Contract *TcbridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTcbridge creates a new instance of Tcbridge, bound to a specific deployed contract.
func NewTcbridge(address common.Address, backend bind.ContractBackend) (*Tcbridge, error) {
	contract, err := bindTcbridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tcbridge{TcbridgeCaller: TcbridgeCaller{contract: contract}, TcbridgeTransactor: TcbridgeTransactor{contract: contract}, TcbridgeFilterer: TcbridgeFilterer{contract: contract}}, nil
}

// NewTcbridgeCaller creates a new read-only instance of Tcbridge, bound to a specific deployed contract.
func NewTcbridgeCaller(address common.Address, caller bind.ContractCaller) (*TcbridgeCaller, error) {
	contract, err := bindTcbridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TcbridgeCaller{contract: contract}, nil
}

// NewTcbridgeTransactor creates a new write-only instance of Tcbridge, bound to a specific deployed contract.
func NewTcbridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*TcbridgeTransactor, error) {
	contract, err := bindTcbridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TcbridgeTransactor{contract: contract}, nil
}

// NewTcbridgeFilterer creates a new log filterer instance of Tcbridge, bound to a specific deployed contract.
func NewTcbridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*TcbridgeFilterer, error) {
	contract, err := bindTcbridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TcbridgeFilterer{contract: contract}, nil
}

// bindTcbridge binds a generic wrapper to an already deployed contract.
func bindTcbridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TcbridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tcbridge *TcbridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tcbridge.Contract.TcbridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tcbridge *TcbridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tcbridge.Contract.TcbridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tcbridge *TcbridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tcbridge.Contract.TcbridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tcbridge *TcbridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tcbridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tcbridge *TcbridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tcbridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tcbridge *TcbridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tcbridge.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tcbridge *TcbridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tcbridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tcbridge *TcbridgeSession) Owner() (common.Address, error) {
	return _Tcbridge.Contract.Owner(&_Tcbridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tcbridge *TcbridgeCallerSession) Owner() (common.Address, error) {
	return _Tcbridge.Contract.Owner(&_Tcbridge.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x15f570dc.
//
// Solidity: function burn(address token, uint256 amount, string externalAddr) returns()
func (_Tcbridge *TcbridgeTransactor) Burn(opts *bind.TransactOpts, token common.Address, amount *big.Int, externalAddr string) (*types.Transaction, error) {
	return _Tcbridge.contract.Transact(opts, "burn", token, amount, externalAddr)
}

// Burn is a paid mutator transaction binding the contract method 0x15f570dc.
//
// Solidity: function burn(address token, uint256 amount, string externalAddr) returns()
func (_Tcbridge *TcbridgeSession) Burn(token common.Address, amount *big.Int, externalAddr string) (*types.Transaction, error) {
	return _Tcbridge.Contract.Burn(&_Tcbridge.TransactOpts, token, amount, externalAddr)
}

// Burn is a paid mutator transaction binding the contract method 0x15f570dc.
//
// Solidity: function burn(address token, uint256 amount, string externalAddr) returns()
func (_Tcbridge *TcbridgeTransactorSession) Burn(token common.Address, amount *big.Int, externalAddr string) (*types.Transaction, error) {
	return _Tcbridge.Contract.Burn(&_Tcbridge.TransactOpts, token, amount, externalAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address validator) returns()
func (_Tcbridge *TcbridgeTransactor) Initialize(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Tcbridge.contract.Transact(opts, "initialize", validator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address validator) returns()
func (_Tcbridge *TcbridgeSession) Initialize(validator common.Address) (*types.Transaction, error) {
	return _Tcbridge.Contract.Initialize(&_Tcbridge.TransactOpts, validator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address validator) returns()
func (_Tcbridge *TcbridgeTransactorSession) Initialize(validator common.Address) (*types.Transaction, error) {
	return _Tcbridge.Contract.Initialize(&_Tcbridge.TransactOpts, validator)
}

// Mint is a paid mutator transaction binding the contract method 0x5530f4a5.
//
// Solidity: function mint(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_Tcbridge *TcbridgeTransactor) Mint(opts *bind.TransactOpts, tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Tcbridge.contract.Transact(opts, "mint", tokens, recipients, amounts)
}

// Mint is a paid mutator transaction binding the contract method 0x5530f4a5.
//
// Solidity: function mint(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_Tcbridge *TcbridgeSession) Mint(tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Tcbridge.Contract.Mint(&_Tcbridge.TransactOpts, tokens, recipients, amounts)
}

// Mint is a paid mutator transaction binding the contract method 0x5530f4a5.
//
// Solidity: function mint(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_Tcbridge *TcbridgeTransactorSession) Mint(tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Tcbridge.Contract.Mint(&_Tcbridge.TransactOpts, tokens, recipients, amounts)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa3bf277e.
//
// Solidity: function mint(address token, address[] recipients, uint256[] amounts) returns()
func (_Tcbridge *TcbridgeTransactor) Mint0(opts *bind.TransactOpts, token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Tcbridge.contract.Transact(opts, "mint0", token, recipients, amounts)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa3bf277e.
//
// Solidity: function mint(address token, address[] recipients, uint256[] amounts) returns()
func (_Tcbridge *TcbridgeSession) Mint0(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Tcbridge.Contract.Mint0(&_Tcbridge.TransactOpts, token, recipients, amounts)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa3bf277e.
//
// Solidity: function mint(address token, address[] recipients, uint256[] amounts) returns()
func (_Tcbridge *TcbridgeTransactorSession) Mint0(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Tcbridge.Contract.Mint0(&_Tcbridge.TransactOpts, token, recipients, amounts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tcbridge *TcbridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tcbridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tcbridge *TcbridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tcbridge.Contract.RenounceOwnership(&_Tcbridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tcbridge *TcbridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tcbridge.Contract.RenounceOwnership(&_Tcbridge.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tcbridge *TcbridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tcbridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tcbridge *TcbridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tcbridge.Contract.TransferOwnership(&_Tcbridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tcbridge *TcbridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tcbridge.Contract.TransferOwnership(&_Tcbridge.TransactOpts, newOwner)
}

// TcbridgeBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Tcbridge contract.
type TcbridgeBurnIterator struct {
	Event *TcbridgeBurn // Event containing the contract specifics and raw log

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
func (it *TcbridgeBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TcbridgeBurn)
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
		it.Event = new(TcbridgeBurn)
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
func (it *TcbridgeBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TcbridgeBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TcbridgeBurn represents a Burn event raised by the Tcbridge contract.
type TcbridgeBurn struct {
	Token   common.Address
	Burner  common.Address
	Amount  *big.Int
	BtcAddr string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xeab88a0b2198f2928ade5fb787115a9a6ffbbf3705277143953a7c26769157ff.
//
// Solidity: event Burn(address token, address burner, uint256 amount, string btcAddr)
func (_Tcbridge *TcbridgeFilterer) FilterBurn(opts *bind.FilterOpts) (*TcbridgeBurnIterator, error) {

	logs, sub, err := _Tcbridge.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &TcbridgeBurnIterator{contract: _Tcbridge.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xeab88a0b2198f2928ade5fb787115a9a6ffbbf3705277143953a7c26769157ff.
//
// Solidity: event Burn(address token, address burner, uint256 amount, string btcAddr)
func (_Tcbridge *TcbridgeFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *TcbridgeBurn) (event.Subscription, error) {

	logs, sub, err := _Tcbridge.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TcbridgeBurn)
				if err := _Tcbridge.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xeab88a0b2198f2928ade5fb787115a9a6ffbbf3705277143953a7c26769157ff.
//
// Solidity: event Burn(address token, address burner, uint256 amount, string btcAddr)
func (_Tcbridge *TcbridgeFilterer) ParseBurn(log types.Log) (*TcbridgeBurn, error) {
	event := new(TcbridgeBurn)
	if err := _Tcbridge.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TcbridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Tcbridge contract.
type TcbridgeInitializedIterator struct {
	Event *TcbridgeInitialized // Event containing the contract specifics and raw log

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
func (it *TcbridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TcbridgeInitialized)
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
		it.Event = new(TcbridgeInitialized)
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
func (it *TcbridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TcbridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TcbridgeInitialized represents a Initialized event raised by the Tcbridge contract.
type TcbridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tcbridge *TcbridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*TcbridgeInitializedIterator, error) {

	logs, sub, err := _Tcbridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TcbridgeInitializedIterator{contract: _Tcbridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tcbridge *TcbridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TcbridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Tcbridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TcbridgeInitialized)
				if err := _Tcbridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Tcbridge *TcbridgeFilterer) ParseInitialized(log types.Log) (*TcbridgeInitialized, error) {
	event := new(TcbridgeInitialized)
	if err := _Tcbridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TcbridgeMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Tcbridge contract.
type TcbridgeMintIterator struct {
	Event *TcbridgeMint // Event containing the contract specifics and raw log

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
func (it *TcbridgeMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TcbridgeMint)
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
		it.Event = new(TcbridgeMint)
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
func (it *TcbridgeMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TcbridgeMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TcbridgeMint represents a Mint event raised by the Tcbridge contract.
type TcbridgeMint struct {
	Tokens     []common.Address
	Recipients []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xe9914506df53b6ba40090fea5ed4edb71623a51062de3125c2dc65b23de6d05e.
//
// Solidity: event Mint(address[] tokens, address[] recipients, uint256[] amounts)
func (_Tcbridge *TcbridgeFilterer) FilterMint(opts *bind.FilterOpts) (*TcbridgeMintIterator, error) {

	logs, sub, err := _Tcbridge.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &TcbridgeMintIterator{contract: _Tcbridge.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xe9914506df53b6ba40090fea5ed4edb71623a51062de3125c2dc65b23de6d05e.
//
// Solidity: event Mint(address[] tokens, address[] recipients, uint256[] amounts)
func (_Tcbridge *TcbridgeFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *TcbridgeMint) (event.Subscription, error) {

	logs, sub, err := _Tcbridge.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TcbridgeMint)
				if err := _Tcbridge.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xe9914506df53b6ba40090fea5ed4edb71623a51062de3125c2dc65b23de6d05e.
//
// Solidity: event Mint(address[] tokens, address[] recipients, uint256[] amounts)
func (_Tcbridge *TcbridgeFilterer) ParseMint(log types.Log) (*TcbridgeMint, error) {
	event := new(TcbridgeMint)
	if err := _Tcbridge.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TcbridgeMint0Iterator is returned from FilterMint0 and is used to iterate over the raw logs and unpacked data for Mint0 events raised by the Tcbridge contract.
type TcbridgeMint0Iterator struct {
	Event *TcbridgeMint0 // Event containing the contract specifics and raw log

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
func (it *TcbridgeMint0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TcbridgeMint0)
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
		it.Event = new(TcbridgeMint0)
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
func (it *TcbridgeMint0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TcbridgeMint0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TcbridgeMint0 represents a Mint0 event raised by the Tcbridge contract.
type TcbridgeMint0 struct {
	Token      common.Address
	Recipients []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint0 is a free log retrieval operation binding the contract event 0xa20ca4d8d83b89ff090c0ea7b3c3c600625d46681874e0c0d1e35a1d1d4964dd.
//
// Solidity: event Mint(address token, address[] recipients, uint256[] amounts)
func (_Tcbridge *TcbridgeFilterer) FilterMint0(opts *bind.FilterOpts) (*TcbridgeMint0Iterator, error) {

	logs, sub, err := _Tcbridge.contract.FilterLogs(opts, "Mint0")
	if err != nil {
		return nil, err
	}
	return &TcbridgeMint0Iterator{contract: _Tcbridge.contract, event: "Mint0", logs: logs, sub: sub}, nil
}

// WatchMint0 is a free log subscription operation binding the contract event 0xa20ca4d8d83b89ff090c0ea7b3c3c600625d46681874e0c0d1e35a1d1d4964dd.
//
// Solidity: event Mint(address token, address[] recipients, uint256[] amounts)
func (_Tcbridge *TcbridgeFilterer) WatchMint0(opts *bind.WatchOpts, sink chan<- *TcbridgeMint0) (event.Subscription, error) {

	logs, sub, err := _Tcbridge.contract.WatchLogs(opts, "Mint0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TcbridgeMint0)
				if err := _Tcbridge.contract.UnpackLog(event, "Mint0", log); err != nil {
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

// ParseMint0 is a log parse operation binding the contract event 0xa20ca4d8d83b89ff090c0ea7b3c3c600625d46681874e0c0d1e35a1d1d4964dd.
//
// Solidity: event Mint(address token, address[] recipients, uint256[] amounts)
func (_Tcbridge *TcbridgeFilterer) ParseMint0(log types.Log) (*TcbridgeMint0, error) {
	event := new(TcbridgeMint0)
	if err := _Tcbridge.contract.UnpackLog(event, "Mint0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TcbridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tcbridge contract.
type TcbridgeOwnershipTransferredIterator struct {
	Event *TcbridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TcbridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TcbridgeOwnershipTransferred)
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
		it.Event = new(TcbridgeOwnershipTransferred)
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
func (it *TcbridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TcbridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TcbridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Tcbridge contract.
type TcbridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tcbridge *TcbridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TcbridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tcbridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TcbridgeOwnershipTransferredIterator{contract: _Tcbridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tcbridge *TcbridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TcbridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tcbridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TcbridgeOwnershipTransferred)
				if err := _Tcbridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Tcbridge *TcbridgeFilterer) ParseOwnershipTransferred(log types.Log) (*TcbridgeOwnershipTransferred, error) {
	event := new(TcbridgeOwnershipTransferred)
	if err := _Tcbridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
