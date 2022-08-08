package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNthFibonacciNumber(t *testing.T) {
	testcases := map[string]map[string]int {
		"tc1": {
			"n": 10,
			"expected": 55,
		},
		"tc2": {
			"n": 15,
			"expected": 610,
		},
	}

	for _,tc := range testcases {
		n := tc["n"]
		expected := tc["expected"]
		actual := dp.GetNthFibonacciNumber(n)
		assert.Equal(t, expected, actual)
	}
}

func TestGetNthFibonacciNumberOptimized(t *testing.T) {
	testcases := map[string]map[string]int {
		"tc1": {
			"n": 10,
			"expected": 55,
		},
		"tc2": {
			"n": 15,
			"expected": 610,
		},
	}

	for _,tc := range testcases {
		n := tc["n"]
		expected := tc["expected"]
		actual := dp.GetNthFibonacciNumberOptimized(n)
		assert.Equal(t, expected, actual)
	}
}
