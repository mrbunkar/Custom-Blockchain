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

func SignBlock(pk *crypto.Privatekey, b *proto.Block) error {
	sign, err := pk.Sign(HashBlock(b))

	if err != nil {
		return err
	}
	b.Header.Signature = sign.Bytes()

	return nil
}

func VerifBlock(block *proto.Block) bool {
	hash := HashBlock(block)

	sg := crypto.SignFromBytes(block.Header.Signature)
	if sg == nil {
		return false
	}

	return sg.Verify(block.Header.PublicKey, hash)
}

func HashHeader(header *proto.Header) []byte {
	headerCopy := pb.Clone(header).(*proto.Header)
	headerCopy.Signature = nil

	b, err := pb.Marshal(header)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)

	return hash[:]
}
