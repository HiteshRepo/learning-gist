package array

import (
	"fmt"
)

func maxSubArray(nums []int) int {

	msf := nums[0]
	meh := 0

	for _, n := range nums {
		meh += n

		if meh > msf {
			msf = meh
		}

		if meh < 0 {
			meh = 0
		}
	}

	return msf
}

func runTestsKadanes() {
	testCases := map[string]map[string]interface{}{
		"tc1": {
			"input":    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			"expected": 6,
		},
		"tc2": {
			"input":    []int{5, 4, -1, 7, 8},
			"expected": 23,
		},
	}

	for name, tc := range testCases {
		fmt.Printf("running test-case: %s\n", name)
		input := tc["input"].([]int)
		expected := tc["expected"].(int)
		actual := maxSubArray(input)
		fmt.Println("is solution correct: ", expected == actual)
	}
}

/*
Runtime: 141 ms, faster than 22.40% of Go online submissions for Maximum Subarray.
Memory Usage: 10.2 MB, less than 8.09% of Go online submissions for Maximum Subarray.
*/
