package array

import (
	"fmt"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
)

func majorityElement(nums []int) []int {
	num1 := -1
	num2 := -1
	count1 := 0
	count2 := 0

	for _, n := range nums {
		if n == num1 {
			count1 += 1
			continue
		}

		if n == num2 {
			count2 += 1
			continue
		}

		if count1 == 0 {
			num1 = n
			count1 = 1
			continue
		}

		if count2 == 0 {
			num2 = n
			count2 = 1
			continue
		}

		count1 -= 1
		count2 -= 1
	}


	count1 = 0
	count2 = 0

	for _, n := range nums {
		if n == num1 {
			count1 += 1
			continue
		}

		if n == num2 {
			count2 += 1
		}
	}

	threshold := len(nums) / 3
	ans := make([]int, 0)
	if count1 > threshold {
		ans = append(ans, num1)
	}

	if count2 > threshold {
		ans = append(ans, num2)
	}

	return ans
}

func RunTestsForMajorityElementGt3() {
	tcs := getTestCasesForMajorityElementsGt3()

	defer utils.Duration(utils.Track("MajorityElementGt3"))
	for name, tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].([]int)

		actual := majorityElement(nums)
		fmt.Printf("is solution correct for %s: %v\n", name, utils.IntArrayEquals(expected, actual))
	}
}

func getTestCasesForMajorityElementsGt3() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"nums":     []int{3, 2, 3},
			"expected": []int{3},
		},
		"tc2": {
			"nums":     []int{1},
			"expected": []int{1},
		},
		"tc3": {
			"nums":     []int{1,2},
			"expected": []int{1,2},
		},
		"tc4": {
			"nums": []int{2,2,1,3},
			"expected": []int{2},
		},
		"tc5": {
			"nums": []int{-1,-1,-1},
			"expected": []int{-1},
		},
	}
}
