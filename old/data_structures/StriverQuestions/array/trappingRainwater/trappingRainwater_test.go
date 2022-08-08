package trappingRainwater_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TrapRainwater(t *testing.T) {
	tcs := getTestcases()

	for _, tc := range tcs {
		height := tc["height"].([]int)
		expected := tc["expected"].(int)

		actual := TrapRainwater(height)
		assert.Equal(t, expected, actual)
	}
}

func Test_TrapRainwaterOptimal(t *testing.T) {
	tcs := getTestcases()

	for _, tc := range tcs {
		height := tc["height"].([]int)
		expected := tc["expected"].(int)

		actual := TrapRainwaterOptimal(height)
		assert.Equal(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"height": []int{0,1,0,2,1,0,1,3,2,1,2,1},
			"expected": 6,
		},
		"tc2": {
			"height": []int{4,2,0,3,2,5},
			"expected": 9,
		},
	}
}
