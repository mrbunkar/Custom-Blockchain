package core

import (
	"github.com/mrbunkar/blockchain/crypto"
)

// Transaction needs to be signed
type Transaction struct {
	Data []byte

	PublicKey crypto.PublicKey
	Signature *crypto.Signature
}

func (tx *Transaction) Sign(prvKey crypto.Privatekey) error {
	s, err := prvKey.Sign(tx.Data)
	if err != nil {
		return err
	}
	tx.PublicKey = prvKey.GenerateKeyPublicKey()
	tx.Signature = s

	return nil
}
