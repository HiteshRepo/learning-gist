package nQueens_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/recursion/nQueens"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NQueens(t *testing.T) {
	tcs := getTestcases()

	for _,tc := range tcs {
		n := tc["n"].(int)
		expected := tc["expected"].([][]string)

		actual := nQueens.SolveNQueens(n)
		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"n": 4,
			"expected": [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
		"tc2": {
			"n": 1,
			"expected": [][]string{{"Q"}},
		},
	}
}
