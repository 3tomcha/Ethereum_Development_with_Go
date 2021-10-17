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

	address := common.HexToAddress("0xb1ed820f4102143ea965d8301c10610488940d46")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}
