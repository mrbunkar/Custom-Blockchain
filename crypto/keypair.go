package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"fmt"
	"math/big"

	"github.com/mrbunkar/blockchain/types"
)

type Privatekey struct {
	key       *ecdsa.PrivateKey
	PublicKey PublicKey
}

type PublicKey []byte

type Signature struct {
	R, S *big.Int
}

func (k *Privatekey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.key, data)

	if err != nil {
		return nil, err
	}
	return &Signature{r, s}, nil
}

func GeneratePrivateKey() *Privatekey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		panic(err)
	}
	privateKey := &Privatekey{
		key: key,
	}

	privateKey.PublicKey = privateKey.GeneratePublicKey()
	return privateKey
}

func (k *Privatekey) GeneratePublicKey() PublicKey {
	return elliptic.MarshalCompressed(k.key.PublicKey, k.key.PublicKey.X, k.key.PublicKey.Y)
}

func (k PublicKey) Address() types.Address {
	// Convert the key to bytes [key.X,]
	hash := sha256.Sum256(k)

	return types.AddressFromBytes(hash[len(hash)-20:])
}

func (sg *Signature) Verify(pubK PublicKey, data []byte) bool {
	fmt.Println("Verify", pubK)
	if pubK == nil {
		return false
	}

	x, y := elliptic.UnmarshalCompressed(elliptic.P256(), pubK)

	key := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	return ecdsa.Verify(key, data, sg.R, sg.S)
}

func (sg *Signature) Bytes() []byte {
	b, err := asn1.Marshal(*sg)

	if err != nil {
		panic(err)
	}
	return b
}

func SignFromBytes(b []byte) *Signature {
	var sig Signature
	_, err := asn1.Unmarshal(b, &sig)
	if err != nil {
		return nil
	}
	return &sig
}
