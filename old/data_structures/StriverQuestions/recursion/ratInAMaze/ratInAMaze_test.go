package ratInAMaze_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/recursion/ratInAMaze"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RatInAMaze(t *testing.T) {
	tcs := getTestcases()

	for _, tc := range tcs {
		maze := tc["maze"].([][]int)
		expected := tc["expected"].([]string)

		actual := ratInAMaze.FindAllPaths(maze)
		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"maze": [][]int{{1,0,0,0}, {1,1,0,1}, {1,1,0,0}, {0,1,1,1}},
			"expected": []string{"DDRDRR", "DRDDRR"},
		},
		"tc2": {
			"maze": [][]int{{1,0}, {1,0}},
			"expected": []string{},
		},
	}
}
