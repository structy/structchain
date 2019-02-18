package main

import (
	"fmt"

    "github.com/structy/structchain"
)

func main() {
	bc := structchain.NewBlockchain()

	bc.AddBlock("Send 1 LVC to Avelino")
	bc.AddBlock("Send 2 more LVC to Avelino")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
