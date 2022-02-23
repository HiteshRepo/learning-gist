package utils

import (
	"math"
)


func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func IntMatricesEquals(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, row := range a {
		if len(row) != len(b[i]) {
			return false
		}

		for j, elem := range a[i] {
			if elem != b[i][j] {
				return false
			}
		}
	}
	return true
}

func CheckFloatEquals(x, y float64) bool {
	tolerance := 0.0000001
	diff := math.Abs(x - y)
	mean := math.Abs(x + y) / 2.0
	factor := diff / mean
	return factor < tolerance
}
