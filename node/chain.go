package node

import (
	"encoding/hex"

	"github.com/mrbunkar/blockchain/proto"
)

type Chain struct {
	blockStore BlockStorer
}

func NewChain(bs BlockStorer) *Chain {
	return &Chain{
		blockStore: bs,
	}
}

func (chain *Chain) AddBlock(block *proto.Block) error {

	// validation
	return chain.blockStore.Put(block)
}

func (chain *Chain) GetBlockByHeight(height int32) (*proto.Block, error) {
	return nil, nil
}

func (chain *Chain) GetBlockByHash(hash []byte) (*proto.Block, error) {
	hashHex := hex.EncodeToString(hash)
	return chain.blockStore.Get(hashHex)
}
