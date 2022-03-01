package array

import (
	"log"
)

func largestSubArrWithSumZero(nums []int) int {
	store := make(map[int]int)

	prefixSum := 0
	maxLen := 0
	for i, n := range nums {
		prefixSum += n

		idx, ok := store[prefixSum]
		if prefixSum == 0 {
			maxLen = i+1
		}

		if ok {
			currLen := i - idx
			if currLen > maxLen {
				maxLen = currLen
			}
		} else {
			store[prefixSum] = i
		}
	}

	return maxLen
}


func RunTestsForLargestSubArrWithSumZero() {
	tcs := getTestCasesForLargestSubArrWithSumZero()

	for name, t := range tcs {
		nums := t["nums"].([]int)
		expected := t["expected"].(int)

		actual := largestSubArrWithSumZero(nums)
		log.Printf("Is the solution for the testcase %s correct: %v", name, expected == actual)
	}
}

func getTestCasesForLargestSubArrWithSumZero() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"nums": []int{1, -1, 3, 2, -2, -8, 1, 7, 10, 2, 3},
			"expected": 5,
		},
	}
}
