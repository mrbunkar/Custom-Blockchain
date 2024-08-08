package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/mrbunkar/blockchain/types"
)

type Privatekey struct {
	key *ecdsa.PrivateKey
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

type Signature struct {
	R, S *big.Int
}

func (k Privatekey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.key, data)

	if err != nil {
		return nil, err
	}
	return &Signature{r, s}, nil
}

func GeneratePrivateKey() Privatekey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		panic(err)
	}
	return Privatekey{
		key: key,
	}
}

func (k Privatekey) GenerateKeyPublicKey() PublicKey {
	return PublicKey{
		key: &k.key.PublicKey,
	}
}

func (k PublicKey) Address() types.Address {
	// Conver the key to bytes [key.X,]

	keySlice := elliptic.MarshalCompressed(k.key, k.key.X, k.key.Y)

	hash := sha256.Sum256(keySlice)

	return types.Address(hash[len(hash)-20:])
}

func (sg Signature) Verify(pubK PublicKey, data []byte) bool {
	return ecdsa.Verify(pubK.key, data, sg.R, sg.S)
}
