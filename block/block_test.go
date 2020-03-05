package block_test

import (
	"blockchain/block"
	"testing"
)

func TestConstructor(t *testing.T) {
	block := block.NewBlock("Hi", []byte("test"))
	if block == nil {
		t.Fail()
	}
}

func TestGenesis(t *testing.T) {
	bigBang := block.GenesisBlock()
	if bigBang == nil {
		t.Fail()
	}
}
