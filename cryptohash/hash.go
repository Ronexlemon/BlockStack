package cryptohash

import (
	"crypto/sha256"
	"encoding/hex"
)


type Hash struct{
	HashValue string
}

func (h *Hash) HashData(data string)string{
	hasher  := sha256.New()
	hasher.Write([]byte(data))
	h.HashValue = hex.EncodeToString(hasher.Sum(nil))
	return h.HashValue
}