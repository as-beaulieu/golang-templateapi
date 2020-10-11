package blockchain

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
