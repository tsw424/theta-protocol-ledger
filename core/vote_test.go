package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thetatoken/ukulele/common"
	"github.com/thetatoken/ukulele/rlp"
)

func TestEncoding(t *testing.T) {
	assert := assert.New(t)

	votes := NewVoteSet()
	votes.AddVote(Vote{
		Block: CreateTestBlock("", "").Hash(),
		ID:    common.HexToAddress("A1"),
		Epoch: 1,
	})
	votes.AddVote(Vote{
		Block: CreateTestBlock("", "").Hash(),
		ID:    common.HexToAddress("A2"),
		Epoch: 1,
	})

	votes2 := NewVoteSet()
	b, err := rlp.EncodeToBytes(votes)
	assert.Nil(err)
	err = rlp.DecodeBytes(b, &votes2)
	assert.Nil(err)

	vs := votes2.Votes()
	vs0 := votes.Votes()

	assert.Equal(2, len(vs))
	assert.Equal(common.HexToAddress("A1"), vs[0].ID)
	assert.NotNil(vs[0].Block)
	assert.Equal(vs0[0].Block, vs[0].Block)

	assert.Equal(common.HexToAddress("A2"), vs[1].ID)
	assert.NotNil(vs[1].Block)
	assert.Equal(vs0[1].Block, vs[1].Block)
}

func TestDedup(t *testing.T) {
	assert := assert.New(t)

	votes1 := NewVoteSet()
	votes1.AddVote(Vote{
		Block: CreateTestBlock("B1", "").Hash(),
		ID:    common.HexToAddress("A1"),
		Epoch: 1,
	})
	// Duplicate votes
	votes1.AddVote(Vote{
		Block: CreateTestBlock("B1", "").Hash(),
		ID:    common.HexToAddress("A1"),
		Epoch: 1,
	})
	votes1.AddVote(Vote{
		Block: CreateTestBlock("B2", "").Hash(),
		ID:    common.HexToAddress("A2"),
		Epoch: 1,
	})
	assert.Equal(2, len(votes1.Votes()))

	votes2 := NewVoteSet()
	// Duplicate vote.
	votes2.AddVote(Vote{
		Block: CreateTestBlock("B1", "").Hash(),
		ID:    common.HexToAddress("A1"),
		Epoch: 1,
	})
	// Duplicate vote from newer epoch.
	votes2.AddVote(Vote{
		Block: CreateTestBlock("B1", "").Hash(),
		ID:    common.HexToAddress("A1"),
		Epoch: 3,
	})
	votes2.AddVote(Vote{
		Block: CreateTestBlock("B2", "").Hash(),
		ID:    common.HexToAddress("A3"),
		Epoch: 1,
	})

	votes := votes1.Merge(votes2)
	assert.Equal(4, len(votes.Votes()))

	votes = votes.KeepLatest()
	assert.Equal(3, len(votes.Votes()))
}
