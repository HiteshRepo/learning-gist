package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCountOfPathsToClimbStairsWithJumpsAndTabulation(t *testing.T) {
	testcases := map[string]map[string]interface{} {
		"tc1": {
			"n": 6,
			"jumps": []int{3,3,0,2,2,3},
			"expected": 8,
		},
	}

	for _,tc := range testcases {
		n := tc["n"].(int)
		jumps := tc["jumps"].([]int)
		expected := tc["expected"].(int)
		actual := dp.GetCountOfPathsToClimbStairsWithJumpsAndTabulation(n, jumps)
		assert.Equal(t, expected, actual)
	}
}
