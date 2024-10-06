package core

import (
	"fmt"
	"testing"

	"github.com/mrbunkar/blockchain/crypto"
	"github.com/mrbunkar/blockchain/proto"
	"github.com/mrbunkar/blockchain/util"
	"github.com/stretchr/testify/assert"
)

// balance = 100
//
//	You want to send 5 to someone, then you need to create
//
// 2 oupute one is of 5 and other is of 95 to you iteself
func TestTransaction(t *testing.T) {
	fromPk := crypto.GeneratePrivateKey()
	toPk := crypto.GeneratePrivateKey()

	fromPbK := fromPk.GeneratePublicKey()
	toPbK := toPk.GeneratePublicKey()

	input := &proto.Input{
		PrevOutHash: util.RandomHash(),
		PrevOutIdx:  0,
		PublicKey:   fromPbK,
	}

	output1 := &proto.Output{
		Amount:   1,
		Reciever: toPbK,
	}

	output2 := &proto.Output{
		Amount:   2,
		Reciever: fromPbK,
	}

	transaction := &proto.Transaction{
		Version: 1,
		Input:   []*proto.Input{input},
		Output:  []*proto.Output{output1, output2},
	}

	SignTransaction(transaction, fromPk)

	fmt.Println(VerifyTransaction(transaction))

	// Test is failing because of verify.
	assert.True(t, VerifyTransaction(transaction))
}
