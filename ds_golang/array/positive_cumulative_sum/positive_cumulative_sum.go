package positive_cumulative_sum

func GetPositiveCumulativeSum(nums []int) []int {
	ans := make([]int, 0)
	cumulativeSums := make([]int, len(nums))

	for i,n := range nums {
		if i == 0 {
			cumulativeSums[i] = n
		} else {
			cumulativeSums[i] = cumulativeSums[i-1] + n
		}

		if cumulativeSums[i] > 0 {
			ans = append(ans, cumulativeSums[i])
		}
	}

	return ans
}
