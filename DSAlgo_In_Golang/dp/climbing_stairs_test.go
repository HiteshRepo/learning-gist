package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCountOfPathsToClimbStairs(t *testing.T) {
	testcases := map[string]map[string]int {
		"tc1": {
			"n": 10,
			"expected": 274,
		},
		"tc2": {
			"n": 15,
			"expected": 5768,
		},
	}

	for _,tc := range testcases {
		n := tc["n"]
		expected := tc["expected"]
		actual := dp.GetCountOfPathsToClimbStairs(n)
		assert.Equal(t, expected, actual)
	}
}

func TestGetCountOfPathsToClimbStairsMemoized(t *testing.T) {
	testcases := map[string]map[string]int {
		"tc1": {
			"n": 10,
			"expected": 274,
		},
		"tc2": {
			"n": 15,
			"expected": 5768,
		},
	}

	for _,tc := range testcases {
		n := tc["n"]
		expected := tc["expected"]
		actual := dp.GetCountOfPathsToClimbStairsMemoized(n)
		assert.Equal(t, expected, actual)
	}
}

func TestGetCountOfPathsToClimbStairsTabulation(t *testing.T) {
	testcases := map[string]map[string]int {
		"tc1": {
			"n": 10,
			"expected": 274,
		},
		"tc2": {
			"n": 15,
			"expected": 5768,
		},
	}

	for _,tc := range testcases {
		n := tc["n"]
		expected := tc["expected"]
		actual := dp.GetCountOfPathsToClimbStairsTabulation(n)
		assert.Equal(t, expected, actual)
	}
}
