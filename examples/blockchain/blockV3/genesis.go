package blockV3

//Genesis creates the first block in a blockchain, so that others may be added in with a reference prevHash
func Genesis() *Block {
	//Simple implementation: only data will be "genesis" as a string, and an empty byte slice for the prevHash
	return CreateBlock("Genesis", []byte{})
}
