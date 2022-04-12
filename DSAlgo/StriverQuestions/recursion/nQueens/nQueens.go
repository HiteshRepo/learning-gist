package nQueens

import "fmt"

func SolveNQueens(n int) [][]string {
	ans := make([][]string, 0)
	grid := make([][]string, 0)

	for i:=0; i<n; i++ {
		temp := make([]string, 0)
		for j:=0; j<n; j++ {
			temp = append(temp, ".")
		}
		grid = append(grid, temp)
	}

	traverseBoard(&grid, &ans, 0, n)

	return ans
}

func traverseBoard(grid *[][]string, ans *[][]string, currRow, n int) {
	if currRow == n {
		res := populateGrid(*grid, n)
		*ans = append(*ans, res)
		return
	}

	for currCol:=0; currCol<n; currCol++ {
		if isValid(currCol, currRow, n, *grid) {
			(*grid)[currRow][currCol] = "Q"
			traverseBoard(grid, ans, currRow+1, n)
			(*grid)[currRow][currCol] = "."
		}
	}
}

func populateGrid(grid [][]string, n int) []string {
	res := make([]string, 0)

	for i:=0; i<n; i++ {
		temp := ""
		for j:=0; j<n; j++ {
			temp = fmt.Sprintf("%s%s", temp, grid[i][j])
		}
		res = append(res, temp)
	}

	return res
}

func isValid(currCol, currRow, n int, grid [][]string) bool {
	return isValidRow(currRow, n, grid) &&
		isValidCol(currCol, n, grid) &&
		isValidDiagonal(currCol,currRow, n, grid)
}

func isValidRow(currRow, n int, grid [][]string) bool {
	for i:=0; i<n; i++ {
		if grid[currRow][i] == "Q" {
			return false
		}
	}

	return true
}

func isValidCol(currCol, n int, grid [][]string) bool {
	for i:=0; i<n; i++ {
		if grid[i][currCol] == "Q" {
			return false
		}
	}

	return true
}

func isValidDiagonal(currCol, currRow, n int, grid [][]string) bool {
	// top left
	i := currRow
	j := currCol
	for i>=0 && j>=0 {
		if grid[i][j] == "Q" {
			return false
		}
		i -= 1
		j -= 1
	}

	// top right
	i = currRow
	j = currCol
	for i>=0 && j<n {
		if grid[i][j] == "Q" {
			return false
		}
		i -= 1
		j += 1
	}

	// bottom left
	i = currRow
	j = currCol
	for i<n && j>=0 {
		if grid[i][j] == "Q" {
			return false
		}
		i += 1
		j -= 1
	}

	// bottom right
	i = currRow
	j = currCol
	for i<n && j<n {
		if grid[i][j] == "Q" {
			return false
		}
		i += 1
		j += 1
	}

	return true
}