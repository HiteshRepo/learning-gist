package dp

func GetNumberOfCoinChangeCombinationPossible(sortedCoins []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for _,coin := range sortedCoins {
		for i,_ := range dp {
			if i<coin {
				continue
			}
			if i >= coin && dp[i-coin] == 1 {
				dp[i] += dp[i-coin]
			}
		}
	}

	return dp[target]
}

func GetNumberOfCoinChangeCombinationPossibleByRecursion(sortedCoins []int, target int) int {
	noOfCombinations := 0
	getNumberOfCoinChangeCombinationPossibleByRecursion(sortedCoins, 0, target, &noOfCombinations)
	return noOfCombinations
}

func getNumberOfCoinChangeCombinationPossibleByRecursion(sortedCoins []int, index int, target int, noOfCombinations *int) {
	if target == 0 {
		*noOfCombinations = (*noOfCombinations) + 1
		return
	}

	if index == len(sortedCoins) {
		return
	}

	currentCoin := sortedCoins[index]
	if currentCoin > target {
		return
	}

	getNumberOfCoinChangeCombinationPossibleByRecursion(sortedCoins, index, target-currentCoin, noOfCombinations) // pick
	getNumberOfCoinChangeCombinationPossibleByRecursion(sortedCoins, index+1, target, noOfCombinations) // not pick
}
