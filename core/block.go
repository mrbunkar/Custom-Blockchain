package core

import (
	"crypto/sha256"

	"github.com/mrbunkar/blockchain/crypto"

	"github.com/mrbunkar/blockchain/proto"
	pb "google.golang.org/protobuf/proto"
)

func HashBlock(block *proto.Block) []byte {
	b, err := pb.Marshal(block)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)

	return hash[:]
}

func SignBlock(pk *crypto.Privatekey, b *proto.Block) *crypto.Signature {
	sign, err := pk.Sign(HashBlock(b))

	if err != nil {
		panic(err)
	}

	return sign
}
