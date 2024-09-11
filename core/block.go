package core

import (
	"crypto/sha256"

	"github.com/mrbunkar/blockchain/crypto"

	"github.com/mrbunkar/blockchain/proto"
	pb "google.golang.org/protobuf/proto"
)

func HashBlock(block *proto.Block) []byte {
	return HashHeader(block.Header)
}

func SignBlock(pk *crypto.Privatekey, b *proto.Block) *crypto.Signature {
	sign, err := pk.Sign(HashBlock(b))

	if err != nil {
		panic(err)
	}

	return sign
}

func HashHeader(header *proto.Header) []byte {
	b, err := pb.Marshal(header)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)

	return hash[:]
}
