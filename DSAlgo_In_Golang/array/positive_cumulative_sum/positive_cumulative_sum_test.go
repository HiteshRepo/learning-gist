package positive_cumulative_sum_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/array/positive_cumulative_sum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetPositiveCumulativeSum(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].([]int)
		actual := positive_cumulative_sum.GetPositiveCumulativeSum(nums)
		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"nums":     []int{1, -2, 3, 4, -6},
			"expected": []int{1, 2, 6},
		},
		"tc2": {
			"nums":     []int{1, -1, -1, -1, 1},
			"expected": []int{1},
		},
		"tc3": {
			"nums":     []int{1, 3, 5, 7},
			"expected": []int{1, 4, 9, 16},
		},
	}
}
