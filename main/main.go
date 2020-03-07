package main

import (
	"blockchain/blockchain"
	"fmt"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Transaction 1")
	bc.AddBlock("Transaction 2")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
