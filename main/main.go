package main

import (
	"blockchain/blockchain"
	"blockchain/cli"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.Db.Close()

	cli := cli.CLI{Bc: bc}
	cli.Run()
}
