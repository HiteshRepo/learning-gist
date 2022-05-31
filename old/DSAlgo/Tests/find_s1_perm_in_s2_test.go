package Tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindS1PermutationInS2(t *testing.T) {
	tcs := getTestCasesForFindS1PermutationInS2()

	for _,tc := range tcs {
		s1 := tc["s1"].(string)
		s2 := tc["s2"].(string)
		expected := tc["expected"].(bool)

		actual := FindS1PermutationInS2(s1, s2)
		assert.Equal(t, expected, actual)
	}
}

func TestFindS1PermutationInS2Optimal(t *testing.T) {
	tcs := getTestCasesForFindS1PermutationInS2()

	for _,tc := range tcs {
		s1 := tc["s1"].(string)
		s2 := tc["s2"].(string)
		expected := tc["expected"].(bool)

		actual := FindS1PermutationInS2Optimal(s1, s2)
		assert.Equal(t, expected, actual)
	}
}

func getTestCasesForFindS1PermutationInS2() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"s1": "abc",
			"s2": "xyzbacmty",
			"expected": true,
		},
		"tc2": {
			"s1": "abcd",
			"s2": "xyzbacmty",
			"expected": false,
		},
	}
}
