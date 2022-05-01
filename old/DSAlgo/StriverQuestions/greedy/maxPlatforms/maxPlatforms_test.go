package maxPlatforms_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MaxPlatforms(t *testing.T) {
	tcs := getTestcases()

	for _,tc := range tcs {
		n := tc["n"].(int)
		arr := tc["arr"].([]int)
		dep := tc["dep"].([]int)

		expected := tc["expected"].(int)

		actual := FindPlatforms(arr, dep, n)
		assert.Equal(t, expected, actual)
	}
}

func Test_MaxPlatformsOptimal(t *testing.T) {
	tcs := getTestcases()

	for _,tc := range tcs {
		n := tc["n"].(int)
		arr := tc["arr"].([]int)
		dep := tc["dep"].([]int)

		expected := tc["expected"].(int)

		actual := FindPlatformsOptimal(arr, dep, n)
		assert.Equal(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"n": 6,
			"arr": []int{900, 940, 950, 1100, 1500, 1800},
			"dep": []int{910, 1200, 1120, 1130, 1900, 2000},
			"expected": 3,
		},
		"tc2": {
			"n": 3,
			"arr": []int{900, 1100, 1235},
			"dep": []int{1000, 1200, 1240},
			"expected": 1,
		},
	}
}
