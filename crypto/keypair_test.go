package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignVerify(t *testing.T) {
	pvKey := GeneratePrivateKey()
	pubKey := pvKey.GeneratePublicKey()

	addres := pubKey.Address()
	fmt.Println(addres.ToString())
	msg := []byte("Hello world")

	sign, err := pvKey.Sign(msg)
	assert.Nil(t, err)
	fmt.Println(sign)

	assert.True(t, sign.Verify(pubKey, msg))
}

func TestSignVerifyFail(t *testing.T) {
	pvKey := GeneratePrivateKey()
	pubKey := pvKey.GeneratePublicKey()

	addres := pubKey.Address()
	fmt.Println(addres.ToString())
	msg := []byte("Hello world")

	sign, err := pvKey.Sign(msg)
	assert.Nil(t, err)

	otherPvKey := GeneratePrivateKey()
	otherPubKey := otherPvKey.GeneratePublicKey()
	fmt.Println(sign)
	fmt.Println(SignFromBytes(sign.Bytes()))
	assert.Equal(t, sign, SignFromBytes(sign.Bytes()))
	assert.False(t, sign.Verify(otherPubKey, msg))
	assert.False(t, sign.Verify(pubKey, []byte("Hello")))
}
