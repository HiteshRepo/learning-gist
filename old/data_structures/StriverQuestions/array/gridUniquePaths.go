package array

import "fmt"

func uniquePathsDP(m int, n int) int {
	store := make(map[string]int)
	moveAndCountDP(m, n, 0, 0, store)
	if store["0,0"] == 0 {
		return 1
	}
	return store["0,0"]
}

func moveAndCountDP(m, n, x, y int, store map[string]int) int {
	if x >= m || y >= n {
		return 0
	}

	if x == m-1 && y == n-1 {
		return 1
	}

	if val, ok := store[fmt.Sprintf("%d,%d", x, y)]; ok {
		return val
	}

	store[fmt.Sprintf("%d,%d", x, y)] = moveAndCountDP(m, n, x+1, y, store) + moveAndCountDP(m, n, x, y+1, store)

	return store[fmt.Sprintf("%d,%d", x, y)]
}

func uniquePathsRecursive(m int, n int) int {
	return moveAndCount(m, n, 0, 0)
}

func moveAndCount(m, n, x, y int) int {
	if x >= m || y >= n {
		return 0
	}

	if x == m-1 && y == n-1 {
		return 1
	}

	return moveAndCount(m, n, x+1, y) + moveAndCount(m, n, x, y+1)
}

func RunTestsForUniquePaths() {
	tcs := getTestCasesForUniquePaths()

	for name, tc := range tcs {
		m := tc["m"].(int)
		n := tc["n"].(int)
		expected := tc["expected"].(int)

		actual := uniquePathsRecursive(m, n)
		fmt.Printf("the solution for tc %s is %v\n", name, actual == expected)
	}
}

func RunTestsForUniquePathsDP() {
	tcs := getTestCasesForUniquePaths()

	for name, tc := range tcs {
		m := tc["m"].(int)
		n := tc["n"].(int)
		expected := tc["expected"].(int)

		actual := uniquePathsDP(m, n)
		fmt.Printf("the solution for tc %s is %v\n", name, actual == expected)
	}
}

func getTestCasesForUniquePaths() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"m": 3,
			"n": 7,
			"expected": 28,
		},
		"tc2": {
			"m": 3,
			"n": 2,
			"expected": 3,
		},
	}
}
