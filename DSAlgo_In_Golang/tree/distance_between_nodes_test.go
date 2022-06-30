package tree_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDistanceBetweenNodes(t *testing.T) {
	nums := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}
	tr := tree.ConstructGenericTree(nums)

	testcases := map[string]map[string]interface{}{
		"tc1": {
			"val1": 100,
			"val2": 90,
			"expected": 4,
		},

		"tc2": {
			"val1": 110,
			"val2": 120,
			"expected": 2,
		},
	}

	for _, tc := range testcases {
		val1 := tc["val1"].(int)
		val2 := tc["val2"].(int)
		expected := tc["expected"].(int)
		actual := tree.GetDistanceBetweenNodes(val1, val2, tr)

		assert.Equal(t, expected, actual)
	}
}
