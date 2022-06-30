package dp

func GetMaximumWeightForGivenCapacity(capacity int, weightArr []int, noOfItemsArr []int) int {
	if len(weightArr) != len(noOfItemsArr) {
		return 0
	}

	dp := make([][]int, len(weightArr)+1)
	for i:=0; i<len(weightArr)+1; i++ {
		cols := make([]int, capacity+1)
		dp[i] = cols
	}

	for i:=1; i<len(weightArr)+1; i++ {
		for j:=1; j<capacity+1; j++ {
			withoutWeightArrAtI := dp[i-1][j]
			if j < weightArr[i-1] {
				dp[i][j] = withoutWeightArrAtI
				continue
			}
			withWeightArrAtI := noOfItemsArr[i-1] + dp[i-1][j-weightArr[i-1]]
			dp[i][j] = max(withWeightArrAtI, withoutWeightArrAtI)
		}
	}

	lastRow := dp[len(dp)-1]
	lastCell := lastRow[len(lastRow)-1]
	return lastCell
}

func GetMaximumWeightForGivenCapacityByRecursion(capacity int, weightArr []int, noOfItemsArr []int) int {
	if len(weightArr) != len(noOfItemsArr) {
		return 0
	}

	allItemsCount := make([]int, 0)
	getMaximumWeightForGivenCapacityByRecursion(capacity, 0, weightArr, noOfItemsArr, &allItemsCount, 0)

	if len(allItemsCount) == 0 {
		return 0
	}

	maxCount := allItemsCount[0]
	for _,count := range allItemsCount {
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}

func getMaximumWeightForGivenCapacityByRecursion(capacity, idx int, weightArr []int, noOfItemsArr []int, allItemsCount *[]int, currItemsCount int) {
	if capacity == 0 {
		*allItemsCount = append(*allItemsCount, currItemsCount)
		return
	}

	if idx >= len(weightArr) {
		return
	}

	getMaximumWeightForGivenCapacityByRecursion(capacity-weightArr[idx], idx+1, weightArr, noOfItemsArr, allItemsCount, currItemsCount+noOfItemsArr[idx])
	getMaximumWeightForGivenCapacityByRecursion(capacity, idx+1, weightArr, noOfItemsArr, allItemsCount, currItemsCount)
}