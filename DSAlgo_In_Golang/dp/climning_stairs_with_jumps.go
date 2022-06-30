package dp

func GetCountOfPathsToClimbStairsWithJumpsAndTabulation(n int, jumps []int) int {
	dp := make([]int, n+1)
	dp[n] = 1

	for i:=n-1; i>=0; i-- {
		for j:=1; j<=jumps[i]; j++ {
			if i+j < len(dp) {
				dp[i] += dp[i+j]
			}
		}
	}

	return dp[0]
}