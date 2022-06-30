package dp

func GetCountOfPathsToClimbStairs(n int) int {
	return countPaths(n)
}

func countPaths(n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	countPathsBySteps1 := countPaths(n - 1)
	countPathsBySteps2 := countPaths(n - 2)
	countPathsBySteps3 := countPaths(n - 3)

	totalPaths := countPathsBySteps1 + countPathsBySteps2 + countPathsBySteps3

	return totalPaths
}

func GetCountOfPathsToClimbStairsMemoized(n int) int {
	memory := make([]int, n+1)
	return countPathsMemoized(n, &memory)
}

func countPathsMemoized(n int, memory *[]int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	if (*memory)[n] > 0 {
		return (*memory)[n]
	}

	countPathsBySteps1 := countPathsMemoized(n-1, memory)
	countPathsBySteps2 := countPathsMemoized(n-2, memory)
	countPathsBySteps3 := countPathsMemoized(n-3, memory)

	totalPaths := countPathsBySteps1 + countPathsBySteps2 + countPathsBySteps3

	(*memory)[n] = totalPaths

	return totalPaths
}

func GetCountOfPathsToClimbStairsTabulation(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1

	for i := 1; i <= n; i++ {
		if i == 1 {
			dp[i] = dp[i-1]
			continue
		}

		if i == 2 {
			dp[i] = dp[i-1] + dp[i-2]
			continue
		}

		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	return dp[n]
}
