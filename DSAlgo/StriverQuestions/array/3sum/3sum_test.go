package threesum_test

import (
	threesum "github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/array/3sum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ThreeSum(t *testing.T) {
	tcs := getTestcases()

	for _, tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].([][]int)

		actual := threesum.ThreeSum(nums)
		assert.Equal(t, expected, actual)
	}
}

func Test_ThreeSumOptimal(t *testing.T) {
	tcs := getTestcases()

	for _, tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].([][]int)

		actual := threesum.ThreeSumOptimal(nums)
		assert.Equal(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"nums": []int{-1,0,1,2,-1,-4},
			"expected": [][]int{{-1,-1,2},{-1,0,1}},
		},
		"tc2": {
			"nums": []int{},
			"expected": [][]int{},
		},
		"tc3": {
			"nums": []int{0},
			"expected": [][]int{},
		},
	}
}
