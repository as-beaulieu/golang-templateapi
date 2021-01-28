package blockV3

import (
	"math/big"
)

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty)) //shifts the number of bytes over by this number

	pow := &ProofOfWork{b, target}

	return pow
}
