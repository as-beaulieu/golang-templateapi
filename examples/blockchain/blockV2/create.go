package blockV2

//CreateBlock takes data and the previous hash from last block and outputs pointer to a new block
func CreateBlock(data string, prevHash []byte) *Block {
	//using block constructor
	block := &Block{[]byte{}, []byte(data), prevHash, 0} //Initial nonce will be 0

	//Modifies V2 to use Proof of Work on each block
	pow := NewProof(block)
	nonce, hash := pow.Run()

	//Now take nonce and block and place into the block structure
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
