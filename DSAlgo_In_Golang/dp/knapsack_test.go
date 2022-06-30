package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMaximumWeightForGivenCapacity(t *testing.T) {
	testcases := map[string]map[string]interface{}{
		"tc1": {
			"capacity": 7,
			"weights":  []int{2, 5, 1, 3, 4},
			"items":    []int{15, 14, 10, 45, 30},
			"expected": 75,
		},
	}

	for _,tc := range testcases {
		capacity := tc["capacity"].(int)
		weights := tc["weights"].([]int)
		items := tc["items"].([]int)

		expected := tc["expected"].(int)

		actual := dp.GetMaximumWeightForGivenCapacity(capacity, weights, items)
		assert.Equal(t, expected, actual)
	}
}

func TestGetMaximumWeightForGivenCapacityByRecursion(t *testing.T) {
	testcases := map[string]map[string]interface{}{
		"tc1": {
			"capacity": 7,
			"weights":  []int{2, 5, 1, 3, 4},
			"items":    []int{15, 14, 10, 45, 30},
			"expected": 75,
		},
	}

	for _,tc := range testcases {
		capacity := tc["capacity"].(int)
		weights := tc["weights"].([]int)
		items := tc["items"].([]int)

		expected := tc["expected"].(int)

		actual := dp.GetMaximumWeightForGivenCapacityByRecursion(capacity, weights, items)
		assert.Equal(t, expected, actual)
	}
}
