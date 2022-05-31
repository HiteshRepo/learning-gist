package amazon_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/amazon"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectCycleInDirectedGraph(t *testing.T) {
	graph := [][]int{
		{1},
		{2},
		{1,3,5},
		{},
		{5},
		{},
	}

	isCyclic := amazon.DetectCycleInDirectedGraph(graph)
	assert.True(t, isCyclic)

	graph = [][]int{
		{1},
		{2},
		{3,5},
		{},
		{5},
		{},
	}

	isCyclic = amazon.DetectCycleInDirectedGraph(graph)
	assert.False(t, isCyclic)
}
