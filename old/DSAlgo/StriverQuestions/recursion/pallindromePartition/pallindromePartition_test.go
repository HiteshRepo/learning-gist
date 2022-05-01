package pallindromePartition_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PalindromePartition(t *testing.T) {
	tcs := getTestCases()

	for _,tc := range tcs {
		s := tc["s"].(string)
		expected := tc["expected"].([][]string)

		actual := Partition(s)
		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1":{
			"s": "aab",
			"expected": [][]string{{"a","a","b"}, {"aa", "b"}},
		},
		"tc2":{
			"s": "a",
			"expected": [][]string{{"a"}},
		},
		"tc3":{
			"s": "aabb",
			"expected": [][]string{{"a","a","b","b"}, {"aa","b","b"}, {"aa","bb"}, {"a","a","bb"}},
		},
	}
}