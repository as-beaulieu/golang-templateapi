package blockV1

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
