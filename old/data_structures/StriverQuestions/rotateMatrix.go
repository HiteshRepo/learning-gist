package main

import (
	"fmt"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
)

func rotate(matrix [][]int)  {
	for i:=0; i<len(matrix); i++ {
		for j:=i; j<len(matrix[0]); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}

	for i:=0; i<len(matrix); i++ {
		for j:=0; j<(len(matrix)/2); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[i][len(matrix)-1-j]
			matrix[i][len(matrix)-1-j] = temp
		}
	}
}

func runTestsForRotateMatrix() {
	testCases := map[string]map[string][][]int{
		"tc1": {
			"input":    {{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}},
			"expected": {{15,13,2,5},{14,3,4,1},{12,6,8,9},{16,7,10,11}},
		},
		"tc2": {
			"input":    {{1,2,3},{4,5,6},{7,8,9}},
			"expected": {{7,4,1}, {8,5,2}, {9,6,3}},
		},
	}

	for name, tc := range testCases {
		fmt.Printf("running test-case: %s\n", name)
		input := tc["input"]
		expected := tc["expected"]
		rotate(input)
		fmt.Println("is solution correct: ", utils.IntMatricesEquals(input, expected))
	}
}

/*
Runtime: 2 ms, faster than 21.54% of Go online submissions for Rotate Image.
Memory Usage: 2.5 MB, less than 5.53% of Go online submissions for Rotate Image.
 */