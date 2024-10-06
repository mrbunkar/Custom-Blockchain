package main

import (
	"context"
	"log"
	"time"

	"github.com/mrbunkar/blockchain/core"
	"github.com/mrbunkar/blockchain/crypto"
	"github.com/mrbunkar/blockchain/node"
	"github.com/mrbunkar/blockchain/proto"
	"github.com/mrbunkar/blockchain/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	MakeNode(":3000", []string{}, "Normal")
	MakeNode(":4000", []string{":3000"}, "Normal")
	time.Sleep(1 * time.Second)
	MakeNode(":3100", []string{}, "Miner")
	MakeNode(":3030", []string{":4000", ":3100"}, "Normal")
	time.Sleep(1 * time.Second)

	for {
		time.Sleep(1 * time.Second)
		MakeTransaction()
	}

}

func MakeNode(listenAddr string, bootstrapNodes []string, nodeType string) *node.Node {
	node := node.NewNode(listenAddr, bootstrapNodes, nodeType)
	go node.Start()
	return node
}

func MakeTransaction() {

	conn, err := grpc.NewClient(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// _, _ = conn.HandleTransaction(context.TODO(), &proto.Transaction{})
	client := proto.NewNodeClient(conn)

	prKey := crypto.GeneratePrivateKey()
	pubKey := prKey.PublicKey
	tx := &proto.Transaction{
		Version: 1,
		Input: []*proto.Input{
			{
				PrevOutHash: util.RandomHash(),
				PrevOutIdx:  0,
				PublicKey:   pubKey,
			},
		},
		Output: []*proto.Output{
			{
				Amount:   1,
				Reciever: pubKey,
			},
		},
	}

	core.SignTransaction(tx, prKey)

	_, err = client.HandleTransaction(context.TODO(), tx)

	if err != nil {
		log.Fatal(err)
	}

}
