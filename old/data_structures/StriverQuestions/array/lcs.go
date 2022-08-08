package array

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
	"log"
)

func longestConsecutiveOptimized(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxLen := -1

	numsClone := make(map[int]struct{})
	for _, n := range nums {
		numsClone[n] = struct{}{}
	}

	for _, n := range nums {
		if _, ok := numsClone[n+1]; ok {
			continue
		}

		x := n
		count := 0
		_, ok := numsClone[x]
		for ok {
			count += 1
			x -= 1
			_, ok = numsClone[x]
		}

		if count > maxLen {
			maxLen = count
		}

		count = 0
	}

	return maxLen
}

func longestConsecutiveNaive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	utils.MergeSort(nums, 0, len(nums)-1)

	currSeq := make([]int, 0)
	maxLen := -1

	j := 1
	currSeq = append(currSeq, nums[j-1])
	for j < len(nums) {
		if nums[j-1] == nums[j] {
			j+=1
			continue
		}
		if (nums[j-1] + 1) != nums[j] {
			if len(currSeq) > maxLen {
				maxLen = len(currSeq)
			}
			currSeq = make([]int, 0)
		}
		currSeq = append(currSeq, nums[j])
		j += 1
	}

	if len(currSeq) > maxLen {
		maxLen = len(currSeq)
	}

	return maxLen
}

func RunTestsForLongestConsecutiveNaive() {
	tcs := getTestCasesForLongestConsecutive()

	for name, t := range tcs {
		nums := t["nums"].([]int)
		expected := t["expected"].(int)

		actual := longestConsecutiveNaive(nums)
		log.Printf("Is the solution for testcase %s correct: %v", name, actual == expected)
	}

}

func RunTestsForLongestConsecutiveOptimized() {
	tcs := getTestCasesForLongestConsecutive()

	for name, t := range tcs {
		nums := t["nums"].([]int)
		expected := t["expected"].(int)

		actual := longestConsecutiveOptimized(nums)
		log.Printf("Is the solution for testcase %s correct: %v", name, actual == expected)
	}

}

func getTestCasesForLongestConsecutive() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1":  {
			"nums": []int{100,4,200,1,3,2},
			"expected": 4,
		},
		"tc2":  {
			"nums": []int{0,3,7,2,5,8,4,6,0,1},
			"expected": 9,
		},
		"tc3": {
			"nums": []int{1,2,0,1},
			"expected": 3,
		},
		"tc4": {
			"nums": []int{1,3,5,2,4},
			"expected": 5,
		},
	}
}
