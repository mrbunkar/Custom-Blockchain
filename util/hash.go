package util

import (
	"crypto/rand"
	"io"
)

func RandomHash() []byte {
	buf := make([]byte, 32)
	io.ReadFull(rand.Reader, buf)
	return buf
}
