package blockchain

import (
	"blockchain/block"
)

// Blockchain struct
type Blockchain struct {
	Blocks []*block.Block
}

// AddBlock method to add a block to blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockchain create a new blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*block.Block{block.NewGenesisBlock()}}
}
