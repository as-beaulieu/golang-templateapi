package blockV2

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
