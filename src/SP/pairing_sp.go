package SP

import (
	"pairing_test/src/tool"
	"reflect"
	"fmt"

	"github.com/Nik-U/pbc"
	"pairing_test/src/ethereum/ethhandler"
)

func SaveOutsourceData(outsourceData tool.Storage, userId int) {
	var Users []int
	FileTable := tool.InputFIT()
	Users = append(Users, userId)
	hashedFile := tool.FileToMMData(outsourceData)
	validFile := tool.FileIndexTable{UserID: Users, HashedFile: hashedFile}
	conn, client := ethhandler.ConnectNetWork()
	auth := ethhandler.AuthUser(client)
	userAddress := ethhandler.StringToAddress("0xD7bdEe86c43402d55F28397603fB38D399D50314")
	reply, err := conn.RegisterOriginalData(auth, validFile.HashedFile, userAddress)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(reply)
	}
	tool.NewFIT(validFile)
	FileTable = append(FileTable, validFile)
}

func DedupChallen(params tool.Params, n int) tool.Chal {
	pairing, _ := pbc.NewPairingFromString(params.Pairing)
	ck, ks1, ks2 := tool.ChallenGen(n, pairing)
	challen := tool.Chal{C: ck, K1: ks1.Bytes(), K2: ks2.Bytes()}
	return challen
}

// Alice generates a keypair and signs a message
func DedupVerify(fileData tool.Storage, params tool.Params, proofs []byte, challen tool.Chal, fitNum int) int {
	var storage tool.Storage
	pairing, _ := pbc.NewPairingFromString(params.Pairing)
	dataBlockNum := len(fileData.MetaData)
	storages := tool.ReadStorage()
	fit := tool.InputFIT()
	for s := 0; s < len(storages); s++ {
		hashedFile := tool.FileToMMData(storages[s])
		if reflect.DeepEqual(hashedFile, fit[fitNum].HashedFile) {
			storage = storages[s]
		}
	}
	aTable, vTable := tool.HashChallen(dataBlockNum, challen, pairing)

	g := pairing.NewG1().SetBytes(params.G)
	pubKey := pairing.NewG1().SetBytes(params.PubKeys[fit[fitNum].UserID[0]].PubKey)
	u := pairing.NewG1().SetBytes(params.U)
	proof := pairing.NewZr().SetBytes(proofs)

	var MSum *pbc.Element
	var metaSum *pbc.Element
	hashedFile := tool.FileToMMData(storage)
	for i := 0; i < challen.C; i++ {
		M := pairing.NewG1().SetBytes(hashedFile[aTable[i]])
		meta := pairing.NewG1().SetBytes(storage.MetaData[aTable[i]])

		if i == 0 {
			metaSum = pairing.NewG1().PowZn(meta, vTable[i])
			MSum = pairing.NewG1().PowZn(M, vTable[i])
		} else {
			MSum.Mul(MSum, pairing.NewG1().PowZn(M, vTable[i]))
			metaSum.Mul(metaSum, pairing.NewG1().PowZn(meta, vTable[i]))
		}
	}
	uProof := pairing.NewG1().PowZn(u, proof)
	right_hand := pairing.NewG1().Mul(uProof, MSum)
	pairing_left := pairing.NewGT().Pair(metaSum, g)
	pairing_right := pairing.NewGT().Pair(right_hand, pubKey)
	DetupVerifyResult := 0
	if pairing_left.Equals(pairing_right) {
		DetupVerifyResult = 1
	}
	return DetupVerifyResult
}

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