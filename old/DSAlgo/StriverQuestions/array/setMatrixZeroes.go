package array

import (
	"fmt"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
)

func setZeroesOptimized(matrix [][]int) {
	rowArr := make([]int, len(matrix))
	columnArr := make([]int, len(matrix[0]))

	for i, row := range matrix {
		for j, elem := range row {
			if elem == 0 {
				rowArr[i] = -1
				columnArr[j] = -1
			}
		}
	}

	for i, row := range matrix {
		for j, _ := range row {
			if rowArr[i] == -1 || columnArr[j] == -1 {
				matrix[i][j] = 0
			}
		}
	}
}

func setZeroes(matrix [][]int) {
	positions := make([][]int, 0)
	for i, row := range matrix {
		for j, elem := range row {
			if elem == 0 {
				setZeroesInRow(matrix, i, &positions)
				setZeroesInCol(matrix, j, &positions)
			}
		}
	}
	for _, pos := range positions {
		matrix[pos[0]][pos[1]] = 0
	}
}

func setZeroesInRow(matrix [][]int, rowNum int, positions *[][]int) {
	i := 0
	for i < len(matrix[rowNum]) {
		*positions = append(*positions, []int{rowNum, i})
		i += 1
	}
}

func setZeroesInCol(matrix [][]int, colNum int, positions *[][]int) {
	i := 0
	for i < len(matrix) {
		*positions = append(*positions, []int{i, colNum})
		i += 1
	}
}

func RunTestsSetMatrixZeroesOptimized() {
	tcs := getTestCasesForSetMatrixZeroes()

	for name, tc := range tcs {
		fmt.Printf("running test-case: %s\n", name)
		input := tc["input"]
		expected := tc["expected"]
		setZeroesOptimized(input)
		fmt.Println("is solution correct: ", utils.IntMatricesEquals(input, expected))
	}
}

func RunTestsSetMatrixZeroes() {
	tcs := getTestCasesForSetMatrixZeroes()

	for name, tc := range tcs {
		fmt.Printf("running test-case: %s\n", name)
		input := tc["input"]
		expected := tc["expected"]
		setZeroes(input)
		fmt.Println("is solution correct: ", utils.IntMatricesEquals(input, expected))
	}
}

func getTestCasesForSetMatrixZeroes() map[string]map[string][][]int {
	return map[string]map[string][][]int {
		"tc1": {
			"input":    {{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			"expected": {{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
		"tc2": {
			"input":    {{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			"expected": {{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
	}
}

/*
164 / 164 test cases passed.
Status: Accepted
Runtime: 16 ms
Memory Usage: 8.4 MB
*/
