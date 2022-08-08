package dp

func GetMaximumWeightForGivenCapacityUnbounded(capacity int, weightArr []int, noOfItemsArr []int) int {
	if len(weightArr) != len(noOfItemsArr) {
		return 0
	}

	dp := make([]int, capacity+1)
	dp[0] = 0

	for i:=1; i<len(dp); i++ {
		maxItems := 0
		for j:=0; j<len(weightArr); j++ {
			if i < weightArr[j] {
				continue
			}

			remainingCapacity := i - weightArr[j]
			remainingItems := dp[remainingCapacity]
			totalItems := remainingItems + noOfItemsArr[j]

			maxItems = max(maxItems, totalItems)
		}
		dp[i] = maxItems
	}

	return dp[len(dp)-1]
}
