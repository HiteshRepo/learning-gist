package dp

func GetCountOfPathsToClimbStairsWithMinimumMoves(n int, jumps []int) int {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)

	dp := make([]int, n+1)
	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		dp[i] = MaxInt - 1
	}

	for i := n - 1; i >= 0; i-- {
		m := dp[i]
		for j := 1; j <= jumps[i]; j++ {
			if i+j < len(dp) {
				m = min(m, dp[i+j])
			}
		}
		dp[i] = m + 1
	}

	return dp[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
