package uniqueSubsets_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_UniqueSubsets(t *testing.T) {
	tcs := getTestcases()

	for _, tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].([][]int)

		actual := SubsetsWithDup(nums)
		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"nums": []int{1,2,2},
			"expected": [][]int{{}, {1}, {1,2}, {1,2,2}, {2}, {2,2}},
		},
		"tc2": {
			"nums": []int{0},
			"expected": [][]int{{}, {0}},
		},
		"tc3": {
			"nums": []int{4,4,4,1,4},
			"expected": [][]int{{}, {1}, {1,4}, {1,4,4}, {1,4,4,4}, {1,4,4,4,4}, {4}, {4,4}, {4,4,4}, {4,4,4,4}},
		},
		"tc4": {
			"nums": []int{-1,1,2},
			"expected": [][]int{{}, {-1}, {-1,1}, {-1,1,2}, {-1,2}, {1}, {1,2}, {2}},
		},
	}
}
