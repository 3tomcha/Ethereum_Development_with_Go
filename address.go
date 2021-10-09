package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	address := common.HexToAddress("0xc8F595E2084DB484f8A80109101D58625223b7C9")
	fmt.Println(address.Hex())
	fmt.Println(address.Hash().Hex())
	fmt.Println(address.Bytes())
}
