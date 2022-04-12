package sudokuSolver

import "fmt"

func SolveSudoku(board [][]byte)  {
	traverseBoard(&board, 0, 0)
}

func traverseBoard(board *[][]byte, currRow, currCol int) bool {
	if currRow == 9 {
		return true
	}

	nextRow := 0
	nextCol := 0

	if currCol == 8 {
		nextRow = currRow + 1
		nextCol = 0
	} else {
		nextRow = currRow
		nextCol = currCol + 1
	}

	if (*board)[currRow][currCol] != []byte(".")[0] {
		return traverseBoard(board, nextRow, nextCol)
	}

	for currVal:=1; currVal<10; currVal++ {
		if isValidCell(currRow, currCol, currVal, *board) {
			(*board)[currRow][currCol] = []byte(fmt.Sprintf("%d", currVal))[0]

			if traverseBoard(board, nextRow, nextCol) {
				return true
			}

			(*board)[currRow][currCol] = []byte(".")[0]
		}
	}

	return false
}

func isValidCell(currRow, currCol, currVal int, board [][]byte) bool {
	return isValidRow(currRow, currVal, board) &&
		isValidCol(currCol, currVal, board) &&
		isValidGrid(currRow, currCol, currVal, board)
}

func isValidRow(currRow, currVal int, board [][]byte) bool {
	for j:=0; j<9; j++ {
		if string(board[currRow][j]) == fmt.Sprintf("%d", currVal) {
			return false
		}
	}
	return true
}

func isValidCol(currCol, currVal int, board [][]byte) bool {
	for i:=0; i<9; i++ {
		if string(board[i][currCol]) == fmt.Sprintf("%d", currVal) {
			return false
		}
	}
	return true
}

func isValidGrid(currRow, currCol, currVal int, board [][]byte) bool {
	x := 3*(currRow/3)
	y := 3*(currCol/3)

	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			if string(board[x+i][y+j]) == fmt.Sprintf("%d", currVal) {
				return false
			}
		}
	}
	return true
}
