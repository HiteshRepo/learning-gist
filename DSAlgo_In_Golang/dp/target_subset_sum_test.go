package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsTargetSumPossible(t *testing.T) {
	testcases := map[string]map[string]interface{}{
		"tc1": {
			"candidates": []int{10, 1, 2, 7, 6, 1, 5},
			"target":     8,
			"expected":   true,
		},
		"tc2": {
			"candidates": []int{2, 5, 2, 1, 2},
			"target":     5,
			"expected":   true,
		},
	}

	for _, tc := range testcases {
		candidates := tc["candidates"].([]int)
		target := tc["target"].(int)

		expected := tc["expected"].(bool)

		actual := dp.IsTargetSumPossible(candidates, target)
		assert.Equal(t, expected, actual)
	}
}

func TestIsTargetSumPossibleOptimized(t *testing.T) {
	testcases := map[string]map[string]interface{}{
		"tc1": {
			"candidates": []int{10, 1, 2, 7, 6, 1, 5},
			"target":     8,
			"expected":   true,
		},
		"tc2": {
			"candidates": []int{2, 5, 2, 1, 2},
			"target":     5,
			"expected":   true,
		},
	}

	for _, tc := range testcases {
		candidates := tc["candidates"].([]int)
		target := tc["target"].(int)

		expected := tc["expected"].(bool)

		actual := dp.IsTargetSumPossibleOptimized(candidates, target)
		assert.Equal(t, expected, actual)
	}
}
