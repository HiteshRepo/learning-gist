package dp

func GetMaxGoldPath(matrix [][]int) int {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	r := len(matrix)
	c := len(matrix[0])

	dp := make([][]int, r)
	for i:=0; i<r; i++ {
		col := make([]int, len(matrix[i]))
		col[len(matrix[i]) - 1] = matrix[i][len(matrix[i]) - 1]
		dp[i] = col
	}

	for i:=c-2; i>=0; i-- {
		for j:=0; j<r; j++ {
			maxCost := MinInt

			if isValidCell(r,c,j,i+1,dp) {
				maxCost = max(maxCost, dp[j][i+1])
			}

			if isValidCell(r,c,j-1,i+1,dp) {
				maxCost = max(maxCost, dp[j-1][i+1])
			}

			if isValidCell(r,c,j+1,i+1,dp) {
				maxCost = max(maxCost, dp[j+1][i+1])
			}

			dp[j][i] = maxCost + matrix[j][i]
		}
	}

	maxCost := MinInt
	for i:=0; i<len(dp); i++ {
		maxCost = max(maxCost, dp[i][0])
	}

	return maxCost
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}
