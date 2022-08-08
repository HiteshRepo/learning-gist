package array

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
	"log"
)

func fourSumBruteOptimal(nums []int, target int) [][]int {
	utils.MergeSort(nums, 0, len(nums)-1)
	res := make([][]int, 0)

	i := 0
	for i < len(nums) {
		for j := i + 1; j < len(nums); {
			left := j + 1
			right := len(nums) - 1
			for left < right && left < len(nums) && right >= 0 && left >= 0 {
				if (nums[left] + nums[right]) == (target - nums[i] - nums[j]) {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					right = skipDuplicates(nums, right, nums[right], -1)
					left = skipDuplicates(nums, left, nums[left], 1)
					continue
				}
				if (nums[left] + nums[right]) > (target - nums[i] - nums[j]) {
					right = skipDuplicates(nums, right, nums[right], -1)
				} else {
					left = skipDuplicates(nums, left, nums[left], 1)
				}
			}
			j = skipDuplicates(nums, j, nums[j], 1)
			if j == -1 {
				break
			}
		}
		i = skipDuplicates(nums, i, nums[i], 1)
		if i == -1 {
			break
		}
	}

	return res
}

func skipDuplicates(nums []int, currPos, currVal, direction int) int {
	if direction == 1 {
		for i, n := range nums {
			if i < currPos {
				continue
			}

			if n == currVal {
				continue
			}

			return i
		}
	} else {
		i := len(nums)
		for i >= 0 {
			if i > currPos {
				i -= 1
				continue
			}

			if nums[i] == currVal {
				i -= 1
				continue
			}

			return i
		}
	}

	return -1
}

func fourSumBruteForce(nums []int, target int) [][]int {
	utils.MergeSort(nums, 0, len(nums)-1)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		j := i + 1
		for j < len(nums)-2 {
			k := j + 1
			for k < len(nums)-1 {
				pos := utils.BinarySearchPos(nums[k+1:], target-(nums[i]+nums[j]+nums[k]))
				if pos > -1 {
					currRes := []int{nums[i], nums[j], nums[k], nums[pos+k+1]}
					if !utils.IntIsArrayInMatrix(res, currRes) {
						res = append(res, currRes)
					}
				}
				k += 1
			}
			j += 1
		}
	}

	return res
}

func RunTestsFor4SumBruteForce() {
	tcs := getTestCasesFor4Sum()

	defer utils.Duration(utils.Track("FourSum - 3 ptr, binary search"))
	for name, tc := range tcs {
		nums := tc["nums"].([]int)
		target := tc["target"].(int)
		expected := tc["expected"].([][]int)

		actual := fourSumBruteForce(nums, target)
		log.Printf("Is the solution of testcase %s correct: %v", name, utils.IntMatricesElementsEquals(expected, actual))
	}
}

func RunTestsFor4SumOptimal() {
	tcs := getTestCasesFor4Sum()

	defer utils.Duration(utils.Track("FourSum - 4 ptr - optimal"))
	for name, tc := range tcs {
		nums := tc["nums"].([]int)
		target := tc["target"].(int)
		expected := tc["expected"].([][]int)

		actual := fourSumBruteOptimal(nums, target)
		log.Printf("Is the solution of testcase %s correct: %v", name, utils.IntMatricesElementsEquals(expected, actual))
	}
}

func getTestCasesFor4Sum() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		//"tc1": {
		//	"nums":     []int{1, 0, -1, 0, -2, 2},
		//	"target":   0,
		//	"expected": [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		//},
		//"tc2": {
		//	"nums": []int{2,2,2,2,2},
		//	"target": 8,
		//	"expected": [][]int{{2,2,2,2}},
		//},
		//"tc3": {
		//	"nums": []int{2,-4,-5,-2,-3,-5,0,4,-2},
		//	"target": -14,
		//	"expected": [][]int{{-5,-5,-4,0},{-5,-5,-2,-2},{-5,-4,-3,-2}},
		//},
		//"tc4": {
		//	"nums": []int{-2,-1,-1,1,1,2,2},
		//	"target": 0,
		//	"expected": [][]int{{-2,-1,1,2},{-1,-1,1,1}},
		//},
		"tc5": {
			"nums": []int{0,0,0,0},
			"target": 1,
			"expected": [][]int{},
		},
	}
}
