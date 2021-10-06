package main

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
	"pairing_test/src/contracts"
)

// ここの２つの変数を修正して実行する
const (
	// ganacheの起動したときのポートを指定 (8545 か 7545)
	GANACHE_PORT = "8545"
	// 先ほど作成したプログラムから取得した。　CONTRACT_ADDRESSを取得
	CONTRACT_ADDRESS = "0x439f1A49Db8e25528Dd6d2Cb4B27E60Ec309Ea07"
)

func main() {
	client, err := ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%s", GANACHE_PORT))
	if err != nil {
		panic(err)
	}
	conn, err := contracts.NewContracts(common.HexToAddress(CONTRACT_ADDRESS), client)
	if err != nil {
		panic(err)
	}

    privateKey, err := crypto.HexToECDSA("75ecb40e80fa420b747b675df970f66d510153f6e374aa8b4cb43bc494d0c816")
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

	_, err = conn.Set(auth, "cds")
	if err != nil {
		fmt.Printf("era")
	}

	reply2, err := conn.Get(&bind.CallOpts{})
	if err != nil {
		fmt.Printf("era")
	}
	fmt.Print(reply2)
}

func ConnectNetWork() contracts.Contracts {
	client, err := ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%s", GANACHE_PORT))
	if err != nil {
		panic(err)
	}
	conn, err := contracts.NewContracts(common.HexToAddress(CONTRACT_ADDRESS), client)
	if err != nil {
		panic(err)
	}
	return conn
}

func AuthUser() {
	privateKey, err := crypto.HexToECDSA("75ecb40e80fa420b747b675df970f66d510153f6e374aa8b4cb43bc494d0c816")
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

func Set(auth bind.TransactOpts, message string) {
	_, err = conn.Set(auth, message)
	if err != nil {
		fmt.Printf(err)
	}
}

func Get() {
	reply2, err := conn.Get(&bind.CallOpts{})
	if err != nil {
		fmt.Printf(err)
	}
	fmt.Print(reply2)
}