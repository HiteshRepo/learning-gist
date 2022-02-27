package array

import (
	"fmt"
	"log"
)

func reversePairsMergeSort(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, lb, ub int) int {
	if ub <= lb {
		return 0
	}
	mid := (lb+ub)/2

	count := mergeSort(nums, lb, mid)
	count += mergeSort(nums, mid+1, ub)

	count += merge(nums, lb, mid, ub)
	return count
}

func merge(nums []int, low, mid, high int) int {
	total := 0

	j := mid + 1
	for i := low; i <=mid; i++ {
		for j <= high && nums[i] > 2*nums[j] {
			j += 1
		}
		total += j - (mid+1)
	}

	i := low
	j = mid+1
	res := make([]int, 0)
	for i <= mid && j <= high {
		if nums[i] > nums[j] {
			res = append(res, nums[j])
			j += 1
		} else {
			res = append(res, nums[i])
			i += 1
		}
	}

	for i <= mid {
		res = append(res, nums[i])
		i += 1
	}

	for j <= high {
		res = append(res, nums[j])
		j += 1
	}

	for x := low; x <= high; x++ {
		nums[x] = res[x-low]
	}

	return total
}

func reversePairsBruteForce(nums []int) int {
	res := 0
	i := 0
	for i < len(nums) - 1 {
		j := i+1
		for j < len(nums) {
			if nums[i] > (2*nums[j]) {
				res += 1
			}
			j += 1
		}
		i += 1
	}
	return  res
}


func RunTestsForReversePairsBruteForce() {
	tcs := getTestCasesForReversePairs()

	for name, tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].(int)

		actual := reversePairsBruteForce(nums)
		log.Println(fmt.Sprintf("Is the solution %s correct: %v\n", name, expected == actual))
	}
}

func RunTestsForReversePairsMergeSort() {
	tcs := getTestCasesForReversePairs()

	for name, tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].(int)

		actual := reversePairsMergeSort(nums)
		log.Println(fmt.Sprintf("Is the solution %s correct: %v\n", name, expected == actual))
	}

}

func getTestCasesForReversePairs() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"nums": []int{1,3,2,3,1},
			"expected": 2,
		},
		"tc2": {
			"nums": []int{2,4,3,5,1},
			"expected": 3,
		},
	}
}
