package tree_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAreTwoTreeMirror(t *testing.T) {
	nums1:= []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tree1 := tree.ConstructGenericTree(nums1)

	nums2 := []int{11, 21, 51, -1, 61, -1, -1, 31, 71, -1, 81, 111, -1, 121, -1, -1, 91, -1, -1, 41, 101, -1, -1, -1}
	tree2 := tree.ConstructGenericTree(nums2)

	nums3 := []int{10, 20, 50, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, 130, -1, -1}
	tree3 := tree.ConstructGenericTree(nums3)

	isSimilar := tree.AreTwoTreeMirror(tree1, tree2)
	assert.False(t, isSimilar)

	isSimilar = tree.AreTwoTreeMirror(tree1, tree3)
	assert.True(t, isSimilar)
}
