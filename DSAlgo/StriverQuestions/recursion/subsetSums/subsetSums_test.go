package subsetSums_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/recursion/subsetSums"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SubsetSumsByRecursion(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		arr := tc["arr"].([]int)
		N := tc["N"].(int)

		expected := tc["expected"].([]int)

		actual := subsetSums.SubsetSumsByRecursion(arr, N)
		assert.Equal(t, expected, actual)
	}
}

func Test_SubsetSumsByPowerSet(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		arr := tc["arr"].([]int)
		N := tc["N"].(int)

		expected := tc["expected"].([]int)

		actual := subsetSums.SubsetSumsByPowerSet(arr, N)
		assert.Equal(t, expected, actual)
	}
}

func Test_SubsetSumsByPowerBit(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		arr := tc["arr"].([]int)
		N := tc["N"].(int)

		expected := tc["expected"].([]int)

		actual := subsetSums.SubsetSumsByPowerBit(arr, N)
		assert.Equal(t, expected, actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"arr": []int{2,3},
			"N": 2,
			"expected": []int{0,2,3,5},
		},
		"tc2": {
			"arr": []int{5,2,1},
			"N": 3,
			"expected": []int{0,1,2,3,5,6,7,8},
		},
	}
}
