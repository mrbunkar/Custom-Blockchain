package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignVerify(t *testing.T) {
	pvKey := GeneratePrivateKey()
	pubKey := pvKey.GenerateKeyPublicKey()

	addres := pubKey.Address()
	fmt.Println(addres.ToString())
	msg := []byte("Hello world")

	sign, err := pvKey.Sign(msg)
	assert.Nil(t, err)
	fmt.Println(sign)

	assert.True(t, sign.Verify(pubKey, msg))
}

func TestSignVeriftFail(t *testing.T) {
	pvKey := GeneratePrivateKey()
	pubKey := pvKey.GenerateKeyPublicKey()

	addres := pubKey.Address()
	fmt.Println(addres.ToString())
	msg := []byte("Hello world")

	sign, err := pvKey.Sign(msg)
	assert.Nil(t, err)

	otherPvKey := GeneratePrivateKey()
	otherPuvKey := otherPvKey.GenerateKeyPublicKey()

	assert.True(t, sign.Verify(otherPuvKey, msg))
}
