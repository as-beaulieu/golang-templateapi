package blockV5

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"log"
	"os"
	"runtime"
)

func InitBlockChain(address string) *BlockChain {
	var lastHash []byte

	if DbExists() == true {
		fmt.Println("Blockchain already exists")
		runtime.Goexit()
	}

	opts := badger.DefaultOptions(dbPath)

	db, err := badger.Open(opts)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(transaction *badger.Txn) error {
		coinbaseTxn := CoinbaseTx(address, genesisData)
		genesis := Genesis(coinbaseTxn)
		fmt.Println("Genesis Proved")
		if err = transaction.Set(genesis.Hash, genesis.Serialize()); err != nil {
			log.Panic(err)
		} //Hash is the key, and serialize the whole block

		if err = transaction.Set([]byte("lh"), genesis.Hash); err != nil {
			log.Panic(err)
		}

		lastHash = genesis.Hash

		return err
	})

	if err != nil {
		log.Panic(err)
	}

	return &BlockChain{lastHash, db} //new blockchain in memory
}

func DbExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}

func ContinueBlockChain(address string) *BlockChain {
	if DbExists() == false {
		fmt.Println("No existing blockchain found, create one!")
		runtime.Goexit()
	}

	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)

	db, err := badger.Open(opts)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		if err != nil {
			log.Panic(err)
		}

		err = item.Value(func(val []byte) error {
			lastHash = append([]byte{}, val...)
			return nil
		})

		return err
	})
	if err != nil {
		log.Panic(err)
	}

	return &BlockChain{lastHash, db}
}
