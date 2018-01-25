package main

import (
	"fmt"

	"github.com/anhle128/gblockchain/models"
)

func main() {
	bc := models.InitBlockchain()

	bc.AddBlock("Mot Con Vit Xoe Ra Hai Cai Canh")
	bc.AddBlock("No Keu Rang Quoac Quoac Quoac")
	bc.AddBlock("Quoac Quoac Quoac")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
