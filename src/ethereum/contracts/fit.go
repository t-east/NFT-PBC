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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spa\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"FIT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fitId\",\"type\":\"uint256\"}],\"name\":\"getHashedFile\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fitId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"registerDedUpData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileData\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"registerOriginalData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_w\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405162000c5d38038062000c5d83398101604081905261003191610117565b600180546001600160a01b0319166001600160a01b038316179055604080518082019091526004808252631d195cdd60e21b60209092019182526100779160029161007e565b5050610182565b82805461008a90610147565b90600052602060002090601f0160209004810192826100ac57600085556100f2565b82601f106100c557805160ff19168380011785556100f2565b828001600101855582156100f2579182015b828111156100f25782518255916020019190600101906100d7565b506100fe929150610102565b5090565b5b808211156100fe5760008155600101610103565b60006020828403121561012957600080fd5b81516001600160a01b038116811461014057600080fd5b9392505050565b600181811c9082168061015b57607f821691505b6020821081141561017c57634e487b7160e01b600052602260045260246000fd5b50919050565b610acb80620001926000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80631040d067146100675780634ed3885e1461007c5780635334ba1f146100a55780635e2b582a146100c65780636d4ce63c146100e6578063f1573dfb146100ee575b600080fd5b61007a61007536600461075b565b610124565b005b61008f61008a366004610826565b610203565b60405161009c91906108c4565b60405180910390f35b6100b86100b33660046108de565b6102ac565b60405190815260200161009c565b6100d96100d43660046109c9565b61047b565b60405161009c91906109e2565b61008f610561565b6101146100fc3660046109c9565b60006020819052908152604090206002015460ff1681565b604051901515815260200161009c565b6001546001600160a01b031633146101745760405162461bcd60e51b815260206004820152600e60248201526d0596f7520617265206e6f742053560941b60448201526064015b60405180910390fd5b60008281526020819052604090206002015460ff166101d55760405162461bcd60e51b815260206004820152601f60248201527f54686973206461746120686173206e6f74206265207265676973746572656400604482015260640161016b565b6000918252602082815260408084206001600160a01b0390931684529190529020805460ff19166001179055565b80516060906102199060029060208501906105f3565b506002805461022790610a44565b80601f016020809104026020016040519081016040528092919081815260200182805461025390610a44565b80156102a05780601f10610275576101008083540402835291602001916102a0565b820191906000526020600020905b81548152906001019060200180831161028357829003601f168201915b50505050509050919050565b6001546000906001600160a01b031633146102fa5760405162461bcd60e51b815260206004820152600e60248201526d0596f7520617265206e6f742053560941b604482015260640161016b565b60008360008151811061030f5761030f610a7f565b602090810291909101810151805190820120600081815280835260408082206001600160a01b0388168352909352919091205490915060ff166103aa5760405162461bcd60e51b815260206004820152602d60248201527f5468697320757365722068617320616c7265616479207265676973746572656460448201526c1030b9903a34329037bbb732b960991b606482015260840161016b565b60008181526020819052604090206002015460ff166104245760405162461bcd60e51b815260206004820152603060248201527f5468697320757365722068617320626520616c7265616479207265676973746560448201526f3932b21030b9903a34329037bbb732b960811b606482015260840161016b565b600081815260208181526040909120855161044792600190920191870190610677565b506000818152602081815260408083206001600160a01b03871684529091529020805460ff19166001179055905092915050565b600081815260208181526040808320600101805482518185028101850190935280835260609492939192909184015b828210156105565783829060005260206000200180546104c990610a44565b80601f01602080910402602001604051908101604052809291908181526020018280546104f590610a44565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b5050505050815260200190600101906104aa565b505050509050919050565b60606002805461057090610a44565b80601f016020809104026020016040519081016040528092919081815260200182805461059c90610a44565b80156105e95780601f106105be576101008083540402835291602001916105e9565b820191906000526020600020905b8154815290600101906020018083116105cc57829003601f168201915b5050505050905090565b8280546105ff90610a44565b90600052602060002090601f0160209004810192826106215760008555610667565b82601f1061063a57805160ff1916838001178555610667565b82800160010185558215610667579182015b8281111561066757825182559160200191906001019061064c565b506106739291506106d0565b5090565b8280548282559060005260206000209081019282156106c4579160200282015b828111156106c457825180516106b49184916020909101906105f3565b5091602001919060010190610697565b506106739291506106e5565b5b8082111561067357600081556001016106d1565b808211156106735760006106f98282610702565b506001016106e5565b50805461070e90610a44565b6000825580601f1061071e575050565b601f01602090049060005260206000209081019061073c91906106d0565b50565b80356001600160a01b038116811461075657600080fd5b919050565b6000806040838503121561076e57600080fd5b8235915061077e6020840161073f565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156107c6576107c6610787565b604052919050565b600067ffffffffffffffff8311156107e8576107e8610787565b6107fb601f8401601f191660200161079d565b905082815283838301111561080f57600080fd5b828260208301376000602084830101529392505050565b60006020828403121561083857600080fd5b813567ffffffffffffffff81111561084f57600080fd5b8201601f8101841361086057600080fd5b61086f848235602084016107ce565b949350505050565b6000815180845260005b8181101561089d57602081850181015186830182015201610881565b818111156108af576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006108d76020830184610877565b9392505050565b60008060408084860312156108f257600080fd5b833567ffffffffffffffff8082111561090a57600080fd5b818601915086601f83011261091e57600080fd5b813560208282111561093257610932610787565b8160051b61094182820161079d565b928352848101820192828101908b85111561095b57600080fd5b83870192505b848310156109ab578235868111156109795760008081fd5b8701603f81018d1361098b5760008081fd5b61099b8d868301358b84016107ce565b8352509183019190830190610961565b98506109bb91505088820161073f565b955050505050509250929050565b6000602082840312156109db57600080fd5b5035919050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015610a3757603f19888603018452610a25858351610877565b94509285019290850190600101610a09565b5092979650505050505050565b600181811c90821680610a5857607f821691505b60208210811415610a7957634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fdfea26469706673582212200e4a653b858ef0c1cc3239a8c969147375ec7bcc5ea349415973ee2db0f143b364736f6c63430008090033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _spa common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _spa)
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
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// FIT is a free data retrieval call binding the contract method 0xf1573dfb.
//
// Solidity: function FIT(uint256 ) view returns(bool isRegistered)
func (_Contracts *ContractsCaller) FIT(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "FIT", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FIT is a free data retrieval call binding the contract method 0xf1573dfb.
//
// Solidity: function FIT(uint256 ) view returns(bool isRegistered)
func (_Contracts *ContractsSession) FIT(arg0 *big.Int) (bool, error) {
	return _Contracts.Contract.FIT(&_Contracts.CallOpts, arg0)
}

// FIT is a free data retrieval call binding the contract method 0xf1573dfb.
//
// Solidity: function FIT(uint256 ) view returns(bool isRegistered)
func (_Contracts *ContractsCallerSession) FIT(arg0 *big.Int) (bool, error) {
	return _Contracts.Contract.FIT(&_Contracts.CallOpts, arg0)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(string)
func (_Contracts *ContractsCaller) Get(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(string)
func (_Contracts *ContractsSession) Get() (string, error) {
	return _Contracts.Contract.Get(&_Contracts.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(string)
func (_Contracts *ContractsCallerSession) Get() (string, error) {
	return _Contracts.Contract.Get(&_Contracts.CallOpts)
}

// GetHashedFile is a free data retrieval call binding the contract method 0x5e2b582a.
//
// Solidity: function getHashedFile(uint256 _fitId) view returns(bytes[])
func (_Contracts *ContractsCaller) GetHashedFile(opts *bind.CallOpts, _fitId *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getHashedFile", _fitId)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetHashedFile is a free data retrieval call binding the contract method 0x5e2b582a.
//
// Solidity: function getHashedFile(uint256 _fitId) view returns(bytes[])
func (_Contracts *ContractsSession) GetHashedFile(_fitId *big.Int) ([][]byte, error) {
	return _Contracts.Contract.GetHashedFile(&_Contracts.CallOpts, _fitId)
}

// GetHashedFile is a free data retrieval call binding the contract method 0x5e2b582a.
//
// Solidity: function getHashedFile(uint256 _fitId) view returns(bytes[])
func (_Contracts *ContractsCallerSession) GetHashedFile(_fitId *big.Int) ([][]byte, error) {
	return _Contracts.Contract.GetHashedFile(&_Contracts.CallOpts, _fitId)
}

// RegisterDedUpData is a paid mutator transaction binding the contract method 0x1040d067.
//
// Solidity: function registerDedUpData(uint256 _fitId, address _userAddress) returns()
func (_Contracts *ContractsTransactor) RegisterDedUpData(opts *bind.TransactOpts, _fitId *big.Int, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerDedUpData", _fitId, _userAddress)
}

// RegisterDedUpData is a paid mutator transaction binding the contract method 0x1040d067.
//
// Solidity: function registerDedUpData(uint256 _fitId, address _userAddress) returns()
func (_Contracts *ContractsSession) RegisterDedUpData(_fitId *big.Int, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterDedUpData(&_Contracts.TransactOpts, _fitId, _userAddress)
}

// RegisterDedUpData is a paid mutator transaction binding the contract method 0x1040d067.
//
// Solidity: function registerDedUpData(uint256 _fitId, address _userAddress) returns()
func (_Contracts *ContractsTransactorSession) RegisterDedUpData(_fitId *big.Int, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterDedUpData(&_Contracts.TransactOpts, _fitId, _userAddress)
}

// RegisterOriginalData is a paid mutator transaction binding the contract method 0x5334ba1f.
//
// Solidity: function registerOriginalData(bytes[] _fileData, address _userAddress) returns(uint256)
func (_Contracts *ContractsTransactor) RegisterOriginalData(opts *bind.TransactOpts, _fileData [][]byte, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerOriginalData", _fileData, _userAddress)
}

// RegisterOriginalData is a paid mutator transaction binding the contract method 0x5334ba1f.
//
// Solidity: function registerOriginalData(bytes[] _fileData, address _userAddress) returns(uint256)
func (_Contracts *ContractsSession) RegisterOriginalData(_fileData [][]byte, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterOriginalData(&_Contracts.TransactOpts, _fileData, _userAddress)
}

// RegisterOriginalData is a paid mutator transaction binding the contract method 0x5334ba1f.
//
// Solidity: function registerOriginalData(bytes[] _fileData, address _userAddress) returns(uint256)
func (_Contracts *ContractsTransactorSession) RegisterOriginalData(_fileData [][]byte, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterOriginalData(&_Contracts.TransactOpts, _fileData, _userAddress)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string _w) returns(string)
func (_Contracts *ContractsTransactor) Set(opts *bind.TransactOpts, _w string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "set", _w)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string _w) returns(string)
func (_Contracts *ContractsSession) Set(_w string) (*types.Transaction, error) {
	return _Contracts.Contract.Set(&_Contracts.TransactOpts, _w)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string _w) returns(string)
func (_Contracts *ContractsTransactorSession) Set(_w string) (*types.Transaction, error) {
	return _Contracts.Contract.Set(&_Contracts.TransactOpts, _w)
}