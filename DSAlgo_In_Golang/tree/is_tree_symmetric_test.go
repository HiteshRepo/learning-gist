package tree_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsTreeSymmetric(t *testing.T) {
	nums:= []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)

	isSymmetric := tree.IsTreeSymmetric(tr)
	assert.False(t, isSymmetric)

	nums = []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, 110, -1, -1}
	tr = tree.ConstructGenericTree(nums)

	isSymmetric = tree.IsTreeSymmetric(tr)
	assert.True(t, isSymmetric)
}
