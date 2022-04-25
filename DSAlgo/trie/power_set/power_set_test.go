package power_set_test

import (
	"github.com/hiteshrepo/learninggist/dsalgo/trie/power_set"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllPossibleStrings(t *testing.T) {
	tcs := getTestcases()

	for _,tc := range tcs {
		str := tc["str"].(string)
		expected := tc["expected"].([]string)
		actual := power_set.GetAllPossibleStrings(str)
		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"str":      "abc",
			"expected": []string{"a", "ab", "abc", "ac", "b", "bc", "", "c"},
		},
		"tc2": {
			"str":      "aa",
			"expected": []string{"a", "aa", "a", ""},
		},
	}
}
