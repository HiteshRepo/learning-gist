package dp

func GetNumberOfCoinChangePermutationPossible(sortedCoins []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i,_ := range dp {
		for _,coin := range sortedCoins {
			if i<coin {
				continue
			}
			if i >= coin && dp[i-coin] > 0 {
				dp[i] += dp[i-coin]
			}
		}
	}

	return dp[target]
}
