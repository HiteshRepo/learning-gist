package combinationSubsetSumWithoutDup_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/recursion/combinationSubsetSumWithoutDup"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CombinationSum2(t *testing.T) {
	tcs := getTestcases()

	for _,tc := range tcs {
		candidates := tc["candidates"].([]int)
		target := tc["target"].(int)

		expected := tc["expected"].([][]int)

		actual := combinationSubsetSumWithoutDup.CombinationSum2(candidates, target)
		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		//"tc1": {
		//	"candidates": []int{10,1,2,7,6,1,5},
		//	"target": 8,
		//	"expected": [][]int{{1,1,6}, {1,2,5}, {1,7}, {2,6}},
		//},
		"tc2": {
			"candidates": []int{2,5,2,1,2},
			"target": 5,
			"expected": [][]int{{1,2,2}, {5}},
		},
	}
}
