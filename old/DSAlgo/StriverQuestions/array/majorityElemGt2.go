package array

import (
	"fmt"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
)

func majorityElementUsingMooresVoting(nums []int) int {
	ele := 0
	count := 0

	for _, n := range nums {
		if count == 0 {
			ele = n
		}

		if ele == n {
			count += 1
		} else {
			count -= 1
		}
	}

	return ele
}

func majorityElementUsingFreqMap(nums []int) int {
	numMap := make(map[int]int)

	for _,n := range nums {
		if _, ok := numMap[n]; ok  {
			numMap[n] += 1
		} else {
			numMap[n] = 1
		}
		if numMap[n] > (len(nums) / 2) {
			return n
		}
	}

	return -1
}

func RunTestsForMajorityElementGt2UsingFreqMap() {
	tcs := getTestCasesForMajorityElementGt2()

	defer utils.Duration(utils.Track("MajorityElementGt2UsingFreqMap"))
	for name, tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].(int)

		actual := majorityElementUsingFreqMap(nums)
		fmt.Printf("is solution correct for %s: %v\n", name, expected == actual)
	}
}

func RunTestsForMajorityElementGt2UsingMooresVoting() {
	tcs := getTestCasesForMajorityElementGt2()

	defer utils.Duration(utils.Track("MajorityElementGt2UsingMooresVoting"))
	for name, tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].(int)

		actual := majorityElementUsingMooresVoting(nums)
		fmt.Printf("is solution correct for %s: %v\n", name, expected == actual)
	}
}

func getTestCasesForMajorityElementGt2() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"nums":     []int{3, 2, 3},
			"expected": 3,
		},
		"tc2": {
			"nums":     []int{2, 2, 1, 1, 1, 2, 2},
			"expected": 2,
		},
		"tc3": {
			"nums":     []int{1},
			"expected": 1,
		},
	}
}
