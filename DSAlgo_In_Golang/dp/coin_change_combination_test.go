package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNumberOfCoinChangeCombinationPossible(t *testing.T) {
	testcases := map[string]map[string]interface{} {
		"tc1": {
			"coins": []int{2,3,5},
			"target": 7,
			"expected": 2,
		},
	}

	for _,tc := range testcases {
		coins := tc["coins"].([]int)
		target := tc["target"].(int)
		expected := tc["expected"].(int)

		actual := dp.GetNumberOfCoinChangeCombinationPossible(coins, target)
		assert.Equal(t, expected, actual)
	}
}

func TestGetNumberOfCoinChangeCombinationPossibleByRecursion(t *testing.T) {
	testcases := map[string]map[string]interface{} {
		"tc1": {
			"coins": []int{2,3,5},
			"target": 7,
			"expected": 2,
		},
	}

	for _,tc := range testcases {
		coins := tc["coins"].([]int)
		target := tc["target"].(int)
		expected := tc["expected"].(int)

		actual := dp.GetNumberOfCoinChangeCombinationPossibleByRecursion(coins, target)
		assert.Equal(t, expected, actual)
	}
}
