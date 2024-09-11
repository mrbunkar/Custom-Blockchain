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

	peerLock       sync.Mutex
	peers          map[proto.NodeClient]*proto.Version
	logger         *zap.SugaredLogger
	bootstrapNodes []string
	proto.UnimplementedNodeServer
}

func NewNode(listenAddr string, bootstrapNodes []string) *Node {
	logger, _ := zap.NewDevelopment()
	return &Node{
		peerLock:       sync.Mutex{},
		peers:          make(map[proto.NodeClient]*proto.Version),
		version:        "Blockchain-0-2",
		Height:         1,
		listenAddr:     listenAddr,
		logger:         logger.Sugar(),
		bootstrapNodes: bootstrapNodes,
	}
}

func (node *Node) isPeer(Addr string) bool {

	ourPeerlist := node.getPeerList()
	for _, peerAddr := range ourPeerlist {
		if peerAddr == Addr {
			return true
		}
	}

	return false
}

func (node *Node) addPeer(peer proto.NodeClient, v *proto.Version) {
	node.peerLock.Lock()
	node.peers[peer] = v
	node.peerLock.Unlock()

	ourVersion := node.getVersion()

	for _, addr := range v.Peerlist {

		if addr != node.listenAddr && !node.isPeer(addr) {
			client, version, err := node.dialRemoteNode(addr)

			if err != nil {
				node.logger.Errorf("[%s]: Error dialing remote node: [%s]", node.listenAddr, addr)
			}

			node.peerLock.Lock()
			node.peers[client] = version
			node.peerLock.Unlock()
			node.logger.Debugf("[%s] Handshake completed with [%s]", ourVersion.ListenAddr, version.ListenAddr)
		}
	}
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

	if len(node.bootstrapNodes) != 0 {
		go node.bootStrapNetwork(node.bootstrapNodes)
	}

	return server.Serve(listen)
}

func (node *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Transaction, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("Recieved peer", peer)

	return nil, nil
}

func (node *Node) Handshake(ctx context.Context, version *proto.Version) (*proto.Version, error) {

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
		Peerlist:   node.getPeerList(),
	}
}

func (node *Node) getPeerList() []string {
	node.peerLock.Lock()
	defer node.peerLock.Unlock()

	peers := []string{}

	for _, version := range node.peers {
		peers = append(peers, version.ListenAddr)
	}

	return peers
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

func (node *Node) bootStrapNetwork(addrs []string) {
	for _, addr := range addrs {
		client, version, err := node.dialRemoteNode(addr)

		if err != nil {
			node.logger.Errorf("[%s]: Error dialing remote node: [%s]", node.listenAddr, addr)
			continue
		}

		node.addPeer(client, version)
		node.logger.Debugf("[%s] Handshake completed with [%s]", node.listenAddr, version.ListenAddr)
	}

	node.logger.Infof("[%s] Bootstrap Netowrk done...", node.listenAddr)
}

func (node *Node) dialRemoteNode(addr string) (proto.NodeClient, *proto.Version, error) {
	client, err := node.NewNodeClient(addr)

	if err != nil {
		node.logger.Errorf("Error in connecting to Client: %s, Error: %s", addr, err)
		return nil, nil, err
	}
	ourVersion := node.getVersion()
	version, err := client.Handshake(context.TODO(), ourVersion)

	if err != nil {
		log.Printf("Error {%s} during hanshake with Client: %s", addr, err)
		return nil, nil, err
	}

	return client, version, nil
}
