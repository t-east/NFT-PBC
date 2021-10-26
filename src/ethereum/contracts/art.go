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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spa\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tpa\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_pairing\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_g\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_u\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ArtLogTable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetParam\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"pairing\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"u\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"g\",\"type\":\"bytes\"}],\"internalType\":\"structIndexTable.Para\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"GetPubKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Logs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"chal\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"k1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"k2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"myu\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gamma\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PubKeys\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pubkey\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pubkey\",\"type\":\"string\"}],\"name\":\"RegisterPubKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileData\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"deleteData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fitId\",\"type\":\"uint256\"}],\"name\":\"getHashedFile\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_logId\",\"type\":\"uint256[]\"}],\"name\":\"getLog\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"chal\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"k1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"k2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"myu\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gamma\",\"type\":\"string\"}],\"internalType\":\"structIndexTable.Log[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fitId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"registerDedUpData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_chal\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_k1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_k2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_myu\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_gamma\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_fitId\",\"type\":\"uint256\"}],\"name\":\"registerLog\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileData\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"registerOriginalData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_logId\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_fitId\",\"type\":\"uint256\"}],\"name\":\"verifyLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001ec738038062001ec7833981016040819052620000349162000258565b600180546001600160a01b038088166001600160a01b0319928316179092556002805492871692909116919091179055825162000079906005906020860190620000b1565b5081516200008f906007906020850190620000b1565b508051620000a5906006906020840190620000b1565b50505050505062000365565b828054620000bf9062000328565b90600052602060002090601f016020900481019282620000e357600085556200012e565b82601f10620000fe57805160ff19168380011785556200012e565b828001600101855582156200012e579182015b828111156200012e57825182559160200191906001019062000111565b506200013c92915062000140565b5090565b5b808211156200013c576000815560010162000141565b80516001600160a01b03811681146200016f57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60006001600160401b0380841115620001a757620001a762000174565b604051601f8501601f19908116603f01168101908282118183101715620001d257620001d262000174565b81604052809350858152868686011115620001ec57600080fd5b600092505b8583101562000211578285015160208483010152602083019250620001f1565b8583111562000224576000602087830101525b5050509392505050565b600082601f8301126200024057600080fd5b62000251838351602085016200018a565b9392505050565b600080600080600060a086880312156200027157600080fd5b6200027c8662000157565b94506200028c6020870162000157565b60408701519094506001600160401b0380821115620002aa57600080fd5b818801915088601f830112620002bf57600080fd5b620002d0898351602085016200018a565b94506060880151915080821115620002e757600080fd5b620002f589838a016200022e565b935060808801519150808211156200030c57600080fd5b506200031b888289016200022e565b9150509295509295909350565b600181811c908216806200033d57607f821691505b602082108114156200035f57634e487b7160e01b600052602260045260246000fd5b50919050565b611b5280620003756000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80635e2b582a1161008c5780638704a867116100665780638704a867146101f35780638a04d5e514610206578063a5fe13f31461021b578063bc5696351461022e57600080fd5b80635e2b582a146101a057806364515e32146101c05780637aafb9f2146101d357600080fd5b80631040d067146100d4578063144acd97146100e957806330bd951414610117578063438696d9146101295780635334ba1f1461014957806354df3ae01461016a575b600080fd5b6100e76100e236600461134f565b610241565b005b6100fc6100f736600461137b565b610320565b60405161010e969594939291906113ec565b60405180910390f35b6100e7610125366004611538565b5050565b61013c61013736600461157d565b610601565b60405161010e91906115ba565b61015c6101573660046116f6565b6109fd565b60405190815260200161010e565b61019061017836600461137b565b60006020819052908152604090206002015460ff1681565b604051901515815260200161010e565b6101b36101ae36600461137b565b610bce565b60405161010e91906117cf565b61015c6101ce366004611868565b610cb4565b6101e66101e1366004611954565b610e10565b60405161010e919061196f565b6100e76102013660046116f6565b610ebc565b61020e610f3d565b60405161010e9190611982565b6101e6610229366004611954565b611129565b6100e761023c3660046119e3565b6111c7565b6001546001600160a01b031633146102915760405162461bcd60e51b815260206004820152600e60248201526d0596f7520617265206e6f742053560941b60448201526064015b60405180910390fd5b60008281526020819052604090206002015460ff166102f25760405162461bcd60e51b815260206004820152601f60248201527f54686973206461746120686173206e6f742062652072656769737465726564006044820152606401610288565b6000918252602082815260408084206001600160a01b0390931684529190529020805460ff19166001179055565b6003602052600090815260409020805460018201805460ff909216929161034690611a18565b80601f016020809104026020016040519081016040528092919081815260200182805461037290611a18565b80156103bf5780601f10610394576101008083540402835291602001916103bf565b820191906000526020600020905b8154815290600101906020018083116103a257829003601f168201915b5050505050908060020180546103d490611a18565b80601f016020809104026020016040519081016040528092919081815260200182805461040090611a18565b801561044d5780601f106104225761010080835404028352916020019161044d565b820191906000526020600020905b81548152906001019060200180831161043057829003601f168201915b50505050509080600301805461046290611a18565b80601f016020809104026020016040519081016040528092919081815260200182805461048e90611a18565b80156104db5780601f106104b0576101008083540402835291602001916104db565b820191906000526020600020905b8154815290600101906020018083116104be57829003601f168201915b5050505050908060040180546104f090611a18565b80601f016020809104026020016040519081016040528092919081815260200182805461051c90611a18565b80156105695780601f1061053e57610100808354040283529160200191610569565b820191906000526020600020905b81548152906001019060200180831161054c57829003601f168201915b50505050509080600501805461057e90611a18565b80601f01602080910402602001604051908101604052809291908181526020018280546105aa90611a18565b80156105f75780601f106105cc576101008083540402835291602001916105f7565b820191906000526020600020905b8154815290600101906020018083116105da57829003601f168201915b5050505050905086565b60606000825167ffffffffffffffff81111561061f5761061f611462565b60405190808252806020026020018201604052801561068b57816020015b6106786040518060c0016040528060001515815260200160608152602001606081526020016060815260200160608152602001606081525090565b81526020019060019003908161063d5790505b50905060005b83518110156109f657600360008583815181106106b0576106b0611a53565b602002602001015181526020019081526020016000206040518060c00160405290816000820160009054906101000a900460ff161515151581526020016001820180546106fc90611a18565b80601f016020809104026020016040519081016040528092919081815260200182805461072890611a18565b80156107755780601f1061074a57610100808354040283529160200191610775565b820191906000526020600020905b81548152906001019060200180831161075857829003601f168201915b5050505050815260200160028201805461078e90611a18565b80601f01602080910402602001604051908101604052809291908181526020018280546107ba90611a18565b80156108075780601f106107dc57610100808354040283529160200191610807565b820191906000526020600020905b8154815290600101906020018083116107ea57829003601f168201915b5050505050815260200160038201805461082090611a18565b80601f016020809104026020016040519081016040528092919081815260200182805461084c90611a18565b80156108995780601f1061086e57610100808354040283529160200191610899565b820191906000526020600020905b81548152906001019060200180831161087c57829003601f168201915b505050505081526020016004820180546108b290611a18565b80601f01602080910402602001604051908101604052809291908181526020018280546108de90611a18565b801561092b5780601f106109005761010080835404028352916020019161092b565b820191906000526020600020905b81548152906001019060200180831161090e57829003601f168201915b5050505050815260200160058201805461094490611a18565b80601f016020809104026020016040519081016040528092919081815260200182805461097090611a18565b80156109bd5780601f10610992576101008083540402835291602001916109bd565b820191906000526020600020905b8154815290600101906020018083116109a057829003601f168201915b5050505050815250508282815181106109d8576109d8611a53565b602002602001018190525080806109ee90611a69565b915050610691565b5092915050565b6001546000906001600160a01b03163314610a4b5760405162461bcd60e51b815260206004820152600e60248201526d0596f7520617265206e6f742053560941b6044820152606401610288565b600083600081518110610a6057610a60611a53565b602090810291909101810151805190820120600081815280835260408082206001600160a01b0388168352909352919091205490915060ff1615610afc5760405162461bcd60e51b815260206004820152602d60248201527f5468697320757365722068617320616c7265616479207265676973746572656460448201526c1030b9903a34329037bbb732b960991b6064820152608401610288565b60008181526020819052604090206002015460ff1615610b775760405162461bcd60e51b815260206004820152603060248201527f5468697320757365722068617320626520616c7265616479207265676973746560448201526f3932b21030b9903a34329037bbb732b960811b6064820152608401610288565b6000818152602081815260409091208551610b9a926001909201918701906111e7565b506000818152602081815260408083206001600160a01b03871684529091529020805460ff19166001179055905092915050565b600081815260208181526040808320600101805482518185028101850190935280835260609492939192909184015b82821015610ca9578382906000526020600020018054610c1c90611a18565b80601f0160208091040260200160405190810160405280929190818152602001828054610c4890611a18565b8015610c955780601f10610c6a57610100808354040283529160200191610c95565b820191906000526020600020905b815481529060010190602001808311610c7857829003601f168201915b505050505081526020019060010190610bfd565b505050509050919050565b6002546000906001600160a01b03163314610d035760405162461bcd60e51b815260206004820152600f60248201526e596f7520617265206e6f742054504160881b6044820152606401610288565b600088888888888888604051602001610d229796959493929190611a92565b60408051808303601f19018152918152815160209283012060008181526003845291909120805460ff19168c15151781558a51919350610d6b92600190910191908b0190611244565b5060008181526003602090815260409091208851610d91926002909201918a0190611244565b5060008181526003602081815260409092208851610db793919092019190890190611244565b5060008181526003602090815260409091208651610ddd92600490920191880190611244565b5060008181526003602090815260409091208551610e0392600590920191870190611244565b5098975050505050505050565b6001600160a01b0381166000908152600460205260409020805460609190610e3790611a18565b80601f0160208091040260200160405190810160405280929190818152602001828054610e6390611a18565b8015610eb05780601f10610e8557610100808354040283529160200191610eb0565b820191906000526020600020905b815481529060010190602001808311610e9357829003601f168201915b50505050509050919050565b600082600081518110610ed157610ed1611a53565b602090810291909101810151805190820120600081815280835260408082206001600160a01b0387168352808552908220805460ff199081169091558383529184526002810180549092169091558551919350610f3792600190910191908601906111e7565b50505050565b610f6160405180606001604052806060815260200160608152602001606081525090565b6005604051806060016040529081600082018054610f7e90611a18565b80601f0160208091040260200160405190810160405280929190818152602001828054610faa90611a18565b8015610ff75780601f10610fcc57610100808354040283529160200191610ff7565b820191906000526020600020905b815481529060010190602001808311610fda57829003601f168201915b5050505050815260200160018201805461101090611a18565b80601f016020809104026020016040519081016040528092919081815260200182805461103c90611a18565b80156110895780601f1061105e57610100808354040283529160200191611089565b820191906000526020600020905b81548152906001019060200180831161106c57829003601f168201915b505050505081526020016002820180546110a290611a18565b80601f01602080910402602001604051908101604052809291908181526020018280546110ce90611a18565b801561111b5780601f106110f05761010080835404028352916020019161111b565b820191906000526020600020905b8154815290600101906020018083116110fe57829003601f168201915b505050505081525050905090565b60046020526000908152604090208054819061114490611a18565b80601f016020809104026020016040519081016040528092919081815260200182805461117090611a18565b80156111bd5780601f10611192576101008083540402835291602001916111bd565b820191906000526020600020905b8154815290600101906020018083116111a057829003601f168201915b5050505050905081565b336000908152600460209081526040909120825161012592840190611244565b828054828255906000526020600020908101928215611234579160200282015b828111156112345782518051611224918491602090910190611244565b5091602001919060010190611207565b506112409291506112c4565b5090565b82805461125090611a18565b90600052602060002090601f01602090048101928261127257600085556112b8565b82601f1061128b57805160ff19168380011785556112b8565b828001600101855582156112b8579182015b828111156112b857825182559160200191906001019061129d565b506112409291506112e1565b808211156112405760006112d882826112f6565b506001016112c4565b5b8082111561124057600081556001016112e2565b50805461130290611a18565b6000825580601f10611312575050565b601f01602090049060005260206000209081019061133091906112e1565b50565b80356001600160a01b038116811461134a57600080fd5b919050565b6000806040838503121561136257600080fd5b8235915061137260208401611333565b90509250929050565b60006020828403121561138d57600080fd5b5035919050565b60005b838110156113af578181015183820152602001611397565b83811115610f375750506000910152565b600081518084526113d8816020860160208601611394565b601f01601f19169290920160200192915050565b861515815260c06020820152600061140760c08301886113c0565b828103604084015261141981886113c0565b9050828103606084015261142d81876113c0565b9050828103608084015261144181866113c0565b905082810360a084015261145581856113c0565b9998505050505050505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156114a1576114a1611462565b604052919050565b600067ffffffffffffffff8211156114c3576114c3611462565b5060051b60200190565b600082601f8301126114de57600080fd5b813560206114f36114ee836114a9565b611478565b82815260059290921b8401810191818101908684111561151257600080fd5b8286015b8481101561152d5780358352918301918301611516565b509695505050505050565b6000806040838503121561154b57600080fd5b823567ffffffffffffffff81111561156257600080fd5b61156e858286016114cd565b95602094909401359450505050565b60006020828403121561158f57600080fd5b813567ffffffffffffffff8111156115a657600080fd5b6115b2848285016114cd565b949350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561169057603f19898403018552815160c081511515855288820151818a870152611611828701826113c0565b915050878201518582038987015261162982826113c0565b9150506060808301518683038288015261164383826113c0565b925050506080808301518683038288015261165e83826113c0565b9250505060a0808301519250858203818701525061167c81836113c0565b9689019694505050908601906001016115e1565b509098975050505050505050565b600067ffffffffffffffff8311156116b8576116b8611462565b6116cb601f8401601f1916602001611478565b90508281528383830111156116df57600080fd5b828260208301376000602084830101529392505050565b600080604080848603121561170a57600080fd5b833567ffffffffffffffff8082111561172257600080fd5b818601915086601f83011261173657600080fd5b813560206117466114ee836114a9565b82815260059290921b8401810191818101908a84111561176557600080fd5b8286015b848110156117b1578035868111156117815760008081fd5b8701603f81018d136117935760008081fd5b6117a38d868301358b840161169e565b845250918301918301611769565b5097506117c19050888201611333565b955050505050509250929050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561182457603f198886030184526118128583516113c0565b945092850192908501906001016117f6565b5092979650505050505050565b8035801515811461134a57600080fd5b600082601f83011261185257600080fd5b6118618383356020850161169e565b9392505050565b600080600080600080600060e0888a03121561188357600080fd5b61188c88611831565b9650602088013567ffffffffffffffff808211156118a957600080fd5b6118b58b838c01611841565b975060408a01359150808211156118cb57600080fd5b6118d78b838c01611841565b965060608a01359150808211156118ed57600080fd5b6118f98b838c01611841565b955060808a013591508082111561190f57600080fd5b61191b8b838c01611841565b945060a08a013591508082111561193157600080fd5b5061193e8a828b01611841565b92505060c0880135905092959891949750929550565b60006020828403121561196657600080fd5b61186182611333565b60208152600061186160208301846113c0565b60208152600082516060602084015261199e60808401826113c0565b90506020840151601f19808584030160408601526119bc83836113c0565b92506040860151915080858403016060860152506119da82826113c0565b95945050505050565b6000602082840312156119f557600080fd5b813567ffffffffffffffff811115611a0c57600080fd5b6115b284828501611841565b600181811c90821680611a2c57607f821691505b60208210811415611a4d57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fd5b6000600019821415611a8b57634e487b7160e01b600052601160045260246000fd5b5060010190565b87151560f81b8152600060018851611ab08183860160208d01611394565b885190840190611ac68184840160208d01611394565b8851910190611adb8184840160208c01611394565b8751910190611af08184840160208b01611394565b8651910190611b058184840160208a01611394565b01908101939093525050602101969550505050505056fea264697066735822122074edeeedaf423f1858a6dfa783992dc10f8dc6fbcf01d6b89e4a93d6ada26e4464736f6c63430008090033",
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

