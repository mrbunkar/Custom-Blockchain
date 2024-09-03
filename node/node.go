package node

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/mrbunkar/blockchain/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/peer"
)

type Node struct {
	version    string
	Height     int32
	listenAddr string

	peerLock sync.Mutex
	peers    map[proto.NodeClient]*proto.Version
	logger   *zap.SugaredLogger
	proto.UnimplementedNodeServer
}

func NewNode(listenAddr string) *Node {
	logger, _ := zap.NewDevelopment()
	return &Node{
		peerLock:   sync.Mutex{},
		peers:      make(map[proto.NodeClient]*proto.Version),
		version:    "Blockchain-0-2",
		Height:     1,
		listenAddr: listenAddr,
		logger:     logger.Sugar(),
	}
}

func (node *Node) addPeer(peer proto.NodeClient, v *proto.Version) {
	node.peerLock.Lock()
	defer node.peerLock.Unlock()
	// log.Println("Peer connected", v.ListenAddr)
	node.peers[peer] = v
}

func (node *Node) deletePeer(peer proto.NodeClient) {
	node.peerLock.Lock()
	defer node.peerLock.Unlock()

	delete(node.peers, peer)
}

func (node *Node) Start() error {

	var (
		rpcOpts = []grpc.ServerOption{}
		server  = grpc.NewServer(rpcOpts...)
	)

	listen, err := net.Listen("tcp", node.listenAddr)
	if err != nil {
		panic(err)
	}

	grpc.WithTransportCredentials(insecure.NewCredentials())

	proto.RegisterNodeServer(server, node)

	node.logger.Infof("gRPC server listening at %v", listen.Addr().String())

	return server.Serve(listen)
}

func (node *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Transaction, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("Recieved peer", peer)

	return nil, nil
}

func (node *Node) Handshake(ctx context.Context, version *proto.Version) (*proto.Version, error) {
	// peer, _ := peer.FromContext(ctx)
	// ourVersion := &proto.Version{
	// 	Version:    node.version,
	// 	Height:     10,
	// 	ListenAddr: node.listenAddr,
	// }
	ourVersion := node.getVersion()
	client, err := node.NewNodeClient(version.ListenAddr)
	if err != nil {
		node.logger.Errorf("Failed to add client. Handshake failed %v", version.ListenAddr)
		return nil, err
	}

	node.addPeer(client, version)

	node.logger.Debugf("[%s] Handshake completed with [%s]", ourVersion.ListenAddr, version.ListenAddr)
	return ourVersion, nil
}

func (node *Node) getVersion() *proto.Version {
	return &proto.Version{
		Version:    node.version,
		Height:     node.Height,
		ListenAddr: node.listenAddr,
	}
}

func (node *Node) NewNodeClient(listenAddr string) (proto.NodeClient, error) {

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time)
	// defer cancel()
	client, err := grpc.NewClient(listenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}
	client.Connect()

	return proto.NewNodeClient(client), err
}

func (node *Node) BootStrapNetwork(addrs []string) error {
	for _, addr := range addrs {
		client, err := node.NewNodeClient(addr)

		if err != nil {
			node.logger.Errorf("Error in connecting to Client: %s, Error: %s", addr, err)
			continue
		}
		ourVersion := node.getVersion()
		version, err := client.Handshake(context.TODO(), ourVersion)

		if err != nil {
			log.Printf("Error {%s} during hanshake with Client: %s", addr, err)
			continue
		}
		node.addPeer(client, version)
		node.logger.Debugf("[%s] Handshake completed with [%s]", ourVersion.ListenAddr, version.ListenAddr)
	}
	return nil
}
