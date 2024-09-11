package node

import (
	"encoding/hex"
	"fmt"

	"github.com/mrbunkar/blockchain/core"
	"github.com/mrbunkar/blockchain/proto"
)

type HeaderList struct {
	headers []*proto.Header
}

func (hd *HeaderList) AddHeader(header *proto.Header) {
	hd.headers = append(hd.headers, header)
}

func (hd *HeaderList) Len() int {
	return len(hd.headers)
}

func (hd *HeaderList) Get(index int) *proto.Header {
	return hd.headers[index]
}

func (hd *HeaderList) Height() int {
	return len(hd.headers) - 1
}

func NewHeaderList() *HeaderList {
	return &HeaderList{
		headers: []*proto.Header{},
	}
}

type Chain struct {
	blockStore BlockStorer
	headerList *HeaderList
}

func NewChain(bs BlockStorer) *Chain {
	return &Chain{
		blockStore: bs,
		headerList: NewHeaderList(),
	}
}

func (chain *Chain) AddBlock(block *proto.Block) error {

	// Add the block header to headerlist
	chain.headerList.AddHeader(block.Header)
	// validation
	return chain.blockStore.Put(block)
}

func (chain *Chain) GetBlockByHeight(height int) (*proto.Block, error) {
	if chain.Height() < height {
		return nil, fmt.Errorf("Given height [%d] is more than Chain height [%d]",
			height, chain.Height())
	}

	header := chain.headerList.Get(height)
	hash := core.HashHeader(header)

	return chain.GetBlockByHash(hash)

}

func (chain *Chain) GetBlockByHash(hash []byte) (*proto.Block, error) {
	hashHex := hex.EncodeToString(hash)
	return chain.blockStore.Get(hashHex)
}

func (chain *Chain) Height() int {
	return chain.headerList.Height()
}