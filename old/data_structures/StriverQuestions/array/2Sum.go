package array

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
	"log"
)

func twoSum(nums []int, target int) []int {
	history := make(map[int]int)

	for i, n := range nums {
		if val, ok := history[target - n]; ok {
			return []int{val, i}
		}

		history[n] = i
	}

	return []int{}
}

func twoSumBruteForce(nums []int, target int) []int {
	i := 0
	for i < len(nums) {
		j := i+1
		for j < len(nums) {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
			j += 1
		}
		i += 1
	}

	return []int{}
}

func RunTestCasesForTwoSum() {
	tcs := getTestCasesForTwoSum()

	defer utils.Duration(utils.Track("TwoSumOptimized"))
	for name, tc := range tcs {
		nums := tc["nums"].([]int)
		target := tc["target"].(int)
		expected := tc["expected"].([]int)

		actual := twoSum(nums, target)
		log.Printf("Is solution for testcase %s correct: %v", name, utils.IntArrayEquals(expected, actual))
	}
}

func RunTestCasesForTwoSumBruteForce() {
	tcs := getTestCasesForTwoSum()

	defer utils.Duration(utils.Track("TwoSumBruteForce"))
	for name, tc := range tcs {
		nums := tc["nums"].([]int)
		target := tc["target"].(int)
		expected := tc["expected"].([]int)

		actual := twoSumBruteForce(nums, target)
		log.Printf("Is solution for testcase %s correct: %v", name, utils.IntArrayEquals(expected, actual))
	}
}

func getTestCasesForTwoSum() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"nums": []int{2,7,11,15},
			"target": 9,
			"expected": []int{0,1},
		},
		"tc2": {
			"nums": []int{3,2,4},
			"target": 6,
			"expected": []int{1,2},
		},
		"tc3": {
			"nums": []int{3, 3},
			"target": 6,
			"expected": []int{0,1},
		},
	}
}
