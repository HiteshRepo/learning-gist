package utils

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
