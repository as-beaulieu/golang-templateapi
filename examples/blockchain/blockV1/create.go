package blockV1

//CreateBlock takes data and the previous hash from last block and outputs pointer to a new block
func CreateBlock(data string, prevHash []byte) *Block {
	//using block constructor
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}
