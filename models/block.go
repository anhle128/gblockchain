package models

import (
	"bytes"
	"encoding/gob"
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

// Serialize block -> []byte
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	encoder.Encode(b)
	return result.Bytes()
}

// DeserializeBlock []byte -> block
func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	decoder.Decode(&block)
	return &block
}

// NewGenesisBlock new genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
