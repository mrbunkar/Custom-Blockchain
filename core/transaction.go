package core

import (
	"bytes"
	"crypto/sha256"
	"errors"

	"github.com/mrbunkar/blockchain/crypto"
	"github.com/mrbunkar/blockchain/proto"
	pb "google.golang.org/protobuf/proto"
)

func SignTransaction(tx *proto.Transaction, pk *crypto.Privatekey) error {
	hash := HashTransaction(tx)

	for _, input := range tx.Input {
		if !bytes.Equal(input.PublicKey, pk.PublicKey) {
			return errors.New("private key does not match input's public key")
		}
		sign, err := pk.Sign(hash)
		if err != nil {
			return err
		}
		input.Signature = sign.Bytes()
	}

	return nil
}

func HashTransaction(tx *proto.Transaction) []byte {
	txCopy := pb.Clone(tx).(*proto.Transaction)

	for _, input := range txCopy.Input {
		input.Signature = nil
	}

	b, err := pb.Marshal(txCopy)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)
	return hash[:]
}

func VerifyTransaction(tx *proto.Transaction) bool {

	hash := HashTransaction(tx)

	for _, input := range tx.Input {
		sg := crypto.SignFromBytes(input.Signature)
		if !sg.Verify(input.PublicKey, hash) {
			return false
		}
	}
	return true
}
