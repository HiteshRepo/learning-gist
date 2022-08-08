package findAllPermutations_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetAllPermutations(t *testing.T) {
	tcs := getTestcases()

	for _, tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].([][]int)

		actual := Permute(nums)
		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"nums": []int{1,2,3},
			"expected": [][]int{{1,2,3}, {1,3,2}, {2,1,3}, {2,3,1}, {3,1,2}, {3,2,1}},
		},
		"tc2": {
			"nums": []int{0,1},
			"expected": [][]int{{0,1}, {1,0}},
		},
		"tc3": {
			"nums": []int{1},
			"expected": [][]int{{1}},
		},
	}
}
