package consecutiveOnes_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/array/consecutiveOnes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MaxConsecutiveOnes(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].(int)

		actual := consecutiveOnes.FindMaxConsecutiveOnes(nums)
		assert.Equal(t, expected, actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"nums": []int{1,1,0,1,1,1},
			"expected": 3,
		},
		"tc2": {
			"nums": []int{1,0,1,1,0,1},
			"expected": 2,
		},
	}
}
