package core

import (
	"github.com/mrbunkar/blockchain/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct {
}

// func (BlockHasher) Hash(b *Block) types.Hash {
// 	buf := new(bytes.Buffer)
// 	enocder := gob.NewEncoder(buf)

// 	if err := enocder.Encode(b.Header); err != nil {
// 		panic(err)
// 	}

// 	h := sha256.Sum256(buf.Bytes())
// 	return types.Hash(h)
// }
