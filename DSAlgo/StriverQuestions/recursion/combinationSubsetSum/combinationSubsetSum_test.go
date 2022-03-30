package combinationSubsetSum_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/recursion/combinationSubsetSum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CombinationSubsetSum(t *testing.T) {
	tcs := getTestCases()

	for _,tc := range tcs {
		candidates := tc["candidates"].([]int)
		target := tc["target"].(int)

		expected := tc["expected"].([][]int)

		actual := combinationSubsetSum.CombinationSum(candidates, target)
		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"candidates": []int{2,3,6,7},
			"target": 7,
			"expected": [][]int{{2,2,3}, {7}},
		},
		"tc2": {
			"candidates": []int{2,3,5},
			"target": 8,
			"expected": [][]int{{2,2,2,2}, {2,3,3}, {3,5}},
		},
		"tc3": {
			"candidates": []int{2},
			"target": 1,
			"expected": [][]int{},
		},
	}
}