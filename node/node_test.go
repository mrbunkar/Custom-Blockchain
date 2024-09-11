package node

import (
	"testing"
	"time"
)

func MakeNode(listenAddr string, bootstrapNodes []string) *Node {
	node := NewNode(listenAddr, bootstrapNodes)
	go node.Start()
	return node
}

func TestNode(t *testing.T) {
	MakeNode(":3000", []string{})
	MakeNode(":4000", []string{":3000"})
	time.Sleep(1 * time.Second)
	MakeNode(":3030", []string{":4000", ":3100"})
	time.Sleep(3 * time.Second)
}
