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
)

func main() {
	MakeNode(":3000", []string{})
	MakeNode(":4000", []string{":3000"})
	// go func() {
	// 	for {
	// 		time.Sleep(2 * time.Second)
	// 		MakeTransaction()
	// 	}
	// }()

	// go node.BootStrapNetwork()

	// log.Fatal(node.Start(":3000"))
	select {}
}

func MakeNode(listenAddr string, bootstrapNodes []string) *node.Node {
	node := node.NewNode(listenAddr)
	go node.Start()

	time.Sleep(1 * time.Second)
	if len(bootstrapNodes) != 0 {
		if err := node.BootStrapNetwork(bootstrapNodes); err != nil {
			log.Printf("Error bootstrapping network: %v\n", err)
		}
	}

	return node
}

func MakeTransaction() {

	// conn, err := grpc.NewClient(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// // _, _ = conn.HandleTransaction(context.TODO(), &proto.Transaction{})
	// client := proto.NewNodeClient(conn)

	node := node.NewNode(":4000")
	fmt.Println(1)
	client, err := node.NewNodeClient(":5000")
	fmt.Println(2)

	if err != nil {
		panic(err)
	}

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
