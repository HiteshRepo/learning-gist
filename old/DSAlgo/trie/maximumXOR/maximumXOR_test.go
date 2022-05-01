package maximumXOR_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaximumXOR(t *testing.T) {
	tcs := getTestcases()

	for _,tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].(int)

		actual := FindMaximumXOR(nums)
		assert.Equal(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"nums": []int{3,10,5,25,2,8},
			"expected": 28,
		},
		"tc2": {
			"nums": []int{14,70,53,83,49,91,36,80,92,51,66,70},
			"expected": 127,
		},
	}
}
