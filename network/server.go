package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts

	rpcCh  chan RPC
	quitCh chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC),
		quitCh:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(2 * time.Second)
free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			fmt.Println("Doing stuff every 2 second")
		}
	}

	fmt.Println("Server shutdown")
}

func (s *Server) initTransports() {
	for _, transport := range s.Transports {
		go func(tr Transport) {
			// rpc := tr.Consume()
			// s.rpcCh <- rpc
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(transport)
	}
}
