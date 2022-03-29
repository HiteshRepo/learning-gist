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

func IntArrayElementsEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range a {
		if !IntIsElementIn(b, v) {
			return false
		}
	}
	return true
}

func IntArrayElementsEqualsExact(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func RuneArrayElementsEqualsExact(a []rune, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func IntIsElementIn(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}

	return false
}

func IntMatricesElementsEquals(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, row1 := range a {
		if !IntIsArrayInMatrix(b, row1) {
			return false
		}
	}

	return true
}

func IntIsArrayInMatrix(a [][]int, b []int) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}

	for _, row1 := range a {
		if IntArrayElementsEqualsExact(b, row1) {
			return true
		}
	}

	return false
}

func IntIsArrayInMatrixUnExact(a [][]int, b []int) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}

	for _, row1 := range a {
		if IntArrayElementsEquals(b, row1) {
			return true
		}
	}

	return false
}

func RuneIsArrayInMatrix(a [][]rune, b []rune) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}

	for _, row1 := range a {
		if RuneArrayElementsEqualsExact(b, row1) {
			return true
		}
	}

	return false
}

func BinarySearch(row []int, target int) bool {
	l := 0
	r := len(row) - 1

	for l < r {
		m := (l + (r-1)) / 2
		if row[m] == target {
			return true
		}

		if row[m] < target {
			l = m+1
		} else {
			r = m-1
		}
	}

	if l == r {
		return row[l] == target
	}

	return false
}

func BinarySearchPos(row []int, target int) int {
	l := 0
	r := len(row) - 1

	for l < r {
		m := (l + (r-1)) / 2
		if row[m] == target {
			return m
		}

		if row[m] < target {
			l = m+1
		} else {
			r = m-1
		}
	}

	if l == r && row[l] == target{
		return l
	}

	return -1
}

func MergeSort(nums []int, lb, ub int) {
	if ub <= lb {
		return
	}
	mid := (lb+ub)/2

	MergeSort(nums, lb, mid)
	MergeSort(nums, mid+1, ub)

	merge(nums, lb, mid, ub)
}

func merge(nums []int, low, mid, high int) {
	i := low
	j := mid+1
	res := make([]int, 0)
	for i <= mid && j <= high {
		if nums[i] > nums[j] {
			res = append(res, nums[j])
			j += 1
		} else {
			res = append(res, nums[i])
			i += 1
		}
	}

	for i <= mid {
		res = append(res, nums[i])
		i += 1
	}

	for j <= high {
		res = append(res, nums[j])
		j += 1
	}

	for x := low; x <= high; x++ {
		nums[x] = res[x-low]
	}
}