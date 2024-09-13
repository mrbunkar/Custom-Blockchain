package main

// Server
// Transport layer TCP, UDP, websocket
// Blockchain

// func main() {
// 	trlocal := network.NewLocalTransport("LOCAL")
// 	trRemote := network.NewLocalTransport("REMOTE")

// 	trlocal.Connect(trRemote)
// 	trRemote.Connect(trlocal)

// 	go func() {
// 		for {
// 			trRemote.SendMessage(trlocal.Addr(), []byte("Hello world"))
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()

// 	opts := network.ServerOpts{
// 		Transports: []network.Transport{trlocal},
// 	}
// 	server := network.NewServer(opts)
// 	server.Start()

// }

import (
	"context"
	"log"
	"time"

	"github.com/mrbunkar/blockchain/crypto"
	"github.com/mrbunkar/blockchain/node"
	"github.com/mrbunkar/blockchain/proto"
	"github.com/mrbunkar/blockchain/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	MakeNode(":3000", []string{})
	MakeNode(":4000", []string{":3000"})
	time.Sleep(1 * time.Second)
	MakeNode(":3100", []string{})
	MakeNode(":3030", []string{":4000", ":i"})
	time.Sleep(3)
	// go func() {
	// 	for {
	// 		time.Sleep(2 * time.Second)
	// 		MakeTransaction()
	// 	}
	// }()

	// MakeTransaction()

	// go node.BootStrapNetwork()

	// log.Fatal(node.Start(":3000"))
	select {}
}

func MakeNode(listenAddr string, bootstrapNodes []string) *node.Node {
	node := node.NewNode(listenAddr, bootstrapNodes)
	go node.Start()
	return node
}

func MakeTransaction() {

	conn, err := grpc.NewClient(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// _, _ = conn.HandleTransaction(context.TODO(), &proto.Transaction{})
	client := proto.NewNodeClient(conn)

	prKey := crypto.GeneratePrivateKey()
	pubKey := prKey.GeneratePublicKey()
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

	// version := &proto.Version{
	// 	Version:    "Blockchain-0-1",
	// 	Height:     1,
	// 	ListenAddr: ":4000",
	// }
	// version, err = client.Handshake(context.TODO(), version)

	_, err = client.HandleTransaction(context.TODO(), tx)

	if err != nil {
		log.Fatal(err)
	}

}
