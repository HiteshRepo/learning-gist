package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNumberOfCoinChangePermutationPossible(t *testing.T) {
	testcases := map[string]map[string]interface{} {
		"tc1": {
			"coins": []int{2,3,5},
			"target": 7,
			"expected": 5,
		},
		"tc2": {
			"coins": []int{2,3,5,6},
			"target": 10,
			"expected": 17,
		},
	}

	for _,tc := range testcases {
		coins := tc["coins"].([]int)
		target := tc["target"].(int)
		expected := tc["expected"].(int)

		actual := dp.GetNumberOfCoinChangePermutationPossible(coins, target)
		assert.Equal(t, expected, actual)
	}
}
