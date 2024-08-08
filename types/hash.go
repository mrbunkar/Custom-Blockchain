package types

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Hash [32]uint8

func (h Hash) IsZero() bool {
	for i := 0; i < 32; i++ {
		if h[i] != 0 {
			return false
		}
	}
	return true
}

func (h Hash) ToString() string {
	slice := make([]byte, 32)

	for i := 0; i < 32; i++ {
		slice[i] = h[i]
	}

	return hex.EncodeToString(slice)
}

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		msg := fmt.Sprintf("Given bytes with lenght %d shoudl be of lenghth 32\n", len(b))
		panic(msg)
	}

	var value [32]uint8

	for i := 0; i < 32; i++ {
		value[i] = b[i]
	}

	return Hash(value)
}

func RandomBytes(size int) []byte {
	token := make([]byte, 32)
	rand.Read(token)
	return token
}

func RandomHash() Hash {
	token := RandomBytes(32)
	return HashFromBytes(token)
}
