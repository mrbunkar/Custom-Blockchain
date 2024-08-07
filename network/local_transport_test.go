package network

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tr1 := NewLocalTransport("1")
	tr2 := NewLocalTransport("2")

	tr1.Connect(tr2)
	tr2.Connect(tr1)

	assert.Equal(t, tr1.peers[tr2.addr], tr2)
	assert.Equal(t, tr2.peers[tr1.addr], tr1)
}

func TestSendMessage(t *testing.T) {
	tr1 := NewLocalTransport("1")
	tr2 := NewLocalTransport("2")

	tr1.Connect(tr2)
	tr2.Connect(tr1)

	msg := []byte("Hello world")

	assert.Nil(t, tr1.SendMessage(tr2.addr, msg))

	rpc := <-tr2.consumCh
	fmt.Println("Payload", rpc.Payload)
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, tr1.addr)
}
