package core

import (
	"testing"

	"github.com/mrbunkar/blockchain/crypto"
	"github.com/stretchr/testify/assert"
)

func RandomTransction() *Transaction {
	return &Transaction{
		Data: []byte("Foo"),
	}
}

func TestTransaction(t *testing.T) {
	tx := RandomTransction()

	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Signature)
	assert.Nil(t, tx.PublicKey)
}
