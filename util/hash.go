package util

import (
	"crypto/rand"
	"io"
	"time"

	"github.com/mrbunkar/blockchain/proto"
)

func RandomHash() []byte {
	buf := make([]byte, 32)
	io.ReadFull(rand.Reader, buf)
	return buf
}

func RandomBLock(height int32) *proto.Block {
	header := &proto.Header{
		Version:       "",
		Height:        height,
		PrevBlockHash: RandomHash(),
		DataHash:      RandomHash(),
		Timestamp:     time.Now().UnixNano(),
		Nonce:         32,
	}

	tx := &proto.Transaction{}

	return &proto.Block{
		Header:      header,
		Transaction: []*proto.Transaction{tx},
	}
}
