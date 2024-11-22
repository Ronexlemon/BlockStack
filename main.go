package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/RonexLemon/chain/block"
	"github.com/RonexLemon/chain/cryptohash"
	"github.com/RonexLemon/chain/nonce"
	"github.com/RonexLemon/chain/types"
)

func ReadData() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	return strings.TrimSpace(data)
}

func main() {
	inputChannel := make(chan string)
	done := make(chan bool)
	blockChain := block.BlockChain{}
	hashit := cryptohash.Hash{}
	nonceGenerator := nonce.Nonce{}
	
	// Create the genesis block and initialize the blockchain
	genesis := blockChain.NewGenisBlock()
	hashit.HashData(genesis) 
	blockChain.NewBlock(genesis)

	go func() {
		for {
			fmt.Println("Enter input (type 'next' to continue or 'exit' to quit): ")
			data := ReadData()

			// Create a new block using previous block's hash and nonce
			prevBlock := blockChain.GetRecentBlock()
			//generate nonce
			nonceGenerator.Generate()
			newBlock := types.Block{
				PrevHash:  prevBlock.Hash,  // Use the previous block's hash
				Nonce:     nonceGenerator.Nonce, // Get a new nonce
				Height:    prevBlock.Height + 1,  // Increment height
				Transactions: data,  // Use the input data as the transaction
				Timestamp: int64(time.Now().Unix()), // Get the current timestamp
			}

			// Calculate the hash for the new block
			hashVal := hashit.HashData(&newBlock)
			newBlock.Hash = hashVal // Set the calculated hash

			// Add the new block to the blockchain
			blockChain.NewBlock(&newBlock)

			// Output the block information
			fmt.Printf("PreviousHash: %s, Hash: %s, Timestamp: %d, Nonce: %d, Height: %d\n",
				newBlock.PrevHash, newBlock.Hash, newBlock.Timestamp, newBlock.Nonce, newBlock.Height)

			if data == "exit" {
				done <- true
				return
			}
			inputChannel <- data
		}
	}()

	// Main loop to process user inputs
	for {
		select {
		case input := <-inputChannel:
			fmt.Println("You have entered:", input)
		case <-done:
			fmt.Println("Exiting program")
			return
		}
	}
}
