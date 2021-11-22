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
	Chal   uint32
	K1     []byte
	K2     []byte
	Myu    []byte
	Gamma  []byte
	ArtId  []byte
}

// IndexTableLogTable is an auto generated low-level Go binding around an user-defined struct.
type IndexTableLogTable struct {
	Log   IndexTableLog
	LogId []byte
}

// IndexTablePara is an auto generated low-level Go binding around an user-defined struct.
type IndexTablePara struct {
	Pairing string
	U       []byte
	G       []byte
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spa\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tpa\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_pairing\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_g\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_u\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"ArtLogTable\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"nftId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetArtIds\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_artId\",\"type\":\"bytes\"}],\"name\":\"GetHashedFile\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetLog\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"chal\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"k1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"k2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"myu\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"gamma\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"artId\",\"type\":\"bytes\"}],\"internalType\":\"structIndexTable.Log\",\"name\":\"Log\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"LogId\",\"type\":\"bytes\"}],\"internalType\":\"structIndexTable.LogTable[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetParam\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"pairing\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"u\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"g\",\"type\":\"bytes\"}],\"internalType\":\"structIndexTable.Para\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_artId\",\"type\":\"bytes\"}],\"name\":\"GetPubKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"Logs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"chal\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"k1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"k2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"myu\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"gamma\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"artId\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PubKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"RegisterPubKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_logId\",\"type\":\"bytes\"}],\"name\":\"SetAuditResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_chal\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_k1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_k2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_logId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_artId\",\"type\":\"bytes\"}],\"name\":\"SetLogId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"_artId\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"registerArt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_myu\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_gamma\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_artId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_logId\",\"type\":\"bytes\"}],\"name\":\"registerAuditProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_logId\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_artId\",\"type\":\"uint256\"}],\"name\":\"verifyLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200230e3803806200230e833981016040819052620000349162000258565b600180546001600160a01b038088166001600160a01b0319928316179092556002805492871692909116919091179055825162000079906006906020860190620000b1565b5081516200008f906008906020850190620000b1565b508051620000a5906007906020840190620000b1565b50505050505062000365565b828054620000bf9062000328565b90600052602060002090601f016020900481019282620000e357600085556200012e565b82601f10620000fe57805160ff19168380011785556200012e565b828001600101855582156200012e579182015b828111156200012e57825182559160200191906001019062000111565b506200013c92915062000140565b5090565b5b808211156200013c576000815560010162000141565b80516001600160a01b03811681146200016f57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60006001600160401b0380841115620001a757620001a762000174565b604051601f8501601f19908116603f01168101908282118183101715620001d257620001d262000174565b81604052809350858152868686011115620001ec57600080fd5b600092505b8583101562000211578285015160208483010152602083019250620001f1565b8583111562000224576000602087830101525b5050509392505050565b600082601f8301126200024057600080fd5b62000251838351602085016200018a565b9392505050565b600080600080600060a086880312156200027157600080fd5b6200027c8662000157565b94506200028c6020870162000157565b60408701519094506001600160401b0380821115620002aa57600080fd5b818801915088601f830112620002bf57600080fd5b620002d0898351602085016200018a565b94506060880151915080821115620002e757600080fd5b620002f589838a016200022e565b935060808801519150808211156200030c57600080fd5b506200031b888289016200022e565b9150509295509295909350565b600181811c908216806200033d57607f821691505b602082108114156200035f57634e487b7160e01b600052602260045260246000fd5b50919050565b611f9980620003756000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638a04d5e51161008c578063c14e4df111610066578063c14e4df114610253578063e870b3ae14610273578063efe7011f1461027b578063f32ab34c1461028e57600080fd5b80638a04d5e514610218578063a5fe13f31461022d578063b4e69d531461024057600080fd5b806330bd9514116100c857806330bd951414610140578063820ebe7014610152578063849202b114610165578063879826a5146101f257600080fd5b80630a18c4cc146100ef5780631b192fb1146101045780632521dfe114610117575b600080fd5b6101026100fd36600461176c565b6102a3565b005b6101026101123660046117a9565b6102c3565b61012a61012536600461176c565b610446565b60405161013791906118ca565b60405180910390f35b61010261014e366004611908565b5050565b6101026101603660046119c0565b610515565b6101bb61017336600461176c565b8051602081830181018051600082529282019190930120915280546002909101546001600160a01b0380831692600160a01b9081900463ffffffff1692918216910460ff1684565b604080516001600160a01b03958616815263ffffffff94909416602085015291909316908201529015156060820152608001610137565b61020561020036600461176c565b6106f9565b6040516101379796959493929190611aa4565b6102206109f2565b6040516101379190611b27565b61012a61023b366004611b88565b610bde565b61010261024e366004611ba3565b610c7c565b61026661026136600461176c565b610cfd565b6040516101379190611bf8565b610266610df7565b610102610289366004611c5a565b610ed0565b610296610fc5565b6040516101379190611d07565b336000908152600460209081526040909120825161014e92840190611509565b6002546001600160a01b031633146103145760405162461bcd60e51b815260206004820152600f60248201526e596f7520617265206e6f742054504160881b60448201526064015b60405180910390fd5b846003836040516103259190611e37565b908152602001604051809103902060000160016101000a81548163ffffffff021916908363ffffffff160217905550836003836040516103659190611e37565b90815260200160405180910390206001019080519060200190610389929190611509565b508260038360405161039b9190611e37565b908152602001604051809103902060020190805190602001906103bf929190611509565b50806003836040516103d19190611e37565b908152602001604051809103902060050190805190602001906103f5929190611509565b506000816040516104069190611e37565b908152604051602091819003820190206003018054600181018255600091825290829020845161043e93919092019190850190611509565b505050505050565b6060600080836040516104599190611e37565b9081526040805160209281900383019020546001600160a01b031660008181526004909352912080549192509061048f90611e53565b80601f01602080910402602001604051908101604052809291908181526020018280546104bb90611e53565b80156105085780601f106104dd57610100808354040283529160200191610508565b820191906000526020600020905b8154815290600101906020018083116104eb57829003601f168201915b5050505050915050919050565b6001546001600160a01b031633146105605760405162461bcd60e51b815260206004820152600e60248201526d0596f7520617265206e6f742053560941b604482015260640161030b565b6000826040516105709190611e37565b9081526040519081900360200190206002015460ff600160a01b90910416156105f45760405162461bcd60e51b815260206004820152603060248201527f5468697320757365722068617320626520616c7265616479207265676973746560448201526f3932b21030b9903a34329037bbb732b960811b606482015260840161030b565b336000836040516106059190611e37565b908152602001604051809103902060020160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508260008360405161064b9190611e37565b9081526020016040518091039020600101908051906020019061066f92919061158d565b50806000836040516106819190611e37565b9081526040516020918190038201902080546001600160a01b0319166001600160a01b0393909316929092179091556005805460018101825560009190915283516106f3927f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db090920191850190611509565b50505050565b80516020818301810180516003825292820191909301209152805460018201805460ff83169361010090930463ffffffff1692919061073790611e53565b80601f016020809104026020016040519081016040528092919081815260200182805461076390611e53565b80156107b05780601f10610785576101008083540402835291602001916107b0565b820191906000526020600020905b81548152906001019060200180831161079357829003601f168201915b5050505050908060020180546107c590611e53565b80601f01602080910402602001604051908101604052809291908181526020018280546107f190611e53565b801561083e5780601f106108135761010080835404028352916020019161083e565b820191906000526020600020905b81548152906001019060200180831161082157829003601f168201915b50505050509080600301805461085390611e53565b80601f016020809104026020016040519081016040528092919081815260200182805461087f90611e53565b80156108cc5780601f106108a1576101008083540402835291602001916108cc565b820191906000526020600020905b8154815290600101906020018083116108af57829003601f168201915b5050505050908060040180546108e190611e53565b80601f016020809104026020016040519081016040528092919081815260200182805461090d90611e53565b801561095a5780601f1061092f5761010080835404028352916020019161095a565b820191906000526020600020905b81548152906001019060200180831161093d57829003601f168201915b50505050509080600501805461096f90611e53565b80601f016020809104026020016040519081016040528092919081815260200182805461099b90611e53565b80156109e85780601f106109bd576101008083540402835291602001916109e8565b820191906000526020600020905b8154815290600101906020018083116109cb57829003601f168201915b5050505050905087565b610a1660405180606001604052806060815260200160608152602001606081525090565b6006604051806060016040529081600082018054610a3390611e53565b80601f0160208091040260200160405190810160405280929190818152602001828054610a5f90611e53565b8015610aac5780601f10610a8157610100808354040283529160200191610aac565b820191906000526020600020905b815481529060010190602001808311610a8f57829003601f168201915b50505050508152602001600182018054610ac590611e53565b80601f0160208091040260200160405190810160405280929190818152602001828054610af190611e53565b8015610b3e5780601f10610b1357610100808354040283529160200191610b3e565b820191906000526020600020905b815481529060010190602001808311610b2157829003601f168201915b50505050508152602001600282018054610b5790611e53565b80601f0160208091040260200160405190810160405280929190818152602001828054610b8390611e53565b8015610bd05780601f10610ba557610100808354040283529160200191610bd0565b820191906000526020600020905b815481529060010190602001808311610bb357829003601f168201915b505050505081525050905090565b600460205260009081526040902080548190610bf990611e53565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2590611e53565b8015610c725780601f10610c4757610100808354040283529160200191610c72565b820191906000526020600020905b815481529060010190602001808311610c5557829003601f168201915b5050505050905081565b6002546001600160a01b03163314610cc85760405162461bcd60e51b815260206004820152600f60248201526e596f7520617265206e6f742054504160881b604482015260640161030b565b81600382604051610cd99190611e37565b908152604051908190036020019020805491151560ff199092169190911790555050565b6060600082604051610d0f9190611e37565b9081526020016040518091039020600101805480602002602001604051908101604052809291908181526020016000905b82821015610dec578382906000526020600020018054610d5f90611e53565b80601f0160208091040260200160405190810160405280929190818152602001828054610d8b90611e53565b8015610dd85780601f10610dad57610100808354040283529160200191610dd8565b820191906000526020600020905b815481529060010190602001808311610dbb57829003601f168201915b505050505081526020019060010190610d40565b505050509050919050565b60606005805480602002602001604051908101604052809291908181526020016000905b82821015610ec7578382906000526020600020018054610e3a90611e53565b80601f0160208091040260200160405190810160405280929190818152602001828054610e6690611e53565b8015610eb35780601f10610e8857610100808354040283529160200191610eb3565b820191906000526020600020905b815481529060010190602001808311610e9657829003601f168201915b505050505081526020019060010190610e1b565b50505050905090565b336001600160a01b0316600083604051610eea9190611e37565b908152604051908190036020019020600201546001600160a01b031614610f535760405162461bcd60e51b815260206004820152601960248201527f596f752063616e6e6f742072656769737465722050726f6f6600000000000000604482015260640161030b565b83600382604051610f649190611e37565b90815260200160405180910390206003019080519060200190610f88929190611509565b5082600382604051610f9a9190611e37565b90815260200160405180910390206004019080519060200190610fbe929190611509565b5050505050565b60055460609060009067ffffffffffffffff811115610fe657610fe66116b5565b60405190808252806020026020018201604052801561101f57816020015b61100c6115e6565b8152602001906001900390816110045790505b50905060005b6005548110156115035760005b60006005838154811061104757611047611e88565b9060005260206000200160405161105e9190611e9e565b908152604051908190036020019020600301548110156114f05760606110826115e6565b60006005858154811061109757611097611e88565b906000526020600020016040516110ae9190611e9e565b908152602001604051809103902060030183815481106110d0576110d0611e88565b9060005260206000200180546110e590611e53565b80601f016020809104026020016040519081016040528092919081815260200182805461111190611e53565b801561115e5780601f106111335761010080835404028352916020019161115e565b820191906000526020600020905b81548152906001019060200180831161114157829003601f168201915b505050505091506003826040516111759190611e37565b9081526040519081900360200190205460ff166114db57602081018290526040516003906111a4908490611e37565b908152604080516020928190038301812060e082018352805460ff811615158352610100900463ffffffff16938201939093526001830180549193928401916111ec90611e53565b80601f016020809104026020016040519081016040528092919081815260200182805461121890611e53565b80156112655780601f1061123a57610100808354040283529160200191611265565b820191906000526020600020905b81548152906001019060200180831161124857829003601f168201915b5050505050815260200160028201805461127e90611e53565b80601f01602080910402602001604051908101604052809291908181526020018280546112aa90611e53565b80156112f75780601f106112cc576101008083540402835291602001916112f7565b820191906000526020600020905b8154815290600101906020018083116112da57829003601f168201915b5050505050815260200160038201805461131090611e53565b80601f016020809104026020016040519081016040528092919081815260200182805461133c90611e53565b80156113895780601f1061135e57610100808354040283529160200191611389565b820191906000526020600020905b81548152906001019060200180831161136c57829003601f168201915b505050505081526020016004820180546113a290611e53565b80601f01602080910402602001604051908101604052809291908181526020018280546113ce90611e53565b801561141b5780601f106113f05761010080835404028352916020019161141b565b820191906000526020600020905b8154815290600101906020018083116113fe57829003601f168201915b5050505050815260200160058201805461143490611e53565b80601f016020809104026020016040519081016040528092919081815260200182805461146090611e53565b80156114ad5780601f10611482576101008083540402835291602001916114ad565b820191906000526020600020905b81548152906001019060200180831161149057829003601f168201915b505050919092525050508152845181908690869081106114cf576114cf611e88565b60200260200101819052505b505080806114e890611f3a565b915050611032565b50806114fb81611f3a565b915050611025565b50919050565b82805461151590611e53565b90600052602060002090601f016020900481019282611537576000855561157d565b82601f1061155057805160ff191683800117855561157d565b8280016001018555821561157d579182015b8281111561157d578251825591602001919060010190611562565b50611589929150611646565b5090565b8280548282559060005260206000209081019282156115da579160200282015b828111156115da57825180516115ca918491602090910190611509565b50916020019190600101906115ad565b5061158992915061165b565b60405180604001604052806116396040518060e00160405280600015158152602001600063ffffffff16815260200160608152602001606081526020016060815260200160608152602001606081525090565b8152602001606081525090565b5b808211156115895760008155600101611647565b8082111561158957600061166f8282611678565b5060010161165b565b50805461168490611e53565b6000825580601f10611694575050565b601f0160209004906000526020600020908101906116b29190611646565b50565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156116f4576116f46116b5565b604052919050565b600082601f83011261170d57600080fd5b813567ffffffffffffffff811115611727576117276116b5565b61173a601f8201601f19166020016116cb565b81815284602083860101111561174f57600080fd5b816020850160208301376000918101602001919091529392505050565b60006020828403121561177e57600080fd5b813567ffffffffffffffff81111561179557600080fd5b6117a1848285016116fc565b949350505050565b600080600080600060a086880312156117c157600080fd5b853563ffffffff811681146117d557600080fd5b9450602086013567ffffffffffffffff808211156117f257600080fd5b6117fe89838a016116fc565b9550604088013591508082111561181457600080fd5b61182089838a016116fc565b9450606088013591508082111561183657600080fd5b61184289838a016116fc565b9350608088013591508082111561185857600080fd5b50611865888289016116fc565b9150509295509295909350565b60005b8381101561188d578181015183820152602001611875565b838111156106f35750506000910152565b600081518084526118b6816020860160208601611872565b601f01601f19169290920160200192915050565b6020815260006118dd602083018461189e565b9392505050565b600067ffffffffffffffff8211156118fe576118fe6116b5565b5060051b60200190565b6000806040838503121561191b57600080fd5b823567ffffffffffffffff81111561193257600080fd5b8301601f8101851361194357600080fd5b80356020611958611953836118e4565b6116cb565b82815260059290921b8301810191818101908884111561197757600080fd5b938201935b838510156119955784358252938201939082019061197c565b98969091013596505050505050565b80356001600160a01b03811681146119bb57600080fd5b919050565b6000806000606084860312156119d557600080fd5b833567ffffffffffffffff808211156119ed57600080fd5b818601915086601f830112611a0157600080fd5b81356020611a11611953836118e4565b82815260059290921b8401810191818101908a841115611a3057600080fd5b8286015b84811015611a6857803586811115611a4c5760008081fd5b611a5a8d86838b01016116fc565b845250918301918301611a34565b5097505087013592505080821115611a7f57600080fd5b50611a8c868287016116fc565b925050611a9b604085016119a4565b90509250925092565b871515815263ffffffff8716602082015260e060408201526000611acb60e083018861189e565b8281036060840152611add818861189e565b90508281036080840152611af1818761189e565b905082810360a0840152611b05818661189e565b905082810360c0840152611b19818561189e565b9a9950505050505050505050565b602081526000825160606020840152611b43608084018261189e565b90506020840151601f1980858403016040860152611b61838361189e565b9250604086015191508085840301606086015250611b7f828261189e565b95945050505050565b600060208284031215611b9a57600080fd5b6118dd826119a4565b60008060408385031215611bb657600080fd5b82358015158114611bc657600080fd5b9150602083013567ffffffffffffffff811115611be257600080fd5b611bee858286016116fc565b9150509250929050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015611c4d57603f19888603018452611c3b85835161189e565b94509285019290850190600101611c1f565b5092979650505050505050565b60008060008060808587031215611c7057600080fd5b843567ffffffffffffffff80821115611c8857600080fd5b611c94888389016116fc565b95506020870135915080821115611caa57600080fd5b611cb6888389016116fc565b94506040870135915080821115611ccc57600080fd5b611cd8888389016116fc565b93506060870135915080821115611cee57600080fd5b50611cfb878288016116fc565b91505092959194509250565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015611e2957603f19808a8503018652825180518886528051151589870152898101516060611d668189018363ffffffff169052565b8a830151915060e0608081818b0152611d836101208b018561189e565b935082850151925060a0878b860301818c0152611da0858561189e565b945081860151935060c09150878b860301828c0152611dbf858561189e565b945080860151935050868a850301828b0152611ddb848461189e565b94015189850387016101008b0152939250611dfa91508290508361189e565b9350505088810151905084820389860152611e15828261189e565b968901969450505090860190600101611d2e565b509098975050505050505050565b60008251611e49818460208701611872565b9190910192915050565b600181811c90821680611e6757607f821691505b6020821081141561150357634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600080835481600182811c915080831680611eba57607f831692505b6020808410821415611eda57634e487b7160e01b86526022600452602486fd5b818015611eee5760018114611eff57611f2c565b60ff19861689528489019650611f2c565b60008a81526020902060005b86811015611f245781548b820152908501908301611f0b565b505084890196505b509498975050505050505050565b6000600019821415611f5c57634e487b7160e01b600052601160045260246000fd5b506001019056fea264697066735822122009a61d1acac2af09f608cdf2f6da3622f3757d12296d7d39561e44b9c84230a164736f6c63430008090033",
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

// ArtLogTable is a free data retrieval call binding the contract method 0x849202b1.
//
// Solidity: function ArtLogTable(bytes ) view returns(address owner, uint32 nftId, address provider, bool isRegistered)
func (_Contracts *ContractsCaller) ArtLogTable(opts *bind.CallOpts, arg0 []byte) (struct {
	Owner        common.Address
	NftId        uint32
	Provider     common.Address
	IsRegistered bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ArtLogTable", arg0)

	outstruct := new(struct {
		Owner        common.Address
		NftId        uint32
		Provider     common.Address
		IsRegistered bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Provider = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.IsRegistered = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// ArtLogTable is a free data retrieval call binding the contract method 0x849202b1.
//
// Solidity: function ArtLogTable(bytes ) view returns(address owner, uint32 nftId, address provider, bool isRegistered)
func (_Contracts *ContractsSession) ArtLogTable(arg0 []byte) (struct {
	Owner        common.Address
	NftId        uint32
	Provider     common.Address
	IsRegistered bool
}, error) {
	return _Contracts.Contract.ArtLogTable(&_Contracts.CallOpts, arg0)
}

// ArtLogTable is a free data retrieval call binding the contract method 0x849202b1.
//
// Solidity: function ArtLogTable(bytes ) view returns(address owner, uint32 nftId, address provider, bool isRegistered)
func (_Contracts *ContractsCallerSession) ArtLogTable(arg0 []byte) (struct {
	Owner        common.Address
	NftId        uint32
	Provider     common.Address
	IsRegistered bool
}, error) {
	return _Contracts.Contract.ArtLogTable(&_Contracts.CallOpts, arg0)
}

// GetArtIds is a free data retrieval call binding the contract method 0xe870b3ae.
//
// Solidity: function GetArtIds() view returns(bytes[])
func (_Contracts *ContractsCaller) GetArtIds(opts *bind.CallOpts) ([][]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "GetArtIds")

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetArtIds is a free data retrieval call binding the contract method 0xe870b3ae.
//
// Solidity: function GetArtIds() view returns(bytes[])
func (_Contracts *ContractsSession) GetArtIds() ([][]byte, error) {
	return _Contracts.Contract.GetArtIds(&_Contracts.CallOpts)
}

// GetArtIds is a free data retrieval call binding the contract method 0xe870b3ae.
//
// Solidity: function GetArtIds() view returns(bytes[])
func (_Contracts *ContractsCallerSession) GetArtIds() ([][]byte, error) {
	return _Contracts.Contract.GetArtIds(&_Contracts.CallOpts)
}

// GetHashedFile is a free data retrieval call binding the contract method 0xc14e4df1.
//
// Solidity: function GetHashedFile(bytes _artId) view returns(bytes[])
func (_Contracts *ContractsCaller) GetHashedFile(opts *bind.CallOpts, _artId []byte) ([][]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "GetHashedFile", _artId)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetHashedFile is a free data retrieval call binding the contract method 0xc14e4df1.
//
// Solidity: function GetHashedFile(bytes _artId) view returns(bytes[])
func (_Contracts *ContractsSession) GetHashedFile(_artId []byte) ([][]byte, error) {
	return _Contracts.Contract.GetHashedFile(&_Contracts.CallOpts, _artId)
}

// GetHashedFile is a free data retrieval call binding the contract method 0xc14e4df1.
//
// Solidity: function GetHashedFile(bytes _artId) view returns(bytes[])
func (_Contracts *ContractsCallerSession) GetHashedFile(_artId []byte) ([][]byte, error) {
	return _Contracts.Contract.GetHashedFile(&_Contracts.CallOpts, _artId)
}

// GetLog is a free data retrieval call binding the contract method 0xf32ab34c.
//
// Solidity: function GetLog() view returns(((bool,uint32,bytes,bytes,bytes,bytes,bytes),bytes)[])
func (_Contracts *ContractsCaller) GetLog(opts *bind.CallOpts) ([]IndexTableLogTable, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "GetLog")

	if err != nil {
		return *new([]IndexTableLogTable), err
	}

	out0 := *abi.ConvertType(out[0], new([]IndexTableLogTable)).(*[]IndexTableLogTable)

	return out0, err

}

// GetLog is a free data retrieval call binding the contract method 0xf32ab34c.
//
// Solidity: function GetLog() view returns(((bool,uint32,bytes,bytes,bytes,bytes,bytes),bytes)[])
func (_Contracts *ContractsSession) GetLog() ([]IndexTableLogTable, error) {
	return _Contracts.Contract.GetLog(&_Contracts.CallOpts)
}

// GetLog is a free data retrieval call binding the contract method 0xf32ab34c.
//
// Solidity: function GetLog() view returns(((bool,uint32,bytes,bytes,bytes,bytes,bytes),bytes)[])
func (_Contracts *ContractsCallerSession) GetLog() ([]IndexTableLogTable, error) {
	return _Contracts.Contract.GetLog(&_Contracts.CallOpts)
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

// GetPubKey is a free data retrieval call binding the contract method 0x2521dfe1.
//
// Solidity: function GetPubKey(bytes _artId) view returns(bytes)
func (_Contracts *ContractsCaller) GetPubKey(opts *bind.CallOpts, _artId []byte) ([]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "GetPubKey", _artId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPubKey is a free data retrieval call binding the contract method 0x2521dfe1.
//
// Solidity: function GetPubKey(bytes _artId) view returns(bytes)
func (_Contracts *ContractsSession) GetPubKey(_artId []byte) ([]byte, error) {
	return _Contracts.Contract.GetPubKey(&_Contracts.CallOpts, _artId)
}

// GetPubKey is a free data retrieval call binding the contract method 0x2521dfe1.
//
// Solidity: function GetPubKey(bytes _artId) view returns(bytes)
func (_Contracts *ContractsCallerSession) GetPubKey(_artId []byte) ([]byte, error) {
	return _Contracts.Contract.GetPubKey(&_Contracts.CallOpts, _artId)
}

// Logs is a free data retrieval call binding the contract method 0x879826a5.
//
// Solidity: function Logs(bytes ) view returns(bool result, uint32 chal, bytes k1, bytes k2, bytes myu, bytes gamma, bytes artId)
func (_Contracts *ContractsCaller) Logs(opts *bind.CallOpts, arg0 []byte) (struct {
	Result bool
	Chal   uint32
	K1     []byte
	K2     []byte
	Myu    []byte
	Gamma  []byte
	ArtId  []byte
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "Logs", arg0)

	outstruct := new(struct {
		Result bool
		Chal   uint32
		K1     []byte
		K2     []byte
		Myu    []byte
		Gamma  []byte
		ArtId  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Result = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Chal = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.K1 = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.K2 = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.Myu = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.Gamma = *abi.ConvertType(out[5], new([]byte)).(*[]byte)
	outstruct.ArtId = *abi.ConvertType(out[6], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Logs is a free data retrieval call binding the contract method 0x879826a5.
//
// Solidity: function Logs(bytes ) view returns(bool result, uint32 chal, bytes k1, bytes k2, bytes myu, bytes gamma, bytes artId)
func (_Contracts *ContractsSession) Logs(arg0 []byte) (struct {
	Result bool
	Chal   uint32
	K1     []byte
	K2     []byte
	Myu    []byte
	Gamma  []byte
	ArtId  []byte
}, error) {
	return _Contracts.Contract.Logs(&_Contracts.CallOpts, arg0)
}

// Logs is a free data retrieval call binding the contract method 0x879826a5.
//
// Solidity: function Logs(bytes ) view returns(bool result, uint32 chal, bytes k1, bytes k2, bytes myu, bytes gamma, bytes artId)
func (_Contracts *ContractsCallerSession) Logs(arg0 []byte) (struct {
	Result bool
	Chal   uint32
	K1     []byte
	K2     []byte
	Myu    []byte
	Gamma  []byte
	ArtId  []byte
}, error) {
	return _Contracts.Contract.Logs(&_Contracts.CallOpts, arg0)
}

// PubKeys is a free data retrieval call binding the contract method 0xa5fe13f3.
//
// Solidity: function PubKeys(address ) view returns(bytes pubkey)
func (_Contracts *ContractsCaller) PubKeys(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "PubKeys", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PubKeys is a free data retrieval call binding the contract method 0xa5fe13f3.
//
// Solidity: function PubKeys(address ) view returns(bytes pubkey)
func (_Contracts *ContractsSession) PubKeys(arg0 common.Address) ([]byte, error) {
	return _Contracts.Contract.PubKeys(&_Contracts.CallOpts, arg0)
}

// PubKeys is a free data retrieval call binding the contract method 0xa5fe13f3.
//
// Solidity: function PubKeys(address ) view returns(bytes pubkey)
func (_Contracts *ContractsCallerSession) PubKeys(arg0 common.Address) ([]byte, error) {
	return _Contracts.Contract.PubKeys(&_Contracts.CallOpts, arg0)
}

// RegisterPubKey is a paid mutator transaction binding the contract method 0x0a18c4cc.
//
// Solidity: function RegisterPubKey(bytes _pubkey) returns()
func (_Contracts *ContractsTransactor) RegisterPubKey(opts *bind.TransactOpts, _pubkey []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "RegisterPubKey", _pubkey)
}

// RegisterPubKey is a paid mutator transaction binding the contract method 0x0a18c4cc.
//
// Solidity: function RegisterPubKey(bytes _pubkey) returns()
func (_Contracts *ContractsSession) RegisterPubKey(_pubkey []byte) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterPubKey(&_Contracts.TransactOpts, _pubkey)
}

// RegisterPubKey is a paid mutator transaction binding the contract method 0x0a18c4cc.
//
// Solidity: function RegisterPubKey(bytes _pubkey) returns()
func (_Contracts *ContractsTransactorSession) RegisterPubKey(_pubkey []byte) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterPubKey(&_Contracts.TransactOpts, _pubkey)
}

// SetAuditResult is a paid mutator transaction binding the contract method 0xb4e69d53.
//
// Solidity: function SetAuditResult(bool _result, bytes _logId) returns()
func (_Contracts *ContractsTransactor) SetAuditResult(opts *bind.TransactOpts, _result bool, _logId []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "SetAuditResult", _result, _logId)
}

// SetAuditResult is a paid mutator transaction binding the contract method 0xb4e69d53.
//
// Solidity: function SetAuditResult(bool _result, bytes _logId) returns()
func (_Contracts *ContractsSession) SetAuditResult(_result bool, _logId []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetAuditResult(&_Contracts.TransactOpts, _result, _logId)
}

// SetAuditResult is a paid mutator transaction binding the contract method 0xb4e69d53.
//
// Solidity: function SetAuditResult(bool _result, bytes _logId) returns()
func (_Contracts *ContractsTransactorSession) SetAuditResult(_result bool, _logId []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetAuditResult(&_Contracts.TransactOpts, _result, _logId)
}

// SetLogId is a paid mutator transaction binding the contract method 0x1b192fb1.
//
// Solidity: function SetLogId(uint32 _chal, bytes _k1, bytes _k2, bytes _logId, bytes _artId) returns()
func (_Contracts *ContractsTransactor) SetLogId(opts *bind.TransactOpts, _chal uint32, _k1 []byte, _k2 []byte, _logId []byte, _artId []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "SetLogId", _chal, _k1, _k2, _logId, _artId)
}

// SetLogId is a paid mutator transaction binding the contract method 0x1b192fb1.
//
// Solidity: function SetLogId(uint32 _chal, bytes _k1, bytes _k2, bytes _logId, bytes _artId) returns()
func (_Contracts *ContractsSession) SetLogId(_chal uint32, _k1 []byte, _k2 []byte, _logId []byte, _artId []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetLogId(&_Contracts.TransactOpts, _chal, _k1, _k2, _logId, _artId)
}

// SetLogId is a paid mutator transaction binding the contract method 0x1b192fb1.
//
// Solidity: function SetLogId(uint32 _chal, bytes _k1, bytes _k2, bytes _logId, bytes _artId) returns()
func (_Contracts *ContractsTransactorSession) SetLogId(_chal uint32, _k1 []byte, _k2 []byte, _logId []byte, _artId []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetLogId(&_Contracts.TransactOpts, _chal, _k1, _k2, _logId, _artId)
}

// RegisterArt is a paid mutator transaction binding the contract method 0x820ebe70.
//
// Solidity: function registerArt(bytes[] _fileData, bytes _artId, address _userAddress) returns()
func (_Contracts *ContractsTransactor) RegisterArt(opts *bind.TransactOpts, _fileData [][]byte, _artId []byte, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerArt", _fileData, _artId, _userAddress)
}

// RegisterArt is a paid mutator transaction binding the contract method 0x820ebe70.
//
// Solidity: function registerArt(bytes[] _fileData, bytes _artId, address _userAddress) returns()
func (_Contracts *ContractsSession) RegisterArt(_fileData [][]byte, _artId []byte, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterArt(&_Contracts.TransactOpts, _fileData, _artId, _userAddress)
}

// RegisterArt is a paid mutator transaction binding the contract method 0x820ebe70.
//
// Solidity: function registerArt(bytes[] _fileData, bytes _artId, address _userAddress) returns()
func (_Contracts *ContractsTransactorSession) RegisterArt(_fileData [][]byte, _artId []byte, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterArt(&_Contracts.TransactOpts, _fileData, _artId, _userAddress)
}

// RegisterAuditProof is a paid mutator transaction binding the contract method 0xefe7011f.
//
// Solidity: function registerAuditProof(bytes _myu, bytes _gamma, bytes _artId, bytes _logId) returns()
func (_Contracts *ContractsTransactor) RegisterAuditProof(opts *bind.TransactOpts, _myu []byte, _gamma []byte, _artId []byte, _logId []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerAuditProof", _myu, _gamma, _artId, _logId)
}

// RegisterAuditProof is a paid mutator transaction binding the contract method 0xefe7011f.
//
// Solidity: function registerAuditProof(bytes _myu, bytes _gamma, bytes _artId, bytes _logId) returns()
func (_Contracts *ContractsSession) RegisterAuditProof(_myu []byte, _gamma []byte, _artId []byte, _logId []byte) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterAuditProof(&_Contracts.TransactOpts, _myu, _gamma, _artId, _logId)
}

// RegisterAuditProof is a paid mutator transaction binding the contract method 0xefe7011f.
//
// Solidity: function registerAuditProof(bytes _myu, bytes _gamma, bytes _artId, bytes _logId) returns()
func (_Contracts *ContractsTransactorSession) RegisterAuditProof(_myu []byte, _gamma []byte, _artId []byte, _logId []byte) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterAuditProof(&_Contracts.TransactOpts, _myu, _gamma, _artId, _logId)
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
