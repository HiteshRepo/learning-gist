package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCountOfPathsToClimbStairsWithMinimumMoves(t *testing.T) {
	testcases := map[string]map[string]interface{} {
		"tc1": {
			"n": 10,
			"jumps": []int{3,2,4,2,0,2,3,1,2,2},
			"expected": 4,
		},
	}

	for _,tc := range testcases {
		n := tc["n"].(int)
		jumps := tc["jumps"].([]int)
		expected := tc["expected"].(int)
		actual := dp.GetCountOfPathsToClimbStairsWithMinimumMoves(n, jumps)
		assert.Equal(t, expected, actual)
	}
}
