package tree_test

import (
	"fmt"
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_Size(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)

	expected := 12
	actual := tr.Size(tr.Root)
	assert.Equal(t, expected, actual)
}

func TestTree_Max(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)

	expected := 120
	actual := tr.Max(tr.Root)
	assert.Equal(t, expected, actual)
}

func TestTree_Height(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)

	expected := 3
	actual := tr.Height(tr.Root)
	assert.Equal(t, expected, actual)
}

func TestTree_PreOrderTraversal(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)
	tr.PreOrderTraversal(tr.Root)
}

func TestTree_PostOrderTraversal(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)
	tr.PostOrderTraversal(tr.Root)
}

func TestTree_LevelOrderTraversal(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)
	tr.LevelOrderTraversal()
}

func TestTree_Mirror(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)
	fmt.Printf("Before mirror: ")
	tr.Display()
	fmt.Println()
	tr.Mirror(tr.Root)
	fmt.Printf("After mirror: ")
	tr.Display()
	fmt.Println()
}

func TestTree_RemoveLeaves(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)
	fmt.Printf("Before Removing Leaves: ")
	tr.Display()
	fmt.Println()
	tr.RemoveLeaves(tr.Root)
	fmt.Printf("After Removing Leaves: ")
	tr.Display()
	fmt.Println()
}

func TestTree_Linearize(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)
	fmt.Printf("Before Linearize: ")
	tr.Display()
	fmt.Println()
	tr.Linearize(tr.Root)
	fmt.Printf("After Linearize: ")
	tr.Display()
	fmt.Println()
}

func TestTree_Find(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)

	testcases := map[string]map[string]interface{}{
		"tc1": {
			"data": 120,
			"expected": true,
		},

		"tc2": {
			"data": 1200,
			"expected": false,
		},
	}

	for _, tc := range testcases {
		data := tc["data"].(int)
		expected := tc["expected"].(bool)
		actual := tr.Find(data)

		assert.Equal(t, expected, actual)
	}
}

func TestTree_NodeToRoot(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)

	testcases := map[string]map[string]interface{}{
		"tc1": {
			"data": 120,
			"expected": []int{120, 80, 30, 10},
		},

		"tc2": {
			"data": 1200,
			"expected": []int{},
		},

		"tc3": {
			"data": 90,
			"expected": []int{90, 30, 10},
		},
	}

	for _, tc := range testcases {
		data := tc["data"].(int)
		expected := tc["expected"].([]int)
		actual := tr.NodeToRoot(data)

		assert.Equal(t, expected, actual)
	}
}

func TestTree_GetSuccessorAndPredecessor(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)

	testcases := map[string]map[string]interface{}{
		"tc1": {
			"data": 120,
			"successor": 90,
			"predecessor": 110,
		},

		"tc2": {
			"data": 40,
			"successor": 100,
			"predecessor": 90,
		},

		"tc3": {
			"data": 10,
			"successor": 20,
			"predecessor": -1,
		},
	}

	for _, tc := range testcases {
		data := tc["data"].(int)
		expectedSuccessor := tc["successor"].(int)
		expectedPredecessor := tc["predecessor"].(int)

		actualSuccessor, actualPredecessor := tr.GetSuccessorAndPredecessor(data)

		assert.Equal(t, expectedSuccessor, actualSuccessor)
		assert.Equal(t, expectedPredecessor, actualPredecessor)
	}
}


func TestTree_GetCeilAndFloor(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)

	testcases := map[string]map[string]interface{}{
		"tc1": {
			"data": 65,
			"ceil": 70,
			"floor": 60,
		},

		"tc2": {
			"data": 5,
			"ceil": 10,
			"floor": -1,
		},

		"tc3": {
			"data": 120,
			"ceil": -1,
			"floor": 110,
		},
	}

	for _, tc := range testcases {
		data := tc["data"].(int)
		expectedCeil := tc["ceil"].(int)
		expectedFloor := tc["floor"].(int)

		actualCeil, actualFloor := tr.GetCeilAndFloor(data)

		assert.Equal(t, expectedCeil, actualCeil)
		assert.Equal(t, expectedFloor, actualFloor)
	}
}

func TestTree_GetKthLargest(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)

	testcases := map[string]map[string]interface{}{
		"tc1": {
			"k": 1,
			"expected": 120,
		},

		"tc2": {
			"k": 2,
			"expected": 110,
		},

		"tc3": {
			"k": 4,
			"expected": 90,
		},
	}

	for _, tc := range testcases {
		k := tc["k"].(int)
		expected := tc["expected"].(int)
		actual := tr.KthLargest(k)
		assert.Equal(t, expected, actual)
	}
}

func TestTree_GetNodeWithMaxSubtreeSum(t *testing.T) {
	nums := []int{10, 20, -50, -1, -60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, -40, -100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)

	expected := 30
	actual := tr.GetNodeWithMaxSubtreeSum()

	assert.Equal(t, expected, actual)
}

func TestTree_Diameter(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)

	expected := 5
	actual := tr.Diameter()

	assert.Equal(t, expected, actual)
}
