package blockchain

import (
	"blockchain/block"
)

type Blockchain struct {
	blocks []*block.Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*block.Block{block.GenesisBlock()}}
}
