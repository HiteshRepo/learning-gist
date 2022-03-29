package uniqueSubsets

import (
	"fmt"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
	"strconv"
	"strings"
)

func SubsetsWithDup(nums []int) [][]int {
	utils.MergeSort(nums, 0, len(nums)-1)
	cloneArr := make([]*int, 0)
	for _, n := range nums {
		num := n
		cloneArr = append(cloneArr, &num)
	}

	subsets := make([][]int, 0)
	subset := make([]*int, len(nums))
	subsetMap := make(map[string]struct{})
	recursivelyGetAllSubsets(cloneArr, subset, 0, &subsets, subsetMap)

	return subsets
}

func recursivelyGetAllSubsets(nums, subset []*int, idx int, subsets *[][]int, subsetMap map[string]struct{})  {
	if idx == len(nums) {
		newSubset := removeNullFromArray(subset)
		strArr := arrToStr(newSubset)
		if _, ok := subsetMap[strArr]; !ok {
			*subsets = append(*subsets, newSubset)
			subsetMap[strArr] = struct{}{}
		}
		return
	}

	subset[idx] = nil
	recursivelyGetAllSubsets(nums, subset, idx+1, subsets, subsetMap)
	subset[idx] = nums[idx]
	recursivelyGetAllSubsets(nums, subset, idx+1, subsets, subsetMap)
}

func removeNullFromArray(nums []*int) []int {
	newNums := make([]int, 0)
	for _, n := range nums {
		if n != nil {
			newNums = append(newNums, *n)
		}
	}

	return newNums
}

func arrToStr(arr []int) string {
	str := ""
	for _,n := range arr {
		str = fmt.Sprintf("%s%d", str,n)
	}
	return str
}

func strToArr(str string) []int {
	strArr := strings.Split(str, "")
	nums := make([]int, 0)
	negative := false
	for _,s := range strArr {
		if s == "-" {
			negative = true
			continue
		}
		n, _ := strconv.Atoi(s)
		if negative {
			nums = append(nums, n*-1)
			negative = false
		} else {
			nums = append(nums, n)
		}
	}
	return nums
}
