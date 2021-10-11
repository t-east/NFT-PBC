package ethhandler

import (
	"fmt"
    "context"
    "crypto/ecdsa"
    "log"
    "math/big"    
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
	CONTRACT_ADDRESS = "0xADEb748E7BbB603f8Bc6dF32AB726d3153Cf58E3"
)

func ConnectNetWork() (*contracts.Contracts, *ethclient.Client) {
	client, err := ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%s", GANACHE_PORT))
	if err != nil {
		panic(err)
	}
	conn, err := contracts.NewContracts(common.HexToAddress(CONTRACT_ADDRESS), client)
	if err != nil {
		panic(err)
	}
	return conn, client
}

func AuthUser(client *ethclient.Client) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA("5f4a70a398a494d4ec34c8fd2e8986c306fc3f37ebe228cb2b273c70fe16ed7f")
    if err != nil {
        log.Fatal(err)
    }
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA) 
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice

	return auth
}

func Set(auth *bind.TransactOpts, conn *contracts.Contracts,message string) {
	_, err := conn.Set(auth, message)
	if err != nil {
		fmt.Println(err)
	}
}

func Get(conn *contracts.Contracts) {
	reply2, err := conn.Get(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(reply2)
}

func StringToAddress(addressStr string) common.Address {
	return common.HexToAddress(addressStr)
}