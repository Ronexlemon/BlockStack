package types


type Block struct{
	Height uint64
	Hash string
	//Transactions []Transaction
	PrevHash string
	Nonce uint64

}