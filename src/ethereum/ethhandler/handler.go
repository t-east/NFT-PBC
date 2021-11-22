package ethhandler

import (
	"fmt"
    "context"
    "crypto/ecdsa"
    "log"
    "os"
    "math/big"   
    "github.com/joho/godotenv" 
    "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"pairing_test/src/ethereum/contracts"
)

// ここの２つの変数を修正して実行する
const (
	GANACHE_PORT = "8545"
)

func ConnectNetWork() (*contracts.Contracts, *ethclient.Client) {
    err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("get-go-env-error")
	}
	contractAffress := os.Getenv("CONTRACT_ADDRESS")
	client, err := ethclient.Dial(fmt.Sprintf("http://gana:%s", GANACHE_PORT))
	if err != nil {
		panic(err)
	}
	conn, err := contracts.NewContracts(common.HexToAddress(contractAffress), client)
	if err != nil {
		panic(err)
	}
	return conn, client
}

func GetUserAddress(privKey string) common.Address {
    fmt.Println(privKey)
    err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("go_env-error")
	}
	userPrivKeyStr := os.Getenv("USER_PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(userPrivKeyStr)
    if err != nil {
        log.Fatal("error!!")
    }
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }
    Address := crypto.PubkeyToAddress(*publicKeyECDSA) 
	return Address
}

func AuthUser(client *ethclient.Client, privKey string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(privKey)
    if err != nil {
        log.Fatal("convert-privKey-error")
    }
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA) 
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal("nonce-error")
    }
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal("gas-error")
    }
    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice

	return auth
}

func GetPara(conn *contracts.Contracts) contracts.IndexTablePara {
	reply2, err := conn.GetParam(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(reply2)
    return reply2
}

func GetArtIds(conn *contracts.Contracts) []string {
	reply2, err := conn.GetParam(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(reply2)
    a := [] string{"Golang", "Java"}
    return a
}

func RegisterArt(privKey string, owner string, hashedData [][]byte, artId []byte) error {
    conn, client := ConnectNetWork()
	auth := AuthUser(client, privKey)
	userAddress := StringToAddress(owner)
	reply, err := conn.RegisterArt(auth, hashedData, artId , userAddress)
    if err != nil{
		log.Fatal(err)
	}
	log.Print(reply)
    return err
}

func SetProof(privKey string, myu []byte, gamma []byte, artId []byte, logId []byte) error {
    conn, client := ConnectNetWork()
	auth := AuthUser(client, privKey)
	reply, err := conn.RegisterAuditProof(auth, myu, gamma, artId, logId)
    if err != nil{
		log.Fatal(err)
	}
	log.Print(reply)
    return err
}

func GetHashData(artId []byte) [][]byte {
	conn, _ := ConnectNetWork()
	reply, err := conn.GetHashedFile(&bind.CallOpts{}, artId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(reply)
    return reply
}

func GetPubkey(artId []byte) []byte {
	conn, _ := ConnectNetWork()
	reply, err := conn.GetPubKey(&bind.CallOpts{}, artId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(reply)
    return reply
}

func GetLog() []contracts.IndexTableLogTable {
	conn, _ := ConnectNetWork()
	logs, _ := conn.GetLog(&bind.CallOpts{})
    return logs
}

func StringToAddress(addressStr string) common.Address {
	return common.HexToAddress(addressStr)
}
