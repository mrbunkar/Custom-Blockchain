package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/mrbunkar/blockchain/types"
)

func RandBLock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     uint64(time.Now().UnixNano()),
	}

	tx := Transaction{
		Data: []byte("Foo"),
	}

	return &Block{
		Header:      header,
		Transaction: []Transaction{tx},
	}
}

func TestHashBlock(t *testing.T) {
	b := RandBLock(32)
	fmt.Println(b.Hash(BlockHasher{}))
}
