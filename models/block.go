package models

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block model
type Block struct {
	Timestamp     int64  // header
	PrevBlockHash []byte // header
	Hash          []byte // header

	Data []byte
}

// SetHash set hash for block
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock create new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().UTC().Unix(),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Data:          []byte(data),
	}
	block.SetHash()
	return block
}

// NewGenesisBlock new genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
