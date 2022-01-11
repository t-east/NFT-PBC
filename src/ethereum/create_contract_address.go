package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	// this would be your generated smart contract bindings

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"pairing_test/src/ethereum/contracts"
	"pairing_test/src/SetUp"
)

func main() {
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
    if err != nil {
        // .env読めなかった場合の処理
    }
	setupPrivateKeyStr := os.Getenv("SETUP_PRIVATE_KEY")
	spAddressStr := os.Getenv("SP_ADDRESS")
	tpaAddressStr := os.Getenv("TPA_ADDRESS")
	userAddressStr := os.Getenv("USER_ADDRESS")
	client, err := ethclient.Dial("http://gana:8545")
	if err != nil {
		panic(err)
	}

	setupPrivateKey, err := crypto.HexToECDSA(setupPrivateKeyStr)
	if err != nil {
		panic(err)
	}

	userAddress := common.HexToAddress(userAddressStr)
	spAddress := common.HexToAddress(spAddressStr)
	tpaAddress := common.HexToAddress(tpaAddressStr)
	pairing, g, u := SetUp.SetUp()

	publicKey := setupPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(setupPrivateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(100)

	address, tx, instance, err := contracts.DeployContracts(auth, client, userAddress, spAddress, tpaAddress, pairing, g, u)
	if err != nil {
		panic(err)
	}

	os.Setenv("CONTRACT_ADDRESS", address.Hex())
	aho := os.Getenv("CONTRACT_ADDRESS")
	fmt.Println(address.Hex())

	_, _ = instance, tx
	fmt.Println(spAddressStr)
	fmt.Println(aho)
}

func loadEnv() error {
	err := godotenv.Load("../.env")
	return err
}