package cumulative_sum

func GetCumulativeSum(arr []int) []int {
	ans := make([]int, len(arr))

	for i,n := range arr {
		if i == 0 {
			ans[i] = n
			continue
		}
		ans[i] = ans[i-1] + n
	}

	return ans
}
