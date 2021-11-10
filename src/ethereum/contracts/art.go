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

// IndexTableLog is an auto generated low-level Go binding around an user-defined struct.
type IndexTableLog struct {
	Result bool
	Chal   string
	K1     string
	K2     string
	Myu    string
	Gamma  string
}

// IndexTablePara is an auto generated low-level Go binding around an user-defined struct.
type IndexTablePara struct {
	Pairing string
	U       []byte
	G       []byte
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spa\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tpa\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_pairing\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_g\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_u\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ArtLogTable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetParam\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"pairing\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"u\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"g\",\"type\":\"bytes\"}],\"internalType\":\"structIndexTable.Para\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"GetPubKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Logs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"chal\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"k1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"k2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"myu\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gamma\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PubKeys\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pubkey\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pubkey\",\"type\":\"string\"}],\"name\":\"RegisterPubKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_artId\",\"type\":\"string\"}],\"name\":\"getHashedFile\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_logId\",\"type\":\"uint256[]\"}],\"name\":\"getLog\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"chal\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"k1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"k2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"myu\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gamma\",\"type\":\"string\"}],\"internalType\":\"structIndexTable.Log[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileData\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"_artId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"registerArt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_chal\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_k1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_k2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_myu\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_gamma\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_artId\",\"type\":\"uint256\"}],\"name\":\"registerLog\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_logId\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_artId\",\"type\":\"uint256\"}],\"name\":\"verifyLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001cdd38038062001cdd833981016040819052620000349162000258565b600180546001600160a01b038088166001600160a01b0319928316179092556002805492871692909116919091179055825162000079906005906020860190620000b1565b5081516200008f906007906020850190620000b1565b508051620000a5906006906020840190620000b1565b50505050505062000365565b828054620000bf9062000328565b90600052602060002090601f016020900481019282620000e357600085556200012e565b82601f10620000fe57805160ff19168380011785556200012e565b828001600101855582156200012e579182015b828111156200012e57825182559160200191906001019062000111565b506200013c92915062000140565b5090565b5b808211156200013c576000815560010162000141565b80516001600160a01b03811681146200016f57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60006001600160401b0380841115620001a757620001a762000174565b604051601f8501601f19908116603f01168101908282118183101715620001d257620001d262000174565b81604052809350858152868686011115620001ec57600080fd5b600092505b8583101562000211578285015160208483010152602083019250620001f1565b8583111562000224576000602087830101525b5050509392505050565b600082601f8301126200024057600080fd5b62000251838351602085016200018a565b9392505050565b600080600080600060a086880312156200027157600080fd5b6200027c8662000157565b94506200028c6020870162000157565b60408701519094506001600160401b0380821115620002aa57600080fd5b818801915088601f830112620002bf57600080fd5b620002d0898351602085016200018a565b94506060880151915080821115620002e757600080fd5b620002f589838a016200022e565b935060808801519150808211156200030c57600080fd5b506200031b888289016200022e565b9150509295509295909350565b600181811c908216806200033d57607f821691505b602082108114156200035f57634e487b7160e01b600052602260045260246000fd5b50919050565b61196880620003756000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638a04d5e5116100715780638a04d5e514610151578063a3339a3d14610166578063a5fe13f314610186578063bc56963514610199578063ed461d3b146101ac578063fdcf2a46146101bf57600080fd5b8063144acd97146100ae57806330bd9514146100dc578063438696d9146100f057806364515e32146101105780637aafb9f214610131575b600080fd5b6100c16100bc366004611135565b610200565b6040516100d3969594939291906111aa565b60405180910390f35b6100ee6100ea3660046112f6565b5050565b005b6101036100fe36600461133b565b6104e1565b6040516100d39190611378565b61012361011e3660046114f0565b6108dd565b6040519081526020016100d3565b61014461013f3660046115f3565b610a3e565b6040516100d3919061160e565b610159610aea565b6040516100d39190611621565b610179610174366004611682565b610cd6565b6040516100d391906116b7565b6101446101943660046115f3565b610dd0565b6100ee6101a7366004611682565b610e6e565b6100ee6101ba366004611719565b610e8e565b6101f06101cd366004611682565b805160208183018101805160008252928201919093012091526002015460ff1681565b60405190151581526020016100d3565b6003602052600090815260409020805460018201805460ff909216929161022690611812565b80601f016020809104026020016040519081016040528092919081815260200182805461025290611812565b801561029f5780601f106102745761010080835404028352916020019161029f565b820191906000526020600020905b81548152906001019060200180831161028257829003601f168201915b5050505050908060020180546102b490611812565b80601f01602080910402602001604051908101604052809291908181526020018280546102e090611812565b801561032d5780601f106103025761010080835404028352916020019161032d565b820191906000526020600020905b81548152906001019060200180831161031057829003601f168201915b50505050509080600301805461034290611812565b80601f016020809104026020016040519081016040528092919081815260200182805461036e90611812565b80156103bb5780601f10610390576101008083540402835291602001916103bb565b820191906000526020600020905b81548152906001019060200180831161039e57829003601f168201915b5050505050908060040180546103d090611812565b80601f01602080910402602001604051908101604052809291908181526020018280546103fc90611812565b80156104495780601f1061041e57610100808354040283529160200191610449565b820191906000526020600020905b81548152906001019060200180831161042c57829003601f168201915b50505050509080600501805461045e90611812565b80601f016020809104026020016040519081016040528092919081815260200182805461048a90611812565b80156104d75780601f106104ac576101008083540402835291602001916104d7565b820191906000526020600020905b8154815290600101906020018083116104ba57829003601f168201915b5050505050905086565b60606000825167ffffffffffffffff8111156104ff576104ff611220565b60405190808252806020026020018201604052801561056b57816020015b6105586040518060c0016040528060001515815260200160608152602001606081526020016060815260200160608152602001606081525090565b81526020019060019003908161051d5790505b50905060005b83518110156108d657600360008583815181106105905761059061184d565b602002602001015181526020019081526020016000206040518060c00160405290816000820160009054906101000a900460ff161515151581526020016001820180546105dc90611812565b80601f016020809104026020016040519081016040528092919081815260200182805461060890611812565b80156106555780601f1061062a57610100808354040283529160200191610655565b820191906000526020600020905b81548152906001019060200180831161063857829003601f168201915b5050505050815260200160028201805461066e90611812565b80601f016020809104026020016040519081016040528092919081815260200182805461069a90611812565b80156106e75780601f106106bc576101008083540402835291602001916106e7565b820191906000526020600020905b8154815290600101906020018083116106ca57829003601f168201915b5050505050815260200160038201805461070090611812565b80601f016020809104026020016040519081016040528092919081815260200182805461072c90611812565b80156107795780601f1061074e57610100808354040283529160200191610779565b820191906000526020600020905b81548152906001019060200180831161075c57829003601f168201915b5050505050815260200160048201805461079290611812565b80601f01602080910402602001604051908101604052809291908181526020018280546107be90611812565b801561080b5780601f106107e05761010080835404028352916020019161080b565b820191906000526020600020905b8154815290600101906020018083116107ee57829003601f168201915b5050505050815260200160058201805461082490611812565b80601f016020809104026020016040519081016040528092919081815260200182805461085090611812565b801561089d5780601f106108725761010080835404028352916020019161089d565b820191906000526020600020905b81548152906001019060200180831161088057829003601f168201915b5050505050815250508282815181106108b8576108b861184d565b602002602001018190525080806108ce90611863565b915050610571565b5092915050565b6002546000906001600160a01b031633146109315760405162461bcd60e51b815260206004820152600f60248201526e596f7520617265206e6f742054504160881b60448201526064015b60405180910390fd5b600088888888888888604051602001610950979695949392919061188c565b60408051808303601f19018152918152815160209283012060008181526003845291909120805460ff19168c15151781558a5191935061099992600190910191908b0190610fe9565b50600081815260036020908152604090912088516109bf926002909201918a0190610fe9565b50600081815260036020818152604090922088516109e593919092019190890190610fe9565b5060008181526003602090815260409091208651610a0b92600490920191880190610fe9565b5060008181526003602090815260409091208551610a3192600590920191870190610fe9565b5098975050505050505050565b6001600160a01b0381166000908152600460205260409020805460609190610a6590611812565b80601f0160208091040260200160405190810160405280929190818152602001828054610a9190611812565b8015610ade5780601f10610ab357610100808354040283529160200191610ade565b820191906000526020600020905b815481529060010190602001808311610ac157829003601f168201915b50505050509050919050565b610b0e60405180606001604052806060815260200160608152602001606081525090565b6005604051806060016040529081600082018054610b2b90611812565b80601f0160208091040260200160405190810160405280929190818152602001828054610b5790611812565b8015610ba45780601f10610b7957610100808354040283529160200191610ba4565b820191906000526020600020905b815481529060010190602001808311610b8757829003601f168201915b50505050508152602001600182018054610bbd90611812565b80601f0160208091040260200160405190810160405280929190818152602001828054610be990611812565b8015610c365780601f10610c0b57610100808354040283529160200191610c36565b820191906000526020600020905b815481529060010190602001808311610c1957829003601f168201915b50505050508152602001600282018054610c4f90611812565b80601f0160208091040260200160405190810160405280929190818152602001828054610c7b90611812565b8015610cc85780601f10610c9d57610100808354040283529160200191610cc8565b820191906000526020600020905b815481529060010190602001808311610cab57829003601f168201915b505050505081525050905090565b6060600082604051610ce89190611916565b9081526020016040518091039020600101805480602002602001604051908101604052809291908181526020016000905b82821015610dc5578382906000526020600020018054610d3890611812565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6490611812565b8015610db15780601f10610d8657610100808354040283529160200191610db1565b820191906000526020600020905b815481529060010190602001808311610d9457829003601f168201915b505050505081526020019060010190610d19565b505050509050919050565b600460205260009081526040902080548190610deb90611812565b80601f0160208091040260200160405190810160405280929190818152602001828054610e1790611812565b8015610e645780601f10610e3957610100808354040283529160200191610e64565b820191906000526020600020905b815481529060010190602001808311610e4757829003601f168201915b5050505050905081565b33600090815260046020908152604090912082516100ea92840190610fe9565b6001546001600160a01b03163314610ed95760405162461bcd60e51b815260206004820152600e60248201526d0596f7520617265206e6f742053560941b6044820152606401610928565b600082604051610ee99190611916565b9081526040519081900360200190206002015460ff1615610f655760405162461bcd60e51b815260206004820152603060248201527f5468697320757365722068617320626520616c7265616479207265676973746560448201526f3932b21030b9903a34329037bbb732b960811b6064820152608401610928565b82600083604051610f769190611916565b90815260200160405180910390206001019080519060200190610f9a92919061106d565b506001600083604051610fad9190611916565b90815260408051602092819003830190206001600160a01b0394909416600090815293909152909120805460ff19169115159190911790555050565b828054610ff590611812565b90600052602060002090601f016020900481019282611017576000855561105d565b82601f1061103057805160ff191683800117855561105d565b8280016001018555821561105d579182015b8281111561105d578251825591602001919060010190611042565b506110699291506110c6565b5090565b8280548282559060005260206000209081019282156110ba579160200282015b828111156110ba57825180516110aa918491602090910190610fe9565b509160200191906001019061108d565b506110699291506110db565b5b8082111561106957600081556001016110c7565b808211156110695760006110ef82826110f8565b506001016110db565b50805461110490611812565b6000825580601f10611114575050565b601f01602090049060005260206000209081019061113291906110c6565b50565b60006020828403121561114757600080fd5b5035919050565b60005b83811015611169578181015183820152602001611151565b83811115611178576000848401525b50505050565b6000815180845261119681602086016020860161114e565b601f01601f19169290920160200192915050565b861515815260c0602082015260006111c560c083018861117e565b82810360408401526111d7818861117e565b905082810360608401526111eb818761117e565b905082810360808401526111ff818661117e565b905082810360a0840152611213818561117e565b9998505050505050505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561125f5761125f611220565b604052919050565b600067ffffffffffffffff82111561128157611281611220565b5060051b60200190565b600082601f83011261129c57600080fd5b813560206112b16112ac83611267565b611236565b82815260059290921b840181019181810190868411156112d057600080fd5b8286015b848110156112eb57803583529183019183016112d4565b509695505050505050565b6000806040838503121561130957600080fd5b823567ffffffffffffffff81111561132057600080fd5b61132c8582860161128b565b95602094909401359450505050565b60006020828403121561134d57600080fd5b813567ffffffffffffffff81111561136457600080fd5b6113708482850161128b565b949350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561144e57603f19898403018552815160c081511515855288820151818a8701526113cf8287018261117e565b91505087820151858203898701526113e7828261117e565b91505060608083015186830382880152611401838261117e565b925050506080808301518683038288015261141c838261117e565b9250505060a0808301519250858203818701525061143a818361117e565b96890196945050509086019060010161139f565b509098975050505050505050565b8035801515811461146c57600080fd5b919050565b600067ffffffffffffffff83111561148b5761148b611220565b61149e601f8401601f1916602001611236565b90508281528383830111156114b257600080fd5b828260208301376000602084830101529392505050565b600082601f8301126114da57600080fd5b6114e983833560208501611471565b9392505050565b600080600080600080600060e0888a03121561150b57600080fd5b6115148861145c565b9650602088013567ffffffffffffffff8082111561153157600080fd5b61153d8b838c016114c9565b975060408a013591508082111561155357600080fd5b61155f8b838c016114c9565b965060608a013591508082111561157557600080fd5b6115818b838c016114c9565b955060808a013591508082111561159757600080fd5b6115a38b838c016114c9565b945060a08a01359150808211156115b957600080fd5b506115c68a828b016114c9565b92505060c0880135905092959891949750929550565b80356001600160a01b038116811461146c57600080fd5b60006020828403121561160557600080fd5b6114e9826115dc565b6020815260006114e9602083018461117e565b60208152600082516060602084015261163d608084018261117e565b90506020840151601f198085840301604086015261165b838361117e565b9250604086015191508085840301606086015250611679828261117e565b95945050505050565b60006020828403121561169457600080fd5b813567ffffffffffffffff8111156116ab57600080fd5b611370848285016114c9565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561170c57603f198886030184526116fa85835161117e565b945092850192908501906001016116de565b5092979650505050505050565b60008060006060848603121561172e57600080fd5b833567ffffffffffffffff8082111561174657600080fd5b818601915086601f83011261175a57600080fd5b8135602061176a6112ac83611267565b82815260059290921b8401810191818101908a84111561178957600080fd5b8286015b848110156117d6578035868111156117a55760008081fd5b8701603f81018d136117b75760008081fd5b6117c88d8683013560408401611471565b84525091830191830161178d565b50975050870135925050808211156117ed57600080fd5b506117fa868287016114c9565b925050611809604085016115dc565b90509250925092565b600181811c9082168061182657607f821691505b6020821081141561184757634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fd5b600060001982141561188557634e487b7160e01b600052601160045260246000fd5b5060010190565b87151560f81b81526000600188516118aa8183860160208d0161114e565b8851908401906118c08184840160208d0161114e565b88519101906118d58184840160208c0161114e565b87519101906118ea8184840160208b0161114e565b86519101906118ff8184840160208a0161114e565b019081019390935250506021019695505050505050565b6000825161192881846020870161114e565b919091019291505056fea26469706673582212200d44292a34552a6b4ec748062f5c49690be99966b9062b0f96a61c667f12c71264736f6c63430008090033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _spa common.Address, _tpa common.Address, _pairing string, _g []byte, _u []byte) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _spa, _tpa, _pairing, _g, _u)
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

// ArtLogTable is a free data retrieval call binding the contract method 0xfdcf2a46.
//
// Solidity: function ArtLogTable(string ) view returns(bool isRegistered)
func (_Contracts *ContractsCaller) ArtLogTable(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ArtLogTable", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ArtLogTable is a free data retrieval call binding the contract method 0xfdcf2a46.
//
// Solidity: function ArtLogTable(string ) view returns(bool isRegistered)
func (_Contracts *ContractsSession) ArtLogTable(arg0 string) (bool, error) {
	return _Contracts.Contract.ArtLogTable(&_Contracts.CallOpts, arg0)
}

// ArtLogTable is a free data retrieval call binding the contract method 0xfdcf2a46.
//
// Solidity: function ArtLogTable(string ) view returns(bool isRegistered)
func (_Contracts *ContractsCallerSession) ArtLogTable(arg0 string) (bool, error) {
	return _Contracts.Contract.ArtLogTable(&_Contracts.CallOpts, arg0)
}

// GetParam is a free data retrieval call binding the contract method 0x8a04d5e5.
//
// Solidity: function GetParam() view returns((string,bytes,bytes))
func (_Contracts *ContractsCaller) GetParam(opts *bind.CallOpts) (IndexTablePara, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "GetParam")

	if err != nil {
		return *new(IndexTablePara), err
	}

	out0 := *abi.ConvertType(out[0], new(IndexTablePara)).(*IndexTablePara)

	return out0, err

}

// GetParam is a free data retrieval call binding the contract method 0x8a04d5e5.
//
// Solidity: function GetParam() view returns((string,bytes,bytes))
func (_Contracts *ContractsSession) GetParam() (IndexTablePara, error) {
	return _Contracts.Contract.GetParam(&_Contracts.CallOpts)
}

// GetParam is a free data retrieval call binding the contract method 0x8a04d5e5.
//
// Solidity: function GetParam() view returns((string,bytes,bytes))
func (_Contracts *ContractsCallerSession) GetParam() (IndexTablePara, error) {
	return _Contracts.Contract.GetParam(&_Contracts.CallOpts)
}

// GetPubKey is a free data retrieval call binding the contract method 0x7aafb9f2.
//
// Solidity: function GetPubKey(address _owner) view returns(string)
func (_Contracts *ContractsCaller) GetPubKey(opts *bind.CallOpts, _owner common.Address) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "GetPubKey", _owner)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPubKey is a free data retrieval call binding the contract method 0x7aafb9f2.
//
// Solidity: function GetPubKey(address _owner) view returns(string)
func (_Contracts *ContractsSession) GetPubKey(_owner common.Address) (string, error) {
	return _Contracts.Contract.GetPubKey(&_Contracts.CallOpts, _owner)
}

