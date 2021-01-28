package blockV4

//Genesis creates the first block in a blockchain, so that others may be added in with a reference prevHash
func Genesis(coinbase *Transaction) *Block {
	//Simple implementation: only data will be "genesis" as a string, and an empty byte slice for the prevHash
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}
