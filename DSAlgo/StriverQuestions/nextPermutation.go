package main

import (
	"fmt"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
)

func nextPermutation(nums []int) {
	if nums == nil || len(nums) == 1 {
		return
	}

	last := len(nums) - 2
	for last >= 0 && nums[last] >= nums[last+1] {
		last -= 1
	}
	index1 := last

	if index1 < 0 {
		reverse(nums, 0, len(nums)-1)
		return
	}

	last = len(nums) - 1
	for nums[last] <= nums[index1] {
		last -= 1
	}
	index2 := last

	swapPerms(nums, index1, index2)
	reverse(nums, index1+1, len(nums)-1)
}

func swapPerms(nums []int, index1, index2 int) {
	nums[index1], nums[index2] = nums[index2], nums[index1]
}

func reverse(nums []int, start, end int) {
	i := start
	j := end
	for i < j {
		swapPerms(nums, i, j)
		i += 1
		j -= 1
	}
}

func runTestsNextPermutation() {
	testCases := map[string]map[string][]int{
		"tc1": {
			"input":    {1, 2, 3},
			"expected": {1, 3, 2},
		},
		"tc2": {
			"input":    {3, 2, 1},
			"expected": {1, 2, 3},
		},
		"tc3": {
			"input":    {1, 1, 5},
			"expected": {1, 5, 1},
		},
		"tc4": {
			"input":    {5, 4, 7, 5, 3, 2},
			"expected": {5, 5, 2, 3, 4, 7},
		},
		"tc5": {
			"input":    {7, 4, 3, 1, 2, 9},
			"expected": {7, 4, 3, 1, 9, 2},
		},
	}

	for name, tc := range testCases {
		fmt.Printf("running test-case: %s\n", name)
		input := tc["input"]
		expected := tc["expected"]
		nextPermutation(input)
		fmt.Println("is solution correct: ", utils.IntArrayEquals(input, expected))
	}
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Next Permutation.
Memory Usage: 2.7 MB, less than 26.96% of Go online submissions for Next Permutation.
*/
