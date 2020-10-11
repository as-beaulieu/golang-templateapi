package blockchain

//BlockChain is the struct that represents the whole blockchain
type BlockChain struct {
	Blocks []*Block
	//Will use more complicated implements will come
}

type Block struct {
	Hash     []byte //Derive hash from Data and PrevHash
	Data     []byte //This is the actual data. Can be anything from ledgers to documents, images, etc
	PrevHash []byte //Last block's hash. Links blocks together like a backlinked list
	//When the block gets more complicated, adding more things like timestamp, blockheight, and other fields
}
