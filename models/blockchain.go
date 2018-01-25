package models

// Blockchain model
type Blockchain struct {
	Blocks []*Block
}

// AddBlock add new block to blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newblock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newblock)
}

// InitBlockchain init new block chain
func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
