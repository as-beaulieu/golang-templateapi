package blockV4

import (
	"crypto/sha256"
	"math/big"
)

//Idea is that after we run our Proof of Work's Run(), we'll have our nonce will will derive the hash, which met the target we wanted
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}
