package amazon

import "math"

// Reverse the bits of an 32 bit unsigned integer A.
func ReverseBits(num int64) int64 {
	bits := make([]int, 32)
	for i:=31; i>=0; i-- {
		bit := (num >> i) & 1
		if bit == 0 {
			bits[31-i] = 0
		} else {
			bits[31-i] = 1
		}
	}

	l := 0
	r := 31
	for l < r {
		swap(l, r, bits)
		l += 1
		r -= 1
	}

	var ans int64
	for i,b := range bits {
		if b == 1 {
			ans += int64(math.Pow(float64(2), float64(31-i)))
		}
	}

	return ans
}

func swap(i,j int, arr []int) {
	arr[i], arr[j] = arr[j], arr[i]
}
