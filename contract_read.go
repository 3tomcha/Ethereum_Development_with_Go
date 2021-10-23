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

	address := common.HexToAddress("0xe13cc8daaef5c8e61e91f67dda0ee83f3e1efe73")
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
