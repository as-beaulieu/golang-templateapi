package blockV3

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"log"
)

func InitBlockChain() *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)

	db, err := badger.Open(opts)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(transaction *badger.Txn) error {
		if _, err := transaction.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found. Creating Genesis block")
			genesis := Genesis()
			fmt.Println("Genesis Proved")
			if err = transaction.Set(genesis.Hash, genesis.Serialize()); err != nil {
				log.Panic(err)
			} //Hash is the key, and serialize the whole block

			if err = transaction.Set([]byte("lh"), genesis.Hash); err != nil {
				log.Panic(err)
			}

			lastHash = genesis.Hash

			return err
		} else { //If we already have a database, and already has a blockchain inside
			fmt.Println("database found, getting lastHash block")
			item, err := transaction.Get([]byte("lh"))
			if err != nil {
				log.Panic(err)
			}
			err = item.Value(func(val []byte) error {
				lastHash = append([]byte{}, val...)
				return nil
			})
			return err
		}
	})

	if err != nil {
		log.Panic(err)
	}

	blockchain := BlockChain{lastHash, db} //new blockchain in memory
	return &blockchain
}
