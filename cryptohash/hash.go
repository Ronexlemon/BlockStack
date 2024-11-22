package cryptohash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/RonexLemon/chain/types"
)


type Hash struct{
	HashValue string
}

func (h *Hash) HashData(data *types.Block)string{
	hasher  := sha256.New()
	serializedData := fmt.Sprintf("%d%s%s%d",
		data.Height,
		data.PrevHash,
		data.Hash,
		data.Nonce,
	)
	hasher.Write([]byte(serializedData))
	h.HashValue = hex.EncodeToString(hasher.Sum(nil))
	return h.HashValue
}