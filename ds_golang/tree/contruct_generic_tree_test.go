package tree_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/tree"
	"testing"
)

func TestConstructGenericTree(t *testing.T) {
	givenArr := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}

	tr := tree.ConstructGenericTree(givenArr)
	tr.Display()
}
