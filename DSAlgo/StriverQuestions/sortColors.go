package main

import (
	"fmt"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
)

func sortColorsSubmission1(nums []int) {
	zeroes := 0
	ones := 0
	twoes := 0

	for _, n := range nums {
		if n == 0 {
			zeroes += 1
		}

		if n == 1 {
			ones += 1
		}

		if n == 2 {
			twoes += 1
		}
	}

	i := 0
	for zeroes > 0 {
		nums[i] = 0
		zeroes -= 1
		i += 1
	}

	for ones > 0 {
		nums[i] = 1
		ones -= 1
		i += 1
	}

	for twoes > 0 {
		nums[i] = 2
		twoes -= 1
		i += 1
	}
}

func sortColorsSubmission2(nums []int) {
	low := 0
	mid := 0
	high := len(nums) - 1

	for mid <= high {
		if nums[mid] == 0 {
			swapColors(low, mid, nums)
			low += 1
			mid += 1
		} else if nums[mid] == 1 {
			mid += 1
		} else if nums[mid] == 2 {
			swapColors(high, mid, nums)
			high -= 1
		}
	}

}

func swapColors(i, j int, nums []int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func runTestsSortColors() {
	testCases := map[string]map[string][]int{
		"tc1": {
			"input":    {2, 0, 2, 1, 1, 0},
			"expected": {0, 0, 1, 1, 2, 2},
		},
		"tc2": {
			"input":    {2, 0, 1},
			"expected": {0, 1, 2},
		},
	}

	for name, tc := range testCases {
		fmt.Printf("running test-case: %s\n", name)
		input := tc["input"]
		expected := tc["expected"]
		//sortColorsSubmission1(input)
		sortColorsSubmission2(input)
		fmt.Println("is solution correct: ", utils.IntArrayEquals(input, expected))
	}
}

/*
Submission-1:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
Memory Usage: 2.1 MB, less than 20.32% of Go online submissions for Sort Colors.

Submission-2:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
Memory Usage: 2.2 MB, less than 7.75% of Go online submissions for Sort Colors.
*/
