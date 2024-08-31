package node

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/mrbunkar/blockchain/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/peer"
)

type Node struct {
	version string
	Height  int32

	peerLock sync.Mutex
	peers    map[proto.NodeClient]*proto.Version
	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{
		peerLock: sync.Mutex{},
		peers:    make(map[proto.NodeClient]*proto.Version),
		version:  "Blockchain-0-2",
		Height:   1,
	}
}

func (node *Node) addPeer(peer proto.NodeClient, v *proto.Version) {
	node.peerLock.Lock()
	defer node.peerLock.Unlock()
	log.Println("Peer connected", v.ListenAddr)
	node.peers[peer] = v
}

func (node *Node) deletePeer(peer proto.NodeClient) {
	node.peerLock.Lock()
	defer node.peerLock.Unlock()

	delete(node.peers, peer)
}

func (node *Node) Start(listenAddr string) error {
	rpcOpts := []grpc.ServerOption{}
	listen, err := net.Listen("tcp", listenAddr)
	if err != nil {
		panic(err)
	}
	grpc.WithTransportCredentials(insecure.NewCredentials())
	server := grpc.NewServer(rpcOpts...)
	proto.RegisterNodeServer(server, node)
	fmt.Printf("gRPC server listening at %v \n", listen.Addr().String())
	return server.Serve(listen)
}

func (node *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Transaction, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("Recieved peer", peer)

	return nil, nil
}

func (node *Node) Handshake(ctx context.Context, version *proto.Version) (*proto.Version, error) {
	// peer, _ := peer.FromContext(ctx)
	ourVersion := &proto.Version{
		Version: node.version,
		Height:  10,
	}
	client, err := node.NewNodeClient(version.ListenAddr)
	if err != nil {
		log.Printf("Failed to add client. Handshake failed %v\n", version.ListenAddr)
		return nil, err
	}

	node.addPeer(client, version)

	fmt.Println("Handshake complete:", version)
	return ourVersion, nil
}

func (node *Node) getVersion() *proto.Version {
	return &proto.Version{
		Version: node.version,
		Height:  node.Height,
	}
}

func (node *Node) NewNodeClient(listenAddr string) (proto.NodeClient, error) {
	client, err := grpc.NewClient(listenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return proto.NewNodeClient(client), err
}

func (node *Node) BootStrapNetwork(addrs []string) {
	for _, addr := range addrs {
		client, err := node.NewNodeClient(addr)

		if err != nil {
			log.Printf("Error in connecting to Client: %s, Error: %s\n", addr, err)
			continue
		}
		ourVersion := node.getVersion()
		version, err := client.Handshake(context.TODO(), ourVersion)

		if err != nil {
			log.Printf("Error {%s} during hanshake with Client: %s\n", addr, err)
		}
		node.addPeer(client, version)
		fmt.Println("Handshake complete:", version)
	}
}
