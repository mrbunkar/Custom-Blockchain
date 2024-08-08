package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/mrbunkar/blockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {

	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: uint64(time.Now().UnixNano()),
		Height:    1,
		Nonce:     1,
	}

	buf := new(bytes.Buffer)
	h.EncodeBinary(buf)

	assert.Nil(t, h.EncodeBinary(buf))
	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)
}

func TestBlockEncDec(t *testing.T) {
	h := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    1,
			Nonce:     1,
		},
		Transaction: nil,
	}
	buf := new(bytes.Buffer)
	h.EncodeBinary(buf)

	assert.Nil(t, h.EncodeBinary(buf))
	hDecode := &Block{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)

	fmt.Println(hDecode)
}

func TestBlockHash(t *testing.T) {
	h := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    1,
			Nonce:     1,
		},
		Transaction: nil,
	}
	b := h.Hash()

	fmt.Println(b)
	assert.False(t, b.IsZero())
}
