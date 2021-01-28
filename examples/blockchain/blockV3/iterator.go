package blockV3

import (
	"github.com/dgraph-io/badger"
	"log"
)

func (chain *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{chain.LastHash, chain.Database}
}

//Because we start with the BlockChain's LastHash, we're iterating backwards through the blocks (Newest -> Genesis)

func (iterator *BlockChainIterator) Next() *Block {
	var block *Block

	err := iterator.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iterator.CurrentHash)
		if err != nil {
			log.Panic(err)
		}
		var encodedBlock []byte
		err = item.Value(func(val []byte) error {
			encodedBlock = append([]byte{}, val...)
			return nil
		})
		if err != nil {
			log.Panic(err)
		}

		block = Deserialize(encodedBlock)

		return err
	})
	if err != nil {
		log.Panic(err)
	}

	iterator.CurrentHash = block.PrevHash //because each block points to its previous block, this sets the next step in the iterator

	return block
}
