package node

import (
	"encoding/hex"
	"sync"

	"github.com/mrbunkar/blockchain/core"
	"github.com/mrbunkar/blockchain/proto"
)

type Mempool struct {
	mu  sync.Mutex
	txs map[string]*proto.Transaction
}

func NewMempool() *Mempool {
	return &Mempool{
		txs: make(map[string]*proto.Transaction),
	}
}

func (pool *Mempool) Check(tx *proto.Transaction) bool {
	hashHex := hex.EncodeToString(core.HashTransaction(tx))

	_, ok := pool.txs[hashHex]
	return ok
}

func (pool *Mempool) StoreTx(tx *proto.Transaction) {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	hashHex := hex.EncodeToString(core.HashTransaction(tx))
	pool.txs[hashHex] = tx
}

func (pool *Mempool) DeleteTx(tx *proto.Transaction) {
	hashHex := hex.EncodeToString(core.HashTransaction(tx))
	delete(pool.txs, hashHex)
}

func (pool *Mempool) Size() int {
	return len(pool.txs)
}
