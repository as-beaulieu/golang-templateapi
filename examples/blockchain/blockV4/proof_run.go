package blockV4

import (
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		//prepare our data
		data := pow.InitData(nonce)

		//convert into sha256 hash
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash) //show that hash is changing until the appropriate hash is found

		//convert hash into big int
		intHash.SetBytes(hash[:])

		//compare big int with proof of work target
		if intHash.Cmp(pow.Target) == -1 {
			break //means that hash is less than target, and that we've signed the block
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}
