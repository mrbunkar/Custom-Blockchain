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
	"fmt"
	"log"
	"time"

	"github.com/mrbunkar/blockchain/node"
	"github.com/mrbunkar/blockchain/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	node := node.NewNode()
	go func() {
		for {
			time.Sleep(2 * time.Second)
			MakeTransaction()
		}
	}()

	log.Fatal(node.Start(":3000"))
}

func MakeTransaction() {

	conn, err := grpc.NewClient(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// _, _ = conn.HandleTransaction(context.TODO(), &proto.Transaction{})
	client := proto.NewNodeClient(conn)
	version := &proto.Version{
		Version:    "Blockchain-0-1",
		Height:     1,
		ListenAddr: ":4000",
	}
	version, err = client.Handshake(context.TODO(), version)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Client", version)
}
