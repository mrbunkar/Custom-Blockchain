package core

import (
	"crypto/sha256"

	"github.com/mrbunkar/blockchain/crypto"
	"github.com/mrbunkar/blockchain/proto"
	pb "google.golang.org/protobuf/proto"
)

func SignTransaction(tx *proto.Transaction, pk *crypto.Privatekey) *crypto.Signature {
	hash := HashTransaction(tx)

	sign, err := pk.Sign(hash)

	if err != nil {
		panic(err)
	}

	return sign
}

func HashTransaction(tx *proto.Transaction) []byte {
	b, err := pb.Marshal(tx)

	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)
	return hash[:]
}

func VerifyTransaction(tx *proto.Transaction) bool {
	for _, input := range tx.Input {
		sg := crypto.SignFromBytes(input.Signature)
		// input.Signature = nil
		hash := HashTransaction(tx)
		if !sg.Verify(input.PublicKey, hash) {
			return false
		}
	}
	return true
}
