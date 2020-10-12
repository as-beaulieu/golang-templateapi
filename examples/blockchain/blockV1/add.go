package blockV1

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1] //get the previous block in the blockchain

	//create the new block
	new := CreateBlock(data, prevBlock.Hash)

	//append this new block to the blockchain
	chain.Blocks = append(chain.Blocks, new)
}
