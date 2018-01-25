package models

import (
	"time"
)

// Block model
type Block struct {
	Timestamp     int64  // header
	PrevBlockHash []byte // header
	Hash          []byte // header
	Nonce         int    // header

	Data []byte
}

// NewBlock create new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().UTC().Unix(),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Data:          []byte(data),
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock new genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
