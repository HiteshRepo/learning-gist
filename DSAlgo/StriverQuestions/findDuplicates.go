package main

import (
	"fmt"
	"sort"
)

func findDuplicate(nums []int) int {
	return usingSorting(nums)
}

func usingSorting(nums []int) int {
	sort.Ints(nums)

	for i:=0; i<len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}

	return -1
}

func usingFrequencyMap(nums []int) int {
	numMap := make(map[int]int)

	for _,n := range nums {
		if exists(n, numMap) {
			numMap[n] = numMap[n]+1
			continue
		}
		numMap[n] = 1
	}

	for k,v := range numMap {
		if v > 1 {
			return k
		}
	}

	return -1
}

func exists(n int, numMap map[int]int) bool {
	for k,_ := range numMap {
		if k == n {
			return true
		}
	}
	return false
}

func runTestsFindDuplicates() {
	testCases := map[string]map[string]interface{}{
		"tc1": {
			"input":    []int{1,3,4,2,2},
			"expected": 2,
		},
		"tc2": {
			"input":    []int{3,1,3,4,2},
			"expected": 3,
		},
		"tc3": {
			"input":    []int{1,3,4,2,1},
			"expected": 1,
		},
	}

	for name, tc := range testCases {
		fmt.Printf("running test-case: %s\n", name)
		input := tc["input"].([]int)
		expected := tc["expected"].(int)
		actual := findDuplicate(input)
		fmt.Println("is solution correct: ", actual == expected)
	}
}

/*
Using sorting:

Runtime: 348 ms, faster than 5.34% of Go online submissions for Find the Duplicate Number.
Memory Usage: 9 MB, less than 53.38% of Go online submissions for Find the Duplicate Number.
*/
