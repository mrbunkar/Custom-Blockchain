package core

import (
	"crypto/sha256"
	"fmt"

	"github.com/mrbunkar/blockchain/crypto"

	"github.com/mrbunkar/blockchain/proto"
	pb "google.golang.org/protobuf/proto"
)

func HashBlock(block *proto.Block) []byte {
	return HashHeader(block.Header)
}

func SignBlock(pk *crypto.Privatekey, b *proto.Block) error {
	b.Header.PublicKey = pk.PublicKey
	sign, err := pk.Sign(HashBlock(b))

	if err != nil {
		return err
	}
	b.Header.Signature = sign.Bytes()
	return nil
}

func VerifyBlock(block *proto.Block) error {
	hash := HashBlock(block)

	sg := crypto.SignFromBytes(block.Header.Signature)

	if sg == nil {
		return fmt.Errorf("missing Signature")
	}

	ok := sg.Verify(block.Header.PublicKey, hash)

	if !ok {
		return fmt.Errorf("sign Verification failed")
	}
	return nil
}

func HashHeader(header *proto.Header) []byte {
	headerCopy := pb.Clone(header).(*proto.Header)
	headerCopy.Signature = nil

	b, err := pb.Marshal(headerCopy)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)

	return hash[:]
}
