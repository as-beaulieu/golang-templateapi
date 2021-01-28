package blockV5

import (
	"github.com/dgraph-io/badger"
	"log"
)

func (chain *BlockChain) AddBlock(transactions []*Transaction) {
	var lastHash []byte

	err := chain.Database.View(func(transaction *badger.Txn) error {
		item, err := transaction.Get([]byte("lh"))
		if err != nil {
			log.Panic(err)
		}
		err = item.Value(func(val []byte) error {
			lastHash = append([]byte{}, val...)
			return nil
		})
		if err != nil {
			log.Panic(err)
		}

		newBlock := CreateBlock(transactions, lastHash)

		err = chain.Database.Update(func(transaction *badger.Txn) error {
			if err := transaction.Set(newBlock.Hash, newBlock.Serialize()); err != nil {
				log.Panic(err)
			}

			if err = transaction.Set([]byte("lh"), newBlock.Hash); err != nil { //Set the new blocks hash as our latest lastHash
				log.Panic(err)
			}

			chain.LastHash = newBlock.Hash

			return err
		})
		if err != nil {
			log.Panic(err)
		}

		return err
	})

	if err != nil {
		log.Panic(err)
	}
}
