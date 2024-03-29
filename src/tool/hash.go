package tool

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"crypto/md5"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"os"

	"github.com/Nik-U/pbc"
)

func HashChallen(n int, c int, k1, k2 []byte, pairing *pbc.Pairing) ([]int, []*pbc.Element) {

	k1Big := new(big.Int).SetBytes(k1)
	k2Big := new(big.Int).SetBytes(k2)
	nBig := big.NewInt(int64(n))
	var aTable []int
	var vTable []*pbc.Element
	for i := 0; i < n; i++ {
		iBig := big.NewInt(int64(i+1))
		ik1Big := new(big.Int).Mod(new(big.Int).Mul(iBig, k1Big), nBig)
		aTable = append(aTable, int(ik1Big.Int64()))
	}
	for j := 0; j < c; j++ {
		iBig := big.NewInt(int64(j+1))
		ik2Big := pairing.NewZr().SetBig(new(big.Int).Mul(iBig, k2Big))
		vTable = append(vTable, ik2Big)
	}
	return aTable, vTable
}

func ReadBinaryFile(filename string, order binary.ByteOrder) []byte {
	// ファイルを開く
	file, _ := os.Open(filename)

	defer file.Close()
	fi, _ := file.Stat()
	//fmt.Printf("%s",file)
	fileSize := int(fi.Size())

	// ファイルから1バイト読み出し
	//bytesList := make([]byte, fileSize)
	b := make([]byte, fileSize)
	file.Read(b)
	return b
}

func HashFile(fileName string) (string, error) {
	inputFile := ReadBinaryFile(fileName, binary.BigEndian)
	if inputFile == nil {
		return "", errors.New("(*>△<)<ナーンナーンっっ")
	}
	r := sha256.Sum256(inputFile)
	return hex.EncodeToString(r[:]), nil
}

func SplitSlice(list []byte, size int) ([][]byte, error) {
	if size <= 0 {
		return nil, fmt.Errorf("size need positive number")
	}
	var result [][]byte
	var tmp = list
	splitNum := len(list) / size
	for i := 0; i < size; i++ {
		if i == size {
			if len(tmp) == 0 {
				break
			}
			r := sha256.Sum256(tmp[:])
			result = append(result, r[:])
		} else {
			r := sha256.Sum256(tmp[0:splitNum])
			result = append(result, r[:])
			tmp = tmp[splitNum:]
		}
	}
	return result, nil
}

func GetBinaryBySHA256(s string) []byte {
	r := sha256.Sum256([]byte(s))
	return r[:]
}
func MD5(s string) []byte {
	hash := md5.Sum([]byte(s))
   return []byte(hex.EncodeToString(hash[:]))
}

func ChallenGen(n int, pairing *pbc.Pairing) (int, *pbc.Element, *pbc.Element) {
	ck := rand.Intn(n) + 1
	k1 := pairing.NewZr().Rand()
	k2 := pairing.NewZr().Rand()
	return ck, k1, k2
}

func FileToMMData(file Storage)[][]byte{
	para, _ := InputPara()
	pairing, _ := pbc.NewPairingFromString(para.Pairing)
	splitedFile, _ := SplitSlice(file.File, len(file.MetaData))
	var MData [][]byte
		for i := 0; i < len(splitedFile); i++ {
			m := pairing.NewG1().SetFromHash(splitedFile[i])
			mm := GetBinaryBySHA256(m.X().String())
			M := pairing.NewG1().SetFromHash(mm)
			MData = append(MData, M.Bytes())
		}
	return MData
}