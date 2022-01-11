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

func GetParam(conn *contracts.Contracts) contracts.IndexTablePara {
	reply2, err := conn.GetParam(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(reply2)
    return reply2
}

func GetPubkey(userAddressStr string) []byte {
	userAddress := StringToAddress(userAddressStr)
	conn, _ := ConnectNetWork()
	reply2, err := conn.GetPubkey(&bind.CallOpts{}, userAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(reply2)
    return reply2
}

func GetArtLogs() []contracts.IndexTableArtLogTable {
	conn, _ := ConnectNetWork()
	reply2, err := conn.GetArtLogs(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(reply2)
    return reply2
}

func GetArtLog(artId []byte) contracts.IndexTableArtLogTable {
	conn, _ := ConnectNetWork()
	reply2, err := conn.GetArtLog(&bind.CallOpts{}, artId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(reply2)
    return reply2
}

func GetAuditLogs() []contracts.IndexTableAuditLogTable {
	conn, _ := ConnectNetWork()
	reply2, err := conn.GetAuditLogs(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(reply2)
    return reply2
}

func GetAuditLog(artId []byte) contracts.IndexTableAuditLogTable {
	conn, _ := ConnectNetWork()
	reply2, err := conn.GetAuditLog(&bind.CallOpts{}, artId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(reply2)
    return reply2
}

func SetArtLog(privKey string, owner string, hashedData [][]byte, artId []byte) error {
    conn, client := ConnectNetWork()
	auth := AuthUser(client, privKey)
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
    if err != nil {
        // .env読めなかった場合の処理
    }
	userAddressStr := os.Getenv("USER_ADDRESS")
	userAddress := common.HexToAddress(userAddressStr)
	_, err = conn.SetArtLog(auth, hashedData, artId , userAddress)
    if err != nil{
		log.Fatal(err)
	}
    return err
}


func SetAuditChal(privKey string, chal uint32, k1 []byte, k2 []byte, logId []byte, artId []byte) error {
	conn, client := ConnectNetWork()
	auth := AuthUser(client, privKey)
	_, err := conn.SetAuditChal(auth, chal, k1, k2, logId, artId)
	if err != nil{
		log.Fatal(err)
	}
	return err
}


func SetAuditProof(privKey string, myu []byte, gamma []byte, artId []byte) error {
    conn, client := ConnectNetWork()
	auth := AuthUser(client, privKey)
	_, err := conn.SetAuditProof(auth, myu, gamma, artId)
    if err != nil{
		log.Fatal(err)
	}
    return err
}

func SetAuditResult(privKey string, result bool, artId []byte) error {
    conn, client := ConnectNetWork()
	auth := AuthUser(client, privKey)
	_, err := conn.SetAuditResult(auth, result, artId)
    if err != nil{
		log.Fatal(err)
	}
    return err
}

func RegisterPubKey(privKey string, pubkey []byte) error {
    conn, client := ConnectNetWork()
	auth := AuthUser(client, privKey)
	_, err := conn.RegisterPubKey(auth, pubkey)
    if err != nil{
		log.Fatal(err)
	}
    return err
}

func StringToAddress(addressStr string) common.Address {
	return common.HexToAddress(addressStr)
}
