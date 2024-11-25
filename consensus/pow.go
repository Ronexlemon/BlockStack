package consensus

import (
	//"bytes"

	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"

	
	"github.com/RonexLemon/chain/types"
)




func  MineBlock( block *types.Block)[]byte{
	var hashit [32]byte
	for {
		data := bytes.Join([][]byte{
			[]byte(strconv.Itoa(int( block.Height))),
			[]byte(block.PrevHash),

			
			[]byte(block.PrevHash),
			[]byte(strconv.Itoa(int(block.Timestamp))),
			[]byte(strconv.Itoa(int(block.Nonce))),
			
			
		}, []byte{})
		hashit = sha256.Sum256(data)

		// Hash the data
		// Example: Find hash starting with two leading zeros.
		if hashit[0] == 0 && hashit[1] == 0 {
			fmt.Printf("Block mined! Nonce: %d, Hash: %x\n", block.Nonce, hashit)
			break
		}
		block.Nonce++

	}
	return hashit[:]
	

}