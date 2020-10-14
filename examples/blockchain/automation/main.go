package main

import (
	cli2 "TemplateApi/examples/blockchain/cli"
	"fmt"
	"os"
	"time"
)

func main() {

	startTime := time.Now()
	//lesson1()
	//lessonV2()
	//lessonV3()
	lessonV4()
	elapsed := time.Since(startTime)
	fmt.Printf("Finished! Application took %s\n", elapsed)

}

func lessonV4() {
	defer os.Exit(0)
	cli := cli2.Commandline{}
	cli.Run()
}

func lessonV3() {
	//defer os.Exit(0)
	//chain := blockV3.InitBlockChain()
	//defer chain.Database.Close()
	//
	//cli := Commandline{}
	//cli.run()
}

func lessonV2() {
	//chain := blockV2.InitBlockChain()
	//
	//chain.AddBlock("first block after genesis")
	//chain.AddBlock("second block after genesis")
	//chain.AddBlock("third block after genesis")
	//
	////To see the blockchain, lets iterate and print each
	//for _, block := range chain.Blocks {
	//	fmt.Println("*****")
	//	fmt.Printf("Previous hash: %x\n", block.PrevHash)
	//	fmt.Printf("Block Data: %s\n", block.Data)
	//	fmt.Printf("Hash: %x\n", block.Hash)
	//
	//	pow := blockV2.NewProof(block)
	//	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	//	fmt.Println("***")
	//	fmt.Println()
	//}
}

func lesson1() {
	//chain := blockV1.InitBlockChain()
	//
	//chain.AddBlock("first block after genesis")
	//chain.AddBlock("second block after genesis")
	//chain.AddBlock("third block after genesis")
	//
	////To see the blockchain, lets iterate and print each
	//for _, block := range chain.Blocks {
	//	fmt.Println("*****")
	//	fmt.Printf("Previous hash: %x\n", block.PrevHash)
	//	fmt.Printf("Block Data: %s\n", block.Data)
	//	fmt.Printf("Hash: %x\n", block.Hash)
	//	fmt.Println("***")
	//}
}
