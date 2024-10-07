package core

import (
	"crypto/rand"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/mrbunkar/blockchain/crypto"

	"github.com/mrbunkar/blockchain/proto"
	"github.com/stretchr/testify/assert"
)

func RandomHash() []byte {
	buf := make([]byte, 32)
	io.ReadFull(rand.Reader, buf)
	return buf
}

func RandomBLock(height int32) *proto.Block {
	header := &proto.Header{
		Version:       "1",
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

func TestHashBlock(t *testing.T) {
	b := RandomBLock(32)
	hash := HashBlock(b)
	fmt.Println(hash)
	assert.Equal(t, 32, len(hash))
}

func TestSign(t *testing.T) {

	b := RandomBLock(32)
	pk := crypto.GeneratePrivateKey()
	// pb := pk.PublicKey

	// b.Header.PublicKey = pb
	fmt.Println("Sign", b.Header.Signature)
	fmt.Println(HashBlock(b))
	SignBlock(pk, b)
	fmt.Println(HashBlock(b))
	ok := VerifyBlock(b)
	fmt.Println(ok)
	if err := VerifyBlock(b); err != nil {
		fmt.Println(err)
	}

	// assert.False(t, VerifBlock(RandomBLock(12)))
}
