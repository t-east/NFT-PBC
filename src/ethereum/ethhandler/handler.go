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
	// ganacheの起動したときのポートを指定 (8545 か 7545)
	GANACHE_PORT = "8545"
	// 先ほど作成したプログラムから取得した。　CONTRACT_ADDRESSを取得
	CONTRACT_ADDRESS = "e5B3cc9306B0E565D481cB2fd674DB0baB3648E4"
    PPriv = "4bb7d1c7d4aa82cb5a94a869137d62909621e8761e9a4bbb00883ae55ac6a783"
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

func GetHashData(conn *contracts.Contracts, artId string) [][]byte {
	reply, err := conn.GetHashedFile(&bind.CallOpts{}, artId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(reply)
    return reply
}

func StringToAddress(addressStr string) common.Address {
	return common.HexToAddress(addressStr)
}
