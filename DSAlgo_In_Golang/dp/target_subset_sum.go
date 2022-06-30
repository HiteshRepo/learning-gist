package dp

func IsTargetSumPossibleOptimized(candidates []int, target int) bool {
	dp := make([][]int, len(candidates) + 1)
	for i:=0; i<len(dp); i++ {
		col := make([]int, target+1)
		dp[i] = col
	}

	for i:=0; i<len(dp); i++ {
		for j:=0; j<len(dp[i]); j++ {

			if i==0 && j==0 {
				dp[i][j] = 1
				continue
			}
			if i==0 {
				continue
			}
			if j==0 {
				dp[i][j] = 1
				continue
			}
			if dp[i-1][j] == 1 {
				dp[i][j] = 1
				continue
			}

			val := candidates[i-1]
			if j >= val && dp[i-1][j-val] == 1 {
				dp[i][j] = 1
			}
		}
	}

	for _,d := range dp {
		if d[len(d)-1] == 1 {
			return true
		}
	}

	return false
}


func IsTargetSumPossible(candidates []int, target int) bool {
	MergeSort(candidates, 0, len(candidates)-1)
	subsets := make([][]int, 0)
	subset := make([]*int, 0)

	cloneArr := make([]*int, 0)
	for _,n := range candidates {
		num := n
		cloneArr = append(cloneArr, &num)
	}

	getAllSubsetsRecursively(cloneArr, subset, 0, &subsets, target)

	return len(subsets) > 0
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

func MergeSort(nums []int, lb, ub int) {
	if ub <= lb {
		return
	}
	mid := (lb+ub)/2

	MergeSort(nums, lb, mid)
	MergeSort(nums, mid+1, ub)

	merge(nums, lb, mid, ub)
}

func merge(nums []int, low, mid, high int) {
	i := low
	j := mid+1
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
}