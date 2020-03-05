package blockchain_test

import (
	"testing"
	"blockchain/blockchain"
)

func TestBlockChain(t *testing.T) {

	blockChain := blockchain.NewBlockchain()
	if blockChain == nil {
		t.Fail()
	}
}

func TestAddBlock(t *testing.T) {
	blockchain := blockchain.NewBlockchain()
	blockchain.AddBlock("Test2")
	if blockchain == nil {
		t.Fail()
	}
}
