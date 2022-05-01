package mColoring_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MColoring(t *testing.T) {
	tcs := getTestcases()

	for _,tc := range tcs {
		m := tc["m"].(int)
		n := tc["n"].(int)
		e := tc["e"].(int)

		edges := tc["edges"].([][]int)

		expected := tc["expected"].(bool)

		actual := MColoring(n, m, e, edges)
		assert.Equal(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"m": 3,
			"n": 4,
			"e": 5,
			"edges": [][]int{{0,1}, {1,2}, {2,3}, {3,0}, {0,2}},
			"expected": true,
		},
		"tc2": {
			"m": 2,
			"n": 3,
			"e": 3,
			"edges": [][]int{{0,1}, {1,2}, {0,2}},
			"expected": false,
		},
	}
}
