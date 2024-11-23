package types


type Block struct{
	Height uint64
	Hash string
	//Transactions []Transaction //will replace with data
	Transactions string
	PrevHash string
	Nonce uint64
	Timestamp int64

}

type BlockChain struct{
	Chain []Block

}

type Transaction struct{
	from string
	to string
	amount float64
	txHash string
	

}