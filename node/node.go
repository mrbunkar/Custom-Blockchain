package node

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/mrbunkar/blockchain/core"
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
	pool           *Mempool
	isValidator    bool
	Chain          *Chain
	proto.UnimplementedNodeServer
}

func NewNode(listenAddr string, bootstrapNodes []string, isValidator bool) *Node {
	logger, _ := zap.NewDevelopment()
	bs := NewMemoryBlockStore()

	return &Node{
		peerLock:       sync.Mutex{},
		peers:          make(map[proto.NodeClient]*proto.Version),
		version:        "Blockchain-0-2",
		Height:         0,
		listenAddr:     listenAddr,
		logger:         logger.Sugar(),
		bootstrapNodes: bootstrapNodes,
		pool:           NewMempool(),
		isValidator:    isValidator,
		Chain:          NewChain(bs),
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

func (node *Node) GenesisBlock() *proto.Block {

	header := &proto.Header{
		Version:       node.version,
		Height:        node.Height,
		PrevBlockHash: []byte{},
		DataHash:      []byte{},
		Timestamp:     time.Now().Unix(),
		Nonce:         0,
	}

	tx := &proto.Transaction{}

	return &proto.Block{
		Header:      header,
		Transaction: []*proto.Transaction{tx},
	}
}

func (node *Node) TxsVerification(txs *proto.Transaction) bool {
	// @TODO Add UTXO model for verification

	for _, input := range txs.Input {
		if len(input.Signature) == 0 {
			return false
		}
	}

	return core.VerifyTransaction(txs)
}

func (node *Node) VerifyBlock(block *proto.Block) error {

	if !core.VerifBlock(block) {
		// Cryptography validation
		return fmt.Errorf("Header Validation failed")
	}

	if err := node.Chain.VerifyBlock(block); err != nil {
		return err
	}

	// Txs Verification
	txs := block.Transaction
	for _, tx := range txs {
		if !node.TxsVerification(tx) {
			return fmt.Errorf("Block: Transaction verification failed")
		}
	}

	for _, tx := range txs {
		if !node.pool.Check(tx) {
			node.pool.DeleteTx(tx)
		}
	}

	return nil
}

func (node *Node) AddNewBlock() *proto.Block {
	// @TODO: Proof Of Stake Algorithi for Add New Block
	return nil
}

func (node *Node) validatorLoop() {
	ticker := time.NewTicker(1 * time.Second)

	for {
		<-ticker.C
		if node.pool.Size() > 5 {
			fmt.Println("Time to mine")
		}
	}
}

func (node *Node) AddGenesisBlock() error {

	genesisBlock := node.GenesisBlock()

	if err := node.Chain.AddBlock(genesisBlock); err != nil {
		node.logger.Panic(err)
		return nil
	}
	node.logger.Debugf("Genesis block added to Chain. We: [%s]", node.listenAddr)

	go func() {
		if err := node.Broadcast(genesisBlock); err != nil {
			node.logger.Error("Broadcast Failed, Error: %s\n", err)
			return
		}
	}()
	return nil
}

func (node *Node) HandleBlock(ctx context.Context, block *proto.Block) (*proto.Block, error) {
	peer, _ := peer.FromContext(ctx)

	if block.Header.Height < node.Height {
		return nil, nil
	}

	if block.Header.Height > node.Height {
		// @TODO: ask for full chain as we have the outdated chain
		return nil, nil
	}

	if err := node.VerifyBlock(block); err != nil {
		return nil, err
	}

	if err := node.Chain.AddBlock(block); err != nil {
		node.logger.Errorf("Add block failed, err: {%s}\n", err)
		return nil, err
	}

	node.Height += 1

	go func() {
		if err := node.Broadcast(block); err != nil {
			// @TODO: check error if we have outdated chain, then we nee to sync
			//  node.SyncWithPeer(*proto.Client)
			node.logger.Error("Broadcast Failed, Error: %s\n", err)
			return
		}
	}()

	node.logger.Debugf("Recieved new block from [%s]. We: [%s]", peer.Addr, node.listenAddr)
	node.logger.Debugf("New Block added to chain. Chain lenght [%d]. We: [%s]\n", node.Height, node.listenAddr)
	return nil, nil
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
		node.bootStrapNetwork(node.bootstrapNodes)
	}

	if node.isValidator {
		go node.validatorLoop()
	}

	if node.isValidator {
		if err := node.AddGenesisBlock(); err != nil {
			node.logger.Errorln(err)
			// What if its already a working node, that means you need to ask for copy when booted
		}
		// What if its already a working node, that means you need to ask for copy when booted
	}

	return server.Serve(listen)
}

func (node *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Transaction, error) {
	peer, _ := peer.FromContext(ctx)

	if !node.pool.Check(tx) {

		if !node.TxsVerification(tx) {
			return nil, fmt.Errorf("Transaction failed Verification process")
		}

		node.pool.StoreTx(tx)
		node.logger.Debugf("Recieved transaction from [%s]. We [%s]", peer.Addr, node.listenAddr)
		go func() {
			if err := node.Broadcast(tx); err != nil {
				node.logger.Errorln("Error broadcating the transaction", err)
				return
			}
		}()
	}
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

func (node *Node) Broadcast(msg any) error {

	// Broadcast a newly transaction to all the nodes on the network

	for peer, version := range node.peers {
		switch v := msg.(type) {
		case *proto.Transaction:
			_, err := peer.HandleTransaction(context.TODO(), v)

			if err != nil {
				node.logger.Errorf("Error broadcasting transaction to peer [%s]. We: [%s\n", version.ListenAddr, node.listenAddr)
				return err
				// continue
			}
		case *proto.Block:
			_, err := peer.HandleBlock(context.TODO(), v)

			if err != nil {
				node.logger.Errorf("Error broadcasting block to peer [%s]\n. We: [%s]", version.ListenAddr, node.listenAddr)
				// @TODO, check what possible scenario for the block failure
				// What is peer node have bigger chain
				return err
			}
		}

	}

	return nil
}
