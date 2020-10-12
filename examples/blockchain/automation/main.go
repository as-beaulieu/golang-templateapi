package main

import (
	"TemplateApi/examples/blockchain/blockV1"
	"TemplateApi/examples/blockchain/blockV2"
	"TemplateApi/examples/blockchain/blockV3"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
)

type Commandline struct {
	blockchain *blockV3.BlockChain
}

func (cli *Commandline) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" add -block BLOCK_DATA - add a block to the chain")
	fmt.Println(" print - Prints the blocks in the chain")
}

func (cli *Commandline) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit() //exits the application by shutting down the goroutines
	}
}

func (cli *Commandline) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Added Block!")
}

func (cli *Commandline) printChain() {
	iterator := cli.blockchain.Iterator()

	for {
		block := iterator.Next()

		fmt.Println("*****")
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Block Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockV3.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("***")
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *Commandline) run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		if err := addBlockCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "print":
		if err := printChainCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func main() {
	//lesson1()

	//startTime := time.Now()
	//lessonV2()
	//elapsed := time.Since(startTime)
	//fmt.Printf("Finished! Application took %s\n", elapsed)

	lessonV3()
}

func lessonV3() {
	defer os.Exit(0)
	chain := blockV3.InitBlockChain()
	defer chain.Database.Close()

	cli := Commandline{chain}
	cli.run()
}

func lessonV2() {
	chain := blockV2.InitBlockChain()

	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	//To see the blockchain, lets iterate and print each
	for _, block := range chain.Blocks {
		fmt.Println("*****")
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Block Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockV2.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("***")
		fmt.Println()
	}
}

func lesson1() {
	chain := blockV1.InitBlockChain()

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
