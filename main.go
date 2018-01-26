package main

import (
	"github.com/anhle128/gblockchain/cli"
	"github.com/anhle128/gblockchain/models"
)

func main() {
	bc := models.InitBlockchain()
	defer bc.CloseDB()

	cli := cli.InitCLI(bc)
	cli.Run()
}
