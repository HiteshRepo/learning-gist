package maximizeXOR_test

import (
	"github.com/hiteshrepo/learninggist/dsalgo/trie/maximizeXOR"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximizeXor(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		nums := tc["nums"].([]int)
		queries := tc["queries"].([][]int)
		expected := tc["expected"].([]int)

		actual := maximizeXOR.MaximizeXor(nums, queries)
		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"nums":     []int{0, 1, 2, 3, 4},
			"queries":  [][]int{{3, 1}, {1, 3}, {5, 6}},
			"expected": []int{3, 3, 7},
		},
		"tc2": {
			"nums":     []int{5, 2, 4, 6, 6, 3},
			"queries":  [][]int{{12, 4}, {8, 1}, {6, 3}},
			"expected": []int{15, -1, 5},
		},
	}
}
