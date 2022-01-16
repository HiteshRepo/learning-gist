package main

import (
	"fmt"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
)

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)

	sort(intervals)

	start := intervals[0][0]
	end := intervals[0][1]

	for i, arr := range intervals {
		if i == 0 {
			continue
		}

		if arr[0] <= end {
			if end < arr[1] {
				end = arr[1]
			}
			continue
		}

		res = append(res, []int{start, end})
		start = arr[0]
		end = arr[1]
	}

	if len(res) == 0 {
		res = append(res, []int{start, end})
	} else if res[len(res)-1][0] != start && res[len(res)-1][1] != end {
		res = append(res, []int{start, end})
	}

	return res
}

func sort(arr [][]int) {
	count := 0
	for count < len(arr) {
		least := arr[count]
		pos := count
		for i, a := range arr {
			if i < count {
				continue
			}
			if least[0] > a[0] {
				least = a
				pos = i
			}
		}
		temp := arr[count]
		arr[count] = arr[pos]
		arr[pos] = temp
		count += 1
	}

}

func runTestsForMergeIntervals() {
	testCases := map[string]map[string][][]int{
		"tc1": {
			"input":    {{1,3},{2,6},{8, 10},{15,18}},
			"expected": {{1,6},{8,10},{15,18}},
		},
		"tc2": {
			"input":    {{1,4}, {4,5}},
			"expected": {{1,5}},
		},
		"tc3": {
			"input":    {{1,4}, {0,4}},
			"expected": {{0,4}},
		},
		"tc4": {
			"input":    {{1,4}, {2, 3}},
			"expected": {{1,4}},
		},
	}

	for name, tc := range testCases {
		fmt.Printf("running test-case: %s\n", name)
		input := tc["input"]
		expected := tc["expected"]
		actual := merge(input)
		fmt.Println("is solution correct: ", utils.IntMatricesEquals(actual, expected))
	}
}

/*
Runtime: 492 ms, faster than 5.17% of Go online submissions for Merge Intervals.
Memory Usage: 7.3 MB, less than 12.55% of Go online submissions for Merge Intervals.
*/