// ArtLogTable is a free data retrieval call binding the contract method 0x54df3ae0.
//
// Solidity: function ArtLogTable(uint256 ) view returns(bool isRegistered)
func (_Contracts *ContractsCaller) ArtLogTable(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ArtLogTable", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ArtLogTable is a free data retrieval call binding the contract method 0x54df3ae0.
//
// Solidity: function ArtLogTable(uint256 ) view returns(bool isRegistered)
func (_Contracts *ContractsSession) ArtLogTable(arg0 *big.Int) (bool, error) {
	return _Contracts.Contract.ArtLogTable(&_Contracts.CallOpts, arg0)
}

// ArtLogTable is a free data retrieval call binding the contract method 0x54df3ae0.
//
// Solidity: function ArtLogTable(uint256 ) view returns(bool isRegistered)
func (_Contracts *ContractsCallerSession) ArtLogTable(arg0 *big.Int) (bool, error) {
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

// RegisterLog is a paid mutator transaction binding the contract method 0x64515e32.
//
// Solidity: function registerLog(bool _result, string _chal, string _k1, string _k2, string _myu, string _gamma, uint256 _fitId) returns(uint256)
func (_Contracts *ContractsTransactor) RegisterLog(opts *bind.TransactOpts, _result bool, _chal string, _k1 string, _k2 string, _myu string, _gamma string, _fitId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerLog", _result, _chal, _k1, _k2, _myu, _gamma, _fitId)
}

// RegisterLog is a paid mutator transaction binding the contract method 0x64515e32.
//
// Solidity: function registerLog(bool _result, string _chal, string _k1, string _k2, string _myu, string _gamma, uint256 _fitId) returns(uint256)
func (_Contracts *ContractsSession) RegisterLog(_result bool, _chal string, _k1 string, _k2 string, _myu string, _gamma string, _fitId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterLog(&_Contracts.TransactOpts, _result, _chal, _k1, _k2, _myu, _gamma, _fitId)
}

// RegisterLog is a paid mutator transaction binding the contract method 0x64515e32.
//
// Solidity: function registerLog(bool _result, string _chal, string _k1, string _k2, string _myu, string _gamma, uint256 _fitId) returns(uint256)
func (_Contracts *ContractsTransactorSession) RegisterLog(_result bool, _chal string, _k1 string, _k2 string, _myu string, _gamma string, _fitId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterLog(&_Contracts.TransactOpts, _result, _chal, _k1, _k2, _myu, _gamma, _fitId)
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

// VerifyLog is a paid mutator transaction binding the contract method 0x30bd9514.
//
// Solidity: function verifyLog(uint256[] _logId, uint256 _fitId) returns()
func (_Contracts *ContractsTransactor) VerifyLog(opts *bind.TransactOpts, _logId []*big.Int, _fitId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "verifyLog", _logId, _fitId)
}

// VerifyLog is a paid mutator transaction binding the contract method 0x30bd9514.
//
// Solidity: function verifyLog(uint256[] _logId, uint256 _fitId) returns()
func (_Contracts *ContractsSession) VerifyLog(_logId []*big.Int, _fitId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.VerifyLog(&_Contracts.TransactOpts, _logId, _fitId)
}

// VerifyLog is a paid mutator transaction binding the contract method 0x30bd9514.
//
// Solidity: function verifyLog(uint256[] _logId, uint256 _fitId) returns()
func (_Contracts *ContractsTransactorSession) VerifyLog(_logId []*big.Int, _fitId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.VerifyLog(&_Contracts.TransactOpts, _logId, _fitId)
}
