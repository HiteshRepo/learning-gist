package sudokuSolver_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/recursion/sudokuSolver"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SudokuSolver(t *testing.T) {
	tcs := getTestcases()

	for _, tc := range tcs {
		input := tc["input"].([][]string)

		board := make([][]byte, 0)
		for _, row := range input {
			byteRow := make([]byte, 0)
			for _, j := range row {
				byteRow = append(byteRow, []byte(j)...)
			}
			board = append(board, byteRow)
		}

		sudokuSolver.SolveSudoku(board)
		expected := tc["expected"].([][]string)

		actual := make([][]string, 0)
		for _, row := range board {
			strRow := make([]string, 0)
			for _, j := range row {
				strRow = append(strRow, string(j))
			}
			actual = append(actual, strRow)
		}

		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"input": [][]string{{"5","3",".",".","7",".",".",".","."},{"6",".",".","1","9","5",".",".","."},{".","9","8",".",".",".",".","6","."},{"8",".",".",".","6",".",".",".","3"},{"4",".",".","8",".","3",".",".","1"},{"7",".",".",".","2",".",".",".","6"},{".","6",".",".",".",".","2","8","."},{".",".",".","4","1","9",".",".","5"},{".",".",".",".","8",".",".","7","9"}},
			"expected": [][]string{{"5","3","4","6","7","8","9","1","2"},{"6","7","2","1","9","5","3","4","8"},{"1","9","8","3","4","2","5","6","7"},{"8","5","9","7","6","1","4","2","3"},{"4","2","6","8","5","3","7","9","1"},{"7","1","3","9","2","4","8","5","6"},{"9","6","1","5","3","7","2","8","4"},{"2","8","7","4","1","9","6","3","5"},{"3","4","5","2","8","6","1","7","9"}},
		},
	}
}
