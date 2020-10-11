package blockchain

import (
	"bytes"
	"crypto/sha256"
)

func (b *Block) DeriveHash() {
	//Take a 2D slice, combine with an empty slice of bytes
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info) //sha 256 is fairly simple compared to real way to hash for a blockchain
	b.Hash = hash[:]
}
