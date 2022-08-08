package dp

func GetMinCostPath(matrix [][]int) int {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)

	r := len(matrix)
	c := len(matrix[0])

	dp := make([][]int, r)
	for i:=0; i<r; i++ {
		col := make([]int, len(matrix[i]))
		dp[i] = col
	}

	dp[r-1][c-1] = matrix[r-1][c-1]

	for i:=r-1; i>=0; i-- {
		for j:=c-1; j>=0; j-- {
			minCost := MaxInt

			if isValidCell(r,c,i+1,j,dp) {
				minCost = min(minCost, dp[i+1][j])
			}

			if isValidCell(r,c,i,j+1,dp) {
				minCost = min(minCost, dp[i][j+1])
			}

			if isValidCell(r,c,i-1,j,dp) {
				minCost = min(minCost, dp[i-1][j])
			}

			if isValidCell(r,c,i,j-1,dp) {
				minCost = min(minCost, dp[i][j-1])
			}

			if i == r-1 && j == c-1 {
				continue
			}

			dp[i][j] = matrix[i][j] + minCost
		}
	}

	return dp[0][0]
}

func isValidCell(r,c,i,j int, matrix [][]int) bool {
	return i>=0 && j>=0 && i<r && j<c && matrix[i][j] != 0
}
