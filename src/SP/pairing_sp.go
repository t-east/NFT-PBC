package SP

import (
	"pairing_test/src/tool"
	"reflect"
	"fmt"

	"github.com/Nik-U/pbc"
	"pairing_test/src/ethereum/ethhandler"
)

// ファイルをストレージに保存
func SaveOutsourceData(outsourceData tool.Storage, userId int) {
	var Users []int
	FileTable := tool.InputFIT()
	Users = append(Users, userId)
	hashedFile := tool.FileToMMData(outsourceData)
	validFile := tool.FileIndexTable{UserID: Users, HashedFile: hashedFile}
	conn, client := ethhandler.ConnectNetWork()
	auth := ethhandler.AuthUser(client)
	userAddress := ethhandler.StringToAddress("0x30b93c46e417bE95A18F5374EEAf5ca6AE457b01")
	// reply, err := conn.DeleteData(auth, validFile.HashedFile, userAddress)
	// fmt.Print("\ndelete\n")
	// if err != nil{
	// 	fmt.Print("\ndelete error\n")
	// 	fmt.Println(err)
	// }
	// fmt.Println(reply)
	
	reply, err := conn.RegisterOriginalData(auth, validFile.HashedFile, userAddress)
	if err != nil{
		fmt.Print("\nregister error\n")
		fmt.Println(err)
	}
	fmt.Println(reply)

	conn, client = ethhandler.ConnectNetWork()
	auth = ethhandler.AuthUser(client)
	userAddress = ethhandler.StringToAddress("0x30b93c46e417bE95A18F5374EEAf5ca6AE457b01")
	ethhandler.Set(auth,conn,"Hello")
	fmt.Print("\nset\n")

	ethhandler.Get(conn)
	fmt.Print("\nget\n")
	
	tool.NewFIT(validFile)
	FileTable = append(FileTable, validFile)
}

// 監査証明作成
func AuditProofGen(params tool.Params, challen []tool.Chal) []tool.ProofT {
	var myuT *pbc.Element
	var gammaT *pbc.Element
	var proofTs []tool.ProofT
	storages := tool.ReadStorage()
	fit := tool.InputFIT()
	pairing, _ := pbc.NewPairingFromString(params.Pairing)
	for i := 0; i < len(fit); i++ {
		for l:=0;l<len(storages);l++{
			hashedFile := tool.FileToMMData(storages[l])
			if reflect.DeepEqual(fit[i].HashedFile,hashedFile){
				var proofT tool.ProofT
				splitedFile, _ := tool.SplitSlice(storages[l].File, len(storages[l].MetaData))
				aTable, vTable := tool.HashChallen(len(storages[l].MetaData), challen[i], pairing)
		
				
				for j := 0; j < challen[i].C; j++ {
					meta := pairing.NewG1().SetBytes(storages[l].MetaData[aTable[j]])
					m := pairing.NewG1().SetFromHash(splitedFile[aTable[j]])
					if j == 0 {
						myuT = pairing.NewZr().MulBig(vTable[j], m.X())
						gammaT = pairing.NewG1().PowZn(meta, vTable[j])
					} else {
						myuT = pairing.NewZr().Add(myuT, pairing.NewZr().MulBig(vTable[j], m.X()))
						gammaT.Mul(gammaT, pairing.NewG1().PowZn(meta, vTable[j]))
					}
				}
				proofT.Myu = myuT.Bytes()
				proofT.Gamma = gammaT.Bytes()
				proofT.ID = challen[i].ID
				proofTs = append(proofTs, proofT)
			}
		}
	}
	return proofTs
}