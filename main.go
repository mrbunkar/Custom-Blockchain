package main

import (
	"time"

	"github.com/mrbunkar/blockchain/network"
)

// Server
// Transport layer TCP, UDP, websocket
// Blockchain

func main() {
	trlocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trlocal.Connect(trRemote)
	trRemote.Connect(trlocal)

	go func() {
		for {
			trRemote.SendMessage(trlocal.Addr(), []byte("Hello world"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trlocal},
	}
	server := network.NewServer(opts)
	server.Start()

}

// func main() {
// 	s := "{[]]}"

// 	fmt.Println(test(s))
// 	fmt.Println(test("{[]}()"))
// }
