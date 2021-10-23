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

	address := common.HexToAddress("6f6a39c3dcc9adfa3551b6fd2990e54d628ea3518863dd8ce534c555872b463c")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
