package combinationSubsetSumWithoutDup

import "github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"

func CombinationSum2(candidates []int, target int) [][]int {
	utils.MergeSort(candidates, 0, len(candidates)-1)
	subsets := make([][]int, 0)
	subset := make([]*int, 0)

	cloneArr := make([]*int, 0)
	for _,n := range candidates {
		num := n
		cloneArr = append(cloneArr, &num)
	}

	getAllSubsetsRecursively(cloneArr, subset, 0, &subsets, target)

	return subsets
}

func getAllSubsetsRecursively(arr, subset []*int, idx int, subsets *[][]int, target int) {
	if target == 0 {
		newSubset := removeNullFromArr(subset)
		*subsets = append(*subsets, newSubset)
		return
	}

	for i:=idx; i<len(arr); i++ {
		if i!=idx && *arr[i] == *arr[i-1] {
			continue
		}
		if *arr[i] > target {
			break
		}
		subset = append(subset, arr[i])
		getAllSubsetsRecursively(arr, subset, i+1, subsets, target - *arr[i])
		subset = subset[0:len(subset)-1]
	}
}

func removeNullFromArr(arr []*int) []int {
	resArr := make([]int, 0)

	for _, n := range arr {
		if n != nil {
			resArr = append(resArr, *n)
		}
	}

	return resArr
}