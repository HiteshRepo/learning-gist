package combinationSubsetSum

func CombinationSum(candidates []int, target int) [][]int {
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
	if idx == len(arr) {
		newSubset := removeNullFromArr(subset)
		if target == 0 {
			*subsets = append(*subsets, newSubset)
		}
		return
	}

	if *arr[idx] <= target {
		subset = append(subset, arr[idx])
		getAllSubsetsRecursively(arr, subset, idx, subsets, target-*arr[idx])
		subset = subset[0:len(subset)-1]
	}

	getAllSubsetsRecursively(arr, subset, idx+1, subsets, target)
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