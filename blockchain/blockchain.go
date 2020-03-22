package blockchain

import (
	"blockchain/block"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// Blockchain struct
type Blockchain struct {
	tip []byte
	Db  *bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// Iterator constructor to return and Iterator object
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.Db}

	return bci
}

// Next method to return next block in blockchain
func (bi *BlockchainIterator) Next() *block.Block {
	var returnedBlock *block.Block

	error := bi.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		blockData := bucket.Get(bi.currentHash)
		returnedBlock = block.DeserializeBlock(blockData)

		return nil
	})

	if error != nil {
		panic("Error retrieving the next block")
	}

	bi.currentHash = returnedBlock.PrevBlockHash

	return returnedBlock
}

// AddBlock method to add a block to blockchain
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.Db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		lastHash = bucket.Get([]byte("l"))

		return nil
	})

	if err != nil {
		panic("Porblem retrieving info from database")
	}

	newBlock := block.NewBlock(data, lastHash)

	_ = bc.Db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put(newBlock.Hash, newBlock.Serialize())
		err = bucket.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash

		if err != nil {
			panic("Porblem retrieving info from database")
		}

		return nil
	})

}

// NewBlockchain create a new blockchain
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))

		if bucket == nil {
			genesis := block.NewGenesisBlock()
			b, _ := tx.CreateBucket([]byte(blocksBucket))
			err = b.Put(genesis.Hash, genesis.Serialize())
			// Store last block reference
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = bucket.Get([]byte("l"))
		}

		if err == nil {
			return nil
		} else {
			return err
		}

	})

	bc := Blockchain{tip, db}

	return &bc

}
