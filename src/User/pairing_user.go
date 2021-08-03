package User

import (
	"encoding/binary"
	"pairing_test/src/tool"
	"reflect"

	"github.com/Nik-U/pbc"
)

// Alice generates a keypair and signs a message
func Upload(para tool.Params, inputData tool.DataList, privKeyByte []byte) ([][]byte, []byte) {
	pairing, _ := pbc.NewPairingFromString(para.Pairing)
	u := pairing.NewG1().SetBytes(para.U)
	n := 10
	// 鍵ペア生成
	privKey := pairing.NewZr().SetBytes(privKeyByte)

	// 署名　sk・H(m)を生成
	inputFile := tool.ReadBinaryFile(inputData.DataName, binary.BigEndian)
	splitedFile, _ := tool.SplitSlice(inputFile, n)
	var metaData [][]byte
	var mData [][]byte
	var MData [][]byte
	for i := 0; i < len(splitedFile); i++ {
		m := pairing.NewG1().SetFromHash(splitedFile[i])
		mData = append(mData, m.Bytes())

		mm := tool.GetBinaryBySHA256(m.X().String())
		M := pairing.NewG1().SetFromHash(mm) 

		um := pairing.NewG1().PowBig(u, m.X())
		temp := pairing.NewG1().Mul(um, M)
		meta := pairing.NewG1().MulZn(temp, privKey)

		metaData = append(metaData, meta.Bytes())
		MData = append(MData, M.Bytes())
	}
	return metaData, inputFile
}

func OutSource(file tool.Storage) tool.Storage {
	tool.SaveData(file)
	return file
}

func DedupProofGen(fileData tool.Storage, params tool.Params, challen tool.Chal) ([]byte, int) {
	pairing, _ := pbc.NewPairingFromString(params.Pairing)
	var mData []*pbc.Element
	var MData [][]byte
	dataBlockNum := len(fileData.MetaData)
	splitedFile, _ := tool.SplitSlice(fileData.File, dataBlockNum)
	for i := 0; i < len(splitedFile); i++ {
		m := pairing.NewG1().SetFromHash(splitedFile[i])
		mm := tool.GetBinaryBySHA256(m.X().String())
		M := pairing.NewG1().SetFromHash(mm)
		mData = append(mData, m)
		MData = append(MData, M.Bytes())
	}
	fitNum := tool.GetFITInfoFromUser(MData)
	aTable, vTable := tool.HashChallen(dataBlockNum, challen, pairing)

	var proofs *pbc.Element
	for j := 0; j < challen.C; j++ {
		if j == 0 {
			proofs = pairing.NewZr().MulBig(vTable[j], mData[aTable[j]].X())
		} else {
			proofs = pairing.NewZr().Add(proofs, pairing.NewZr().MulBig(vTable[j], mData[aTable[j]].X()))
		}
	}

	return proofs.Bytes(), fitNum
}

func AuditChallen(sharedParams string) []tool.Chal {
	fit := tool.InputFIT()
	var challens []tool.Chal
	for i := 0; i < len(fit); i++ {
		pairing, _ := pbc.NewPairingFromString(sharedParams)
		ck, kt1, kt2 := tool.ChallenGen(len(fit[i].HashedFile), pairing)
		challen := tool.Chal{ID: fit[i].FileId,C: ck, K1: kt1.Bytes(), K2: kt2.Bytes()}
		challens = append(challens, challen)
	}
	return challens
}

func CheckDep(outSourceData tool.Storage) int {
	found := 0
	fit := tool.InputFIT()
	for i := 0; i < len(fit); i++ {
		hashedFile := tool.FileToMMData(outSourceData)
		if reflect.DeepEqual(fit[i].HashedFile, hashedFile) {
			found = 1
		}
	}
	return found
}

func CheckLog(user int, params tool.Params) int {
	result := 0
	fit := tool.InputFIT()
	var linkIdTable []int

	//FIT
	for i := 0; i < len(fit); i++ {
		if tool.Contains(fit[i].UserID, user) {
			linkIdTable = append(linkIdTable, fit[i].FileId)
		}
	}
	//Para
	pairing, _ := pbc.NewPairingFromString(params.Pairing)
	u := pairing.NewG1().SetBytes(params.U)
	g := pairing.NewG1().SetBytes(params.G)
	logs := tool.ReadLogs()
	var challen tool.Chal
	var hashedFile [][]byte
	var pubKey *pbc.Element
	for j := 0; j < len(logs); j++ {
		if tool.Contains(linkIdTable, logs[j].FileId) {
			proof := pairing.NewZr().SetBytes(logs[j].Proof.Myu)
			uProof := pairing.NewG1().MulZn(u, proof)
			challen.C = logs[j].Challen.C
			challen.K1 = logs[j].Challen.K1
			challen.K2 = logs[j].Challen.K2
			gammaT := pairing.NewG1().SetBytes(logs[j].Proof.Gamma)
			aTable, vTable := tool.HashChallen(10, challen, pairing)
			for _, v := range fit {
				if v.FileId == logs[j].FileId {
					hashedFile = v.HashedFile
					pubKey = pairing.NewG1().SetBytes(params.PubKeys[v.UserID[0]].PubKey)
					break
				}
			}
			x := uProof
			for i := 0; i < challen.C; i++ {
				M := pairing.NewG1().SetBytes(hashedFile[aTable[i]])
				Mv := pairing.NewG1().MulZn(M, vTable[i])
				x = pairing.NewG1().Mul(x, Mv)
			}
			pairing_left := pairing.NewGT().Pair(gammaT, g)
			pairing_right := pairing.NewGT().Pair(x, pubKey)
			if pairing_left.Equals(pairing_right) {
				result = 1
			}
		}
	}
	return result
}
