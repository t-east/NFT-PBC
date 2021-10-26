package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	// this would be your generated smart contract bindings

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"pairing_test/src/ethereum/contracts"
	"pairing_test/src/SetUp"
)

const PRIVATE_KEY = "001eec6f80d346f468d032964b34995c9c38c6cabc875fe7dd2a5baa8056a880"
const SP_ADDRESS = "0xF69Fc8A0AA6b2d0f3baFA7d40ee501a788B0d65E"
const TPA_ADDRESS = "0x2dFce98b4eC6A9E8b719C96B8b21CB75bf0c82d7"

func main() {
	client, err := ethclient.Dial("http://0.0.0.0:8545")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		panic(err)
	}

	spAddress := common.HexToAddress(SP_ADDRESS)
	tpaAddress := common.HexToAddress(TPA_ADDRESS)
	pairing, g, u := SetUp.SetUp()

	publicKey := privateKey.Public()
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

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(100)

	address, tx, instance, err := contracts.DeployContracts(auth, client, spAddress, tpaAddress, pairing, g, u)
	if err != nil {
		panic(err)
	}

	fmt.Println(address.Hex())

	_, _ = instance, tx
}