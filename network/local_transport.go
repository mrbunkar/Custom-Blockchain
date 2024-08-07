package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr     NetAddr
	consumCh chan RPC

	lock  sync.Mutex
	peers map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:     addr,
		consumCh: make(chan RPC, 1024),
		peers:    make(map[NetAddr]*LocalTransport),
	}
}

func (lt *LocalTransport) Consume() <-chan RPC {
	return lt.consumCh
}

func (lt *LocalTransport) Connect(tr *LocalTransport) error {
	lt.lock.Lock()
	defer lt.lock.Unlock()

	lt.peers[tr.Addr()] = tr

	return nil
}

func (lt *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	lt.lock.Lock()
	defer lt.lock.Unlock()

	peer, ok := lt.peers[to]

	if !ok {
		return fmt.Errorf("%s: cound not send message to %s", lt.addr, to)
	}

	peer.consumCh <- RPC{
		From:    lt.addr,
		Payload: payload,
	}

	return nil
}

func (lt *LocalTransport) Addr() NetAddr {
	return lt.addr
}
