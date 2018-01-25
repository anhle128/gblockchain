package main

import (
	"fmt"

	"github.com/anhle128/gblockchain/models"
)

func main() {
	bc := models.InitBlockchain()

	bc.AddBlock("Send 1 BTC to AnhLe")
	bc.AddBlock("Send 2 more BTC to AnhLe")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
