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
	bigBang := block.NewGenesisBlock()
	if bigBang == nil {
		t.Fail()
	}
}

func TestSerialize(t *testing.T) {
	bloc := block.NewBlock("test", []byte("test"))
	data := bloc.Serialize()
	if data == nil {
		t.Fail()
	}
	blockDes := block.DeserializeBlock(data)
	if len(blockDes.Data) != len([]byte("test")) {
		t.Fail()
	}
}
