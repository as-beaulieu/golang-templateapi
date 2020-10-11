package main

import (
	"TemplateApi/src/blockchain"
	"fmt"
)

func main() {
	lesson1()
}

func lesson1() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	//To see the blockchain, lets iterate and print each
	for _, block := range chain.Blocks {
		fmt.Println("*****")
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Block Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("***")
	}
}
