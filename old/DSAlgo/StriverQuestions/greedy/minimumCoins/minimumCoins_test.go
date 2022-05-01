package minimumCoins_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MinCoins(t *testing.T) {
	tcs := getTestCases()

	for _,tc := range tcs {
		denominations := tc["denominations"].([]int)
		value := tc["value"].(int)

		expected := tc["expected"].(int)

		actual := MinCoins(denominations, value)

		assert.Equal(t, expected, actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"denominations": []int{25, 10, 5},
			"value": 30,
			"expected": 2,
		},
		"tc2": {
			"denominations": []int{9, 6, 5, 1},
			"value": 11,
			"expected": 2,
		},
	}
}