// GetPubKey is a free data retrieval call binding the contract method 0x7aafb9f2.
//
// Solidity: function GetPubKey(address _owner) view returns(string)
func (_Contracts *ContractsCallerSession) GetPubKey(_owner common.Address) (string, error) {
	return _Contracts.Contract.GetPubKey(&_Contracts.CallOpts, _owner)
}

// Logs is a free data retrieval call binding the contract method 0x144acd97.
//
// Solidity: function Logs(uint256 ) view returns(bool result, string chal, string k1, string k2, string myu, string gamma)
func (_Contracts *ContractsCaller) Logs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Result bool
	Chal   string
	K1     string
	K2     string
	Myu    string
	Gamma  string
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "Logs", arg0)

	outstruct := new(struct {
		Result bool
		Chal   string
		K1     string
		K2     string
		Myu    string
		Gamma  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Result = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Chal = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.K1 = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.K2 = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Myu = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Gamma = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// Logs is a free data retrieval call binding the contract method 0x144acd97.
//
// Solidity: function Logs(uint256 ) view returns(bool result, string chal, string k1, string k2, string myu, string gamma)
func (_Contracts *ContractsSession) Logs(arg0 *big.Int) (struct {
	Result bool
	Chal   string
	K1     string
	K2     string
	Myu    string
	Gamma  string
}, error) {
	return _Contracts.Contract.Logs(&_Contracts.CallOpts, arg0)
}

// Logs is a free data retrieval call binding the contract method 0x144acd97.
//
// Solidity: function Logs(uint256 ) view returns(bool result, string chal, string k1, string k2, string myu, string gamma)
func (_Contracts *ContractsCallerSession) Logs(arg0 *big.Int) (struct {
	Result bool
	Chal   string
	K1     string
	K2     string
	Myu    string
	Gamma  string
}, error) {
	return _Contracts.Contract.Logs(&_Contracts.CallOpts, arg0)
}

// PubKeys is a free data retrieval call binding the contract method 0xa5fe13f3.
//
// Solidity: function PubKeys(address ) view returns(string pubkey)
func (_Contracts *ContractsCaller) PubKeys(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "PubKeys", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PubKeys is a free data retrieval call binding the contract method 0xa5fe13f3.
//
// Solidity: function PubKeys(address ) view returns(string pubkey)
func (_Contracts *ContractsSession) PubKeys(arg0 common.Address) (string, error) {
	return _Contracts.Contract.PubKeys(&_Contracts.CallOpts, arg0)
}

// PubKeys is a free data retrieval call binding the contract method 0xa5fe13f3.
//
// Solidity: function PubKeys(address ) view returns(string pubkey)
func (_Contracts *ContractsCallerSession) PubKeys(arg0 common.Address) (string, error) {
	return _Contracts.Contract.PubKeys(&_Contracts.CallOpts, arg0)
}

// GetHashedFile is a free data retrieval call binding the contract method 0xa3339a3d.
//
// Solidity: function getHashedFile(string _artId) view returns(bytes[])
func (_Contracts *ContractsCaller) GetHashedFile(opts *bind.CallOpts, _artId string) ([][]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getHashedFile", _artId)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetHashedFile is a free data retrieval call binding the contract method 0xa3339a3d.
//
// Solidity: function getHashedFile(string _artId) view returns(bytes[])
func (_Contracts *ContractsSession) GetHashedFile(_artId string) ([][]byte, error) {
	return _Contracts.Contract.GetHashedFile(&_Contracts.CallOpts, _artId)
}

// GetHashedFile is a free data retrieval call binding the contract method 0xa3339a3d.
//
// Solidity: function getHashedFile(string _artId) view returns(bytes[])
func (_Contracts *ContractsCallerSession) GetHashedFile(_artId string) ([][]byte, error) {
	return _Contracts.Contract.GetHashedFile(&_Contracts.CallOpts, _artId)
}

// GetLog is a free data retrieval call binding the contract method 0x438696d9.
//
// Solidity: function getLog(uint256[] _logId) view returns((bool,string,string,string,string,string)[])
func (_Contracts *ContractsCaller) GetLog(opts *bind.CallOpts, _logId []*big.Int) ([]IndexTableLog, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLog", _logId)

	if err != nil {
		return *new([]IndexTableLog), err
	}

	out0 := *abi.ConvertType(out[0], new([]IndexTableLog)).(*[]IndexTableLog)

	return out0, err

}

// GetLog is a free data retrieval call binding the contract method 0x438696d9.
//
// Solidity: function getLog(uint256[] _logId) view returns((bool,string,string,string,string,string)[])
func (_Contracts *ContractsSession) GetLog(_logId []*big.Int) ([]IndexTableLog, error) {
	return _Contracts.Contract.GetLog(&_Contracts.CallOpts, _logId)
}

// GetLog is a free data retrieval call binding the contract method 0x438696d9.
//
// Solidity: function getLog(uint256[] _logId) view returns((bool,string,string,string,string,string)[])
func (_Contracts *ContractsCallerSession) GetLog(_logId []*big.Int) ([]IndexTableLog, error) {
	return _Contracts.Contract.GetLog(&_Contracts.CallOpts, _logId)
}

// RegisterPubKey is a paid mutator transaction binding the contract method 0xbc569635.
//
// Solidity: function RegisterPubKey(string _pubkey) returns()
func (_Contracts *ContractsTransactor) RegisterPubKey(opts *bind.TransactOpts, _pubkey string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "RegisterPubKey", _pubkey)
}

// RegisterPubKey is a paid mutator transaction binding the contract method 0xbc569635.
//
// Solidity: function RegisterPubKey(string _pubkey) returns()
func (_Contracts *ContractsSession) RegisterPubKey(_pubkey string) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterPubKey(&_Contracts.TransactOpts, _pubkey)
}

// RegisterPubKey is a paid mutator transaction binding the contract method 0xbc569635.
//
// Solidity: function RegisterPubKey(string _pubkey) returns()
func (_Contracts *ContractsTransactorSession) RegisterPubKey(_pubkey string) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterPubKey(&_Contracts.TransactOpts, _pubkey)
}

// RegisterArt is a paid mutator transaction binding the contract method 0xed461d3b.
//
// Solidity: function registerArt(bytes[] _fileData, string _artId, address _userAddress) returns()
func (_Contracts *ContractsTransactor) RegisterArt(opts *bind.TransactOpts, _fileData [][]byte, _artId string, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerArt", _fileData, _artId, _userAddress)
}

// RegisterArt is a paid mutator transaction binding the contract method 0xed461d3b.
//
// Solidity: function registerArt(bytes[] _fileData, string _artId, address _userAddress) returns()
func (_Contracts *ContractsSession) RegisterArt(_fileData [][]byte, _artId string, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterArt(&_Contracts.TransactOpts, _fileData, _artId, _userAddress)
}

// RegisterArt is a paid mutator transaction binding the contract method 0xed461d3b.
//
// Solidity: function registerArt(bytes[] _fileData, string _artId, address _userAddress) returns()
func (_Contracts *ContractsTransactorSession) RegisterArt(_fileData [][]byte, _artId string, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterArt(&_Contracts.TransactOpts, _fileData, _artId, _userAddress)
}

// RegisterLog is a paid mutator transaction binding the contract method 0x64515e32.
//
// Solidity: function registerLog(bool _result, string _chal, string _k1, string _k2, string _myu, string _gamma, uint256 _artId) returns(uint256)
func (_Contracts *ContractsTransactor) RegisterLog(opts *bind.TransactOpts, _result bool, _chal string, _k1 string, _k2 string, _myu string, _gamma string, _artId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerLog", _result, _chal, _k1, _k2, _myu, _gamma, _artId)
}

// RegisterLog is a paid mutator transaction binding the contract method 0x64515e32.
//
// Solidity: function registerLog(bool _result, string _chal, string _k1, string _k2, string _myu, string _gamma, uint256 _artId) returns(uint256)
func (_Contracts *ContractsSession) RegisterLog(_result bool, _chal string, _k1 string, _k2 string, _myu string, _gamma string, _artId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterLog(&_Contracts.TransactOpts, _result, _chal, _k1, _k2, _myu, _gamma, _artId)
}

// RegisterLog is a paid mutator transaction binding the contract method 0x64515e32.
//
// Solidity: function registerLog(bool _result, string _chal, string _k1, string _k2, string _myu, string _gamma, uint256 _artId) returns(uint256)
func (_Contracts *ContractsTransactorSession) RegisterLog(_result bool, _chal string, _k1 string, _k2 string, _myu string, _gamma string, _artId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterLog(&_Contracts.TransactOpts, _result, _chal, _k1, _k2, _myu, _gamma, _artId)
}

// VerifyLog is a paid mutator transaction binding the contract method 0x30bd9514.
//
// Solidity: function verifyLog(uint256[] _logId, uint256 _artId) returns()
func (_Contracts *ContractsTransactor) VerifyLog(opts *bind.TransactOpts, _logId []*big.Int, _artId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "verifyLog", _logId, _artId)
}

// VerifyLog is a paid mutator transaction binding the contract method 0x30bd9514.
//
// Solidity: function verifyLog(uint256[] _logId, uint256 _artId) returns()
func (_Contracts *ContractsSession) VerifyLog(_logId []*big.Int, _artId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.VerifyLog(&_Contracts.TransactOpts, _logId, _artId)
}

// VerifyLog is a paid mutator transaction binding the contract method 0x30bd9514.
//
// Solidity: function verifyLog(uint256[] _logId, uint256 _artId) returns()
func (_Contracts *ContractsTransactorSession) VerifyLog(_logId []*big.Int, _artId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.VerifyLog(&_Contracts.TransactOpts, _logId, _artId)
}
