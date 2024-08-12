package core

import (
	"io"

	"github.com/mrbunkar/blockchain/crypto"
	"github.com/mrbunkar/blockchain/types"
)

type Header struct {
	Version       uint32
	DataHash      types.Hash
	PrevBlockHash types.Hash
	Timestamp     uint64
	Height        uint32
	Nonce         uint32
}

type Block struct {
	*Header
	Transaction []Transaction
	Validator   crypto.PublicKey
	Signture    *crypto.Signature

	//Cached version of the header hash
	hash types.Hash
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}
	return b.hash
}

func (b *Block) Decoder(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encoder(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func NewBlock(h *Header, t []Transaction) *Block {
	return &Block{
		Header:      h,
		Transaction: t,
	}
}
