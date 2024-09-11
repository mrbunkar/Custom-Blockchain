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

	for i := 0; i < 100; i++ {
		var (
			block = util.RandomBLock(int32(i))
			err   = chain.AddBlock(block)
		)
		assert.Nil(t, err)

		fetchedBlock, err := chain.GetBlockByHash(core.HashBlock(block))
		assert.Nil(t, err)
		assert.Equal(t, block, fetchedBlock)

		fetchedBlockByHeight, err := chain.GetBlockByHeight(i)
		assert.Nil(t, err)
		assert.Equal(t, fetchedBlockByHeight, block)
	}
}

func TestChainHeight(t *testing.T) {
	bs := NewMemoryBlockStore()
	chain := NewChain(bs)

	for i := 0; i < 10; i++ {
		block := util.RandomBLock(int32(i))
		assert.Nil(t, chain.AddBlock(block))
		assert.Equal(t, chain.Height(), i)
	}
}
