package removeDuplicates_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RemoveDuplicates(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].(int)

		actual := RemoveDuplicates(nums)

		assert.Equal(t, expected, actual)
	}
}

func Test_RemoveDuplicatesOptimized(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].(int)

		actual := RemoveDuplicatesOptimized(nums)

		assert.Equal(t, expected, actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"nums": []int{1,1,2},
			"expected": 2,
		},
		"tc2": {
			"nums": []int{0,0,1,1,1,2,2,3,3,4},
			"expected": 5,
		},
	}
}
