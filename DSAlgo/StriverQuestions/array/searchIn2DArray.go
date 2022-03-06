package array

import (
	"fmt"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
)

// only possible if:
// The first integer of each row is greater than the last integer of the previous row.
func searchMatrixConsideringWholeMatrixAsSingleArray(matrix [][]int, target int) bool {
	l := 0
	h := (len(matrix) * len(matrix[l])) - 1

	for l < h {
		m := (l + (h-1)) / 2

		r := m/len(matrix[0])
		c := m%len(matrix[0])

		if matrix[r][c] == target {
			return true
		}

		if matrix[r][c] > target {
			h = m-1
		} else {
			l = m+1
		}
	}

	if l == h && l >= 0 && h >= 0 && l < len(matrix) && h < len(matrix[0]) {
		return matrix[l][h] == target
	}

	return false
}

func searchMatrixRowFromLastTraverse(matrix [][]int, target int) bool {
	row := 0
	column := len(matrix[row]) - 1

	for row < len(matrix) && column >= 0 {
		if matrix[row][column] == target {
			return true
		}

		if matrix[row][column] > target {
			column -= 1
		} else {
			row += 1
		}
	}

	return false
}

func searchMatrixRowBinarySearch(matrix [][]int, target int) bool {
	for _, row := range matrix {
		if utils.BinarySearch(row, target) {
			return true
		}
	}
	return false
}

func searchMatrixLinear(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for _, num := range row {
			if num == target {
				return true
			}
		}
	}
	return false
}

func RunTestsForSearchMatrixLinear() {
	testCases := getTestCases()

	defer utils.Duration(utils.Track("SearchMatrixLinear"))
	for name, tc := range testCases {
		fmt.Printf("running test-case: %s\n", name)
		matrix := tc["matrix"].([][]int)
		target := tc["target"].(int)
		expected := tc["expected"].(bool)
		actual := searchMatrixLinear(matrix, target)
		fmt.Println("is solution correct: ", expected == actual)
	}
}

// RunTestsForSearchMatrixRowFromLastTraverse possible only if array is sorted row wise and column wise
func RunTestsForSearchMatrixRowFromLastTraverse() {
	testCases := getTestCases()

	defer utils.Duration(utils.Track("SearchMatrixRowFromLastTraverse"))
	for name, tc := range testCases {
		fmt.Printf("running test-case: %s\n", name)
		matrix := tc["matrix"].([][]int)
		target := tc["target"].(int)
		expected := tc["expected"].(bool)
		actual := searchMatrixRowFromLastTraverse(matrix, target)
		fmt.Println("is solution correct: ", expected == actual)
	}
}

func RunTestsForSearchMatrixRowBinarySearch() {
	testCases := getTestCases()

	defer utils.Duration(utils.Track("SearchMatrixRowBinarySearch"))
	for name, tc := range testCases {
		fmt.Printf("running test-case: %s\n", name)
		matrix := tc["matrix"].([][]int)
		target := tc["target"].(int)
		expected := tc["expected"].(bool)
		actual := searchMatrixRowBinarySearch(matrix, target)
		fmt.Println("is solution correct: ", expected == actual)
	}
}

func RunTestsForSearchMatrixConsideringWholeMatrixAsSingleArray() {
	testCases := getTestCases()

	defer utils.Duration(utils.Track("SearchMatrixConsideringWholeMatrixAsSingleArray"))
	for name, tc := range testCases {
		fmt.Printf("running test-case: %s\n", name)
		matrix := tc["matrix"].([][]int)
		target := tc["target"].(int)
		expected := tc["expected"].(bool)
		actual := searchMatrixConsideringWholeMatrixAsSingleArray(matrix, target)
		fmt.Println("is solution correct: ", expected == actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"matrix":    [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}},
			"target":    3,
			"expected": true,
		},
		"tc2": {
			"matrix":    [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}},
			"target":    13,
			"expected": false,
		},
		"tc3": {
			"matrix":    [][]int{{1,1}},
			"target":    2,
			"expected": false,
		},
		"tc4": {
			"matrix":    [][]int{{1}},
			"target":    1,
			"expected": true,
		},
		"tc5": {
			"matrix":    [][]int{{1,3}},
			"target":    3,
			"expected": true,
		},
	}
}