package User

import (
	"encoding/binary"
	"pairing_test/src/tool"

	"github.com/Nik-U/pbc"
)

// Alice generates a keypair and signs a message
func Upload(para tool.Params, inputData tool.DataList, privKeyByte []byte) ([][]byte, [][]byte, [][]byte, int, [][]byte) {
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
	return metaData, mData, MData, n, splitedFile
}

func OutSource(file, metadata [][]byte, fileName string, input tool.DataList){
	hashedFile ,_:= tool.HashFile(input.DataName)
	tool.SaveData(file, metadata, fileName, hashedFile)

}

func DedupProofGen(privKeyByte []byte, fileData tool.FileData, params tool.Params, inputData tool.DataList, challen tool.Chal) []byte {
	pairing, _ := pbc.NewPairingFromString(params.Pairing)
	var mData []*pbc.Element
	dataBlockNum := fileData.DataBlockNum
	inputFile := tool.ReadBinaryFile(inputData.DataName, binary.BigEndian)
	splitedFile, _ := tool.SplitSlice(inputFile, dataBlockNum)
	for i := 0; i < len(splitedFile); i++ {
		m := pairing.NewG1().SetFromHash(splitedFile[i])
		mData = append(mData, m)
	}
	aTable, vTable := tool.HashChallen(dataBlockNum, challen, pairing)

	var proofs *pbc.Element
	for j := 0; j < challen.C; j++ {
		if j == 0 {
			proofs = pairing.NewZr().MulBig(vTable[j], mData[aTable[j]].X())
		} else {
			proofs = pairing.NewZr().Add(proofs, pairing.NewZr().MulBig(vTable[j], mData[aTable[j]].X()))
		}
	}
	return proofs.Bytes()
}

func AuditChallen(sharedParams string, n int) tool.Chal {
	pairing, _ := pbc.NewPairingFromString(sharedParams)
	ck, kt1, kt2 := tool.ChallenGen(n, pairing)
	challen := tool.Chal{ck, kt1.Bytes(), kt2.Bytes()}
	return challen
}

func CheckDep(fileData tool.DataList, fileTableData []tool.FileIndexTable) int {
	found := 0
	for i := 0; i < len(fileTableData); i++ {
		hashedFile, err := tool.HashFile(fileData.DataName)
		if err != nil {
			err.Error()
		}
		if fileTableData[i].HashedFile == hashedFile {
			found = 1
		}
	}
	return found
}
