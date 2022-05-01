package array_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/array"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetCumulativeSum(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].([]int)
		actual := array.GetCumulativeSum(nums)
		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"nums":     []int{1, 2, 3, 4},
			"expected": []int{1, 3, 6, 10},
		},
		"tc2": {
			"nums":     []int{1, 1, 1, 1, 1},
			"expected": []int{1, 2, 3, 4, 5},
		},
		"tc3": {
			"nums":     []int{1, 3, 5, 7, 9},
			"expected": []int{1, 4, 9, 16, 25},
		},
	}
}
