package cli

import (
	"blockchain/blockchain"
	"flag"
	"fmt"
	"os"
)

type CLI struct {
	Bc *blockchain.Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("addBlock: Add block to the chain")
	fmt.Println("printChain: Print block chain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.Bc.AddBlock(data)
	fmt.Println("Added block")
}

func (cli *CLI) printChain() {
	bci := cli.Bc.Iterator()

	for {
		block := bci.Next()
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			panic("Error in command")
		}
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			panic("Error in command")
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
