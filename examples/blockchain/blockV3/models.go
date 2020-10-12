package blockV3

import (
	"github.com/dgraph-io/badger"
)

const (
	dbPath = "/tmp/blocks"
)

//BlockChain is the struct that represents the whole blockchain
type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

//In V2, BlockChain was in memory, so was able to print the ProofOfWork effort on the CLI, but now it's persisted in DB
//BlockChainIterator helps to reclaim that feature
type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

type Block struct {
	Hash     []byte //Derive hash from Data and PrevHash
	Data     []byte //This is the actual data. Can be anything from ledgers to documents, images, etc
	PrevHash []byte //Last block's hash. Links blocks together like a backlinked list
	Nonce    int
}

//When the block gets more complicated, adding more things like timestamp, blockheight, and other fields
