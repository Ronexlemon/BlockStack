package nonce

import "sync"




type Nonce struct{
	mu sync.Mutex // protect the noonce for concurrency safety
	Nonce uint64
}

func (n *Nonce)Generate()uint64{
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Nonce++
	return n.Nonce	
}
