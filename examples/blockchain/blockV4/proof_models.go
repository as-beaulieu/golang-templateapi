package blockV4

import (
	"math/big"
)

//Take data from the block

// Create a counter (nonce) which starts at 0, increments upwards, theoretically infinitely

// Create a hash of the data plus the counter

// Check the hash to see if it meets a set of requirements
//	If meets the requirements, then we use the hash, and says it signs the block
//		If does not meet requirements, go back make another hash, and repeat until we have a hash that meets requirements

//	Requirements:
//		First few bytes must contain 0s
//		Bitcoin original difficulty - first 20 digits was 0

//In v1, difficulty is static. In real life, difficulty increases
//	Ideally, you want the time to compute a block to remain the same,

const Difficulty = 15 //Start this at 12, note the change in compute time just to 15, 18, etc.
//12 -> 143.434 ms
//15 -> 1.30096 s
//18 -> 6.18875 s
//21 -> 50.7466 s

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}
