package kthPermutation_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_KthPermutation(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		n := tc["n"].(int)
		k := tc["k"].(int)

		expected := tc["expected"].(string)

		actual := GetPermutation(n, k)
		assert.Equal(t, expected, actual)
	}

}

func Test_KthPermutationOptimal(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		n := tc["n"].(int)
		k := tc["k"].(int)

		expected := tc["expected"].(string)

		actual := GetPermutationOptimal(n, k)
		assert.Equal(t, expected, actual)
	}

}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"n": 3,
			"k": 3,
			"expected": "213",
		},
		"tc2": {
			"n": 4,
			"k": 9,
			"expected": "2314",
		},
		"tc3": {
			"n": 3,
			"k": 1,
			"expected": "123",
		},
	}
}
