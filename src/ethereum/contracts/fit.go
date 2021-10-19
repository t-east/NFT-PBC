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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spa\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"FIT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileData\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"deleteData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fitId\",\"type\":\"uint256\"}],\"name\":\"getHashedFile\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fitId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"registerDedUpData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileData\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"registerOriginalData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_w\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610c07380380610c0783398101604081905261002f91610115565b600180546001600160a01b0319166001600160a01b038316179055604080518082019091526004808252631d195cdd60e21b60209092019182526100759160029161007c565b5050610180565b82805461008890610145565b90600052602060002090601f0160209004810192826100aa57600085556100f0565b82601f106100c357805160ff19168380011785556100f0565b828001600101855582156100f0579182015b828111156100f05782518255916020019190600101906100d5565b506100fc929150610100565b5090565b5b808211156100fc5760008155600101610101565b60006020828403121561012757600080fd5b81516001600160a01b038116811461013e57600080fd5b9392505050565b600181811c9082168061015957607f821691505b6020821081141561017a57634e487b7160e01b600052602260045260246000fd5b50919050565b610a788061018f6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80635e2b582a1161005b5780635e2b582a146100e15780636d4ce63c146101015780638704a86714610109578063f1573dfb1461011c57600080fd5b80631040d067146100825780634ed3885e146100975780635334ba1f146100c0575b600080fd5b610095610090366004610708565b610152565b005b6100aa6100a53660046107d3565b610231565b6040516100b79190610871565b60405180910390f35b6100d36100ce36600461088b565b6102da565b6040519081526020016100b7565b6100f46100ef366004610976565b6103a7565b6040516100b7919061098f565b6100aa61048d565b61009561011736600461088b565b61051f565b61014261012a366004610976565b60006020819052908152604090206002015460ff1681565b60405190151581526020016100b7565b6001546001600160a01b031633146101a25760405162461bcd60e51b815260206004820152600e60248201526d0596f7520617265206e6f742053560941b60448201526064015b60405180910390fd5b60008281526020819052604090206002015460ff166102035760405162461bcd60e51b815260206004820152601f60248201527f54686973206461746120686173206e6f742062652072656769737465726564006044820152606401610199565b6000918252602082815260408084206001600160a01b0390931684529190529020805460ff19166001179055565b80516060906102479060029060208501906105a0565b5060028054610255906109f1565b80601f0160208091040260200160405190810160405280929190818152602001828054610281906109f1565b80156102ce5780601f106102a3576101008083540402835291602001916102ce565b820191906000526020600020905b8154815290600101906020018083116102b157829003601f168201915b50505050509050919050565b6001546000906001600160a01b031633146103285760405162461bcd60e51b815260206004820152600e60248201526d0596f7520617265206e6f742053560941b6044820152606401610199565b60008360008151811061033d5761033d610a2c565b60209081029190910181015180519082012060008181528083526040902086519193506103739260019091019190870190610624565b506000818152602081815260408083206001600160a01b03871684529091529020805460ff19166001179055905092915050565b600081815260208181526040808320600101805482518185028101850190935280835260609492939192909184015b828210156104825783829060005260206000200180546103f5906109f1565b80601f0160208091040260200160405190810160405280929190818152602001828054610421906109f1565b801561046e5780601f106104435761010080835404028352916020019161046e565b820191906000526020600020905b81548152906001019060200180831161045157829003601f168201915b5050505050815260200190600101906103d6565b505050509050919050565b60606002805461049c906109f1565b80601f01602080910402602001604051908101604052809291908181526020018280546104c8906109f1565b80156105155780601f106104ea57610100808354040283529160200191610515565b820191906000526020600020905b8154815290600101906020018083116104f857829003601f168201915b5050505050905090565b60008260008151811061053457610534610a2c565b602090810291909101810151805190820120600081815280835260408082206001600160a01b0387168352808552908220805460ff19908116909155838352918452600281018054909216909155855191935061059a9260019091019190860190610624565b50505050565b8280546105ac906109f1565b90600052602060002090601f0160209004810192826105ce5760008555610614565b82601f106105e757805160ff1916838001178555610614565b82800160010185558215610614579182015b828111156106145782518255916020019190600101906105f9565b5061062092915061067d565b5090565b828054828255906000526020600020908101928215610671579160200282015b8281111561067157825180516106619184916020909101906105a0565b5091602001919060010190610644565b50610620929150610692565b5b80821115610620576000815560010161067e565b808211156106205760006106a682826106af565b50600101610692565b5080546106bb906109f1565b6000825580601f106106cb575050565b601f0160209004906000526020600020908101906106e9919061067d565b50565b80356001600160a01b038116811461070357600080fd5b919050565b6000806040838503121561071b57600080fd5b8235915061072b602084016106ec565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561077357610773610734565b604052919050565b600067ffffffffffffffff83111561079557610795610734565b6107a8601f8401601f191660200161074a565b90508281528383830111156107bc57600080fd5b828260208301376000602084830101529392505050565b6000602082840312156107e557600080fd5b813567ffffffffffffffff8111156107fc57600080fd5b8201601f8101841361080d57600080fd5b61081c8482356020840161077b565b949350505050565b6000815180845260005b8181101561084a5760208185018101518683018201520161082e565b8181111561085c576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006108846020830184610824565b9392505050565b600080604080848603121561089f57600080fd5b833567ffffffffffffffff808211156108b757600080fd5b818601915086601f8301126108cb57600080fd5b81356020828211156108df576108df610734565b8160051b6108ee82820161074a565b928352848101820192828101908b85111561090857600080fd5b83870192505b84831015610958578235868111156109265760008081fd5b8701603f81018d136109385760008081fd5b6109488d868301358b840161077b565b835250918301919083019061090e565b98506109689150508882016106ec565b955050505050509250929050565b60006020828403121561098857600080fd5b5035919050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156109e457603f198886030184526109d2858351610824565b945092850192908501906001016109b6565b5092979650505050505050565b600181811c90821680610a0557607f821691505b60208210811415610a2657634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fdfea2646970667358221220fc6c10d16d940abf9428157bf68a571c0e31273adb71b6a81eb405c789f6f13464736f6c63430008090033",
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

// DeleteData is a paid mutator transaction binding the contract method 0x8704a867.
//
// Solidity: function deleteData(bytes[] _fileData, address _userAddress) returns()
func (_Contracts *ContractsTransactor) DeleteData(opts *bind.TransactOpts, _fileData [][]byte, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deleteData", _fileData, _userAddress)
}

// DeleteData is a paid mutator transaction binding the contract method 0x8704a867.
//
// Solidity: function deleteData(bytes[] _fileData, address _userAddress) returns()
func (_Contracts *ContractsSession) DeleteData(_fileData [][]byte, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.DeleteData(&_Contracts.TransactOpts, _fileData, _userAddress)
}

// DeleteData is a paid mutator transaction binding the contract method 0x8704a867.
//
// Solidity: function deleteData(bytes[] _fileData, address _userAddress) returns()
func (_Contracts *ContractsTransactorSession) DeleteData(_fileData [][]byte, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.DeleteData(&_Contracts.TransactOpts, _fileData, _userAddress)
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
