package block

import (
	"time"

	"github.com/RonexLemon/chain/types"
)
type BlockChain struct {
    types.BlockChain
}

func(b *BlockChain) AddBlock(block *types.Block){
	b.Chain = append(b.Chain,*block)
}
func (b *BlockChain)NewGenisBlock()*types.Block{
	return &types.Block{
		Hash:     "0x00000000000000000000000000000",
		Nonce: 100000,
		PrevHash: "",
		Timestamp: int64(time.Now().Unix()),
		Height: 0,
	}
}