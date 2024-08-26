package core

func RandomTransction() *Transaction {
	return &Transaction{
		Data: []byte("Foo"),
	}
}

// func TestTransaction(t *testing.T) {
// 	tx := RandomTransction()

// 	privKey := crypto.GeneratePrivateKey()
// 	assert.Nil(t, tx.Sign(privKey))
// 	assert.Nil(t, tx.Signature)
// 	assert.Nil(t, tx.PublicKey)
// }
