package node

import (
	"testing"

	"github.com/mrbunkar/blockchain/core"
	"github.com/mrbunkar/blockchain/util"
	"github.com/stretchr/testify/assert"
)

func TestAddBlock(t *testing.T) {
	bs := NewMemoryBlockStore()
	chain := NewChain(bs)

	block := util.RandomBLock(16)
	err := chain.AddBlock(block)

	assert.Nil(t, err)

	fetchedBlock, err := chain.GetBlockByHash(core.HashBlock(block))
	assert.Nil(t, err)
	assert.Equal(t, block, fetchedBlock)
}
