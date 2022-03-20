package threesum

import "github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"

func ThreeSum(nums []int) [][]int {
	res := make([][]int, 0)
	for i:=0; i<len(nums)-2; i++ {
		for j:=i+1; j<len(nums)-1; j++ {
			for k:=j+1; k<len(nums); k++ {
				if nums[i] + nums[j] + nums[k] == 0 {
					r := []int{nums[i],nums[j], nums[k]}
					if !utils.IntIsArrayInMatrix(res, r) {
						res = append(res, r)
					}
				}
			}
		}
	}
	return res
}

func ThreeSumOptimal(nums []int) [][]int {
	res := make([][]int, 0)
	utils.MergeSort(nums, 0, len(nums)-1)

	for i:=0; i<len(nums)-2; i++ {
		lo := i+1
		hi := len(nums)-1
		target := 0 - nums[i]

		for lo < hi {
			if nums[lo] + nums[hi] == target {
				res  = append(res, []int{nums[i], nums[lo], nums[hi]})
				for lo < hi && nums[lo] == nums[lo+1] {
					lo += 1
				}
				for lo < hi && nums[hi] == nums[hi-1] {
					hi -= 1
				}

				lo += 1
				hi -= 1
			}

			if nums[lo] + nums[hi] > target {
				hi -= 1
			}

			if nums[lo] + nums[hi] < target {
				lo += 1
			}
		}
	}
	return res
}
