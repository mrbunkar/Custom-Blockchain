package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"

	"github.com/mrbunkar/blockchain/types"
)

type Header struct {
	Version   uint32
	PrevBlock types.Hash
	Timestamp uint64
	Height    uint32
	Nonce     uint32
}

func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
		return nil
	}
	if err := binary.Write(w, binary.LittleEndian, &h.PrevBlock); err != nil {
		return nil
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Timestamp); err != nil {
		return nil
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Height); err != nil {
		return nil
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Nonce); err != nil {
		return nil
	}

	return nil
}

func (h *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return nil
	}
	if err := binary.Read(r, binary.LittleEndian, &h.PrevBlock); err != nil {
		return nil
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Timestamp); err != nil {
		return nil
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
		return nil
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Nonce); err != nil {
		return nil
	}

	return nil
}

type Block struct {
	Header
	Transaction []Transaction

	//Cached version of the header hash
	hash types.Hash
}

func (b *Block) EncodeBinary(w io.Writer) error {
	if err := b.Header.EncodeBinary(w); err != nil {
		return err
	}
	for _, tx := range b.Transaction {
		if err := tx.EncodeBinary(w); err != nil {
			return err
		}
	}
	return nil
}

func (b *Block) DecodeBinary(r io.Reader) error {
	if err := b.Header.DecodeBinary(r); err != nil {
		return err
	}

	for _, tx := range b.Transaction {
		if err := tx.DecodeBinary(r); err != nil {
			return err
		}
	}

	return nil
}

func (b *Block) Hash() types.Hash {
	buf := new(bytes.Buffer)
	b.Header.EncodeBinary(buf)

	if b.hash.IsZero() {
		b.hash = types.Hash(sha256.Sum256(buf.Bytes()))
	}

	return b.hash
}
