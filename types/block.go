package types


type Block struct{
	Height uint64
	Hash string
	//Transactions []Transaction
	PrevHash string
	Nonce uint64
	Timestamp int64

}

type BlockChain struct{
	Chain []Block

}