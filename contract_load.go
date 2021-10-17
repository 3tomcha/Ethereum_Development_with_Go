package main

import (
	"fmt"
	"log"

	store "github.com/3tomcha/Ethereum_Development_with_Go/contracts" // for demo
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("aa3bc6ba81aee1901dd418507fbaa83478310eedae20f7346375f1929f7eaa04")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
