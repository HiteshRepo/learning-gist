package amazon

// Find Square Root of a whole number without using standard functions
func FindSquareRoot(num int) float64 {

	found := false
	var ans float64
	i := 1
	for !found {

		if i*i == num {
			found = true
			ans = float64(i)
		} else if i*i > num {
			ans = Square(float64(i-1), float64(i), float64(num))
			found = true
		}

		i += 1

	}
	return ans
}

func Square(i, j, num float64) float64 {
	mid := (i + j) / 2
	mul := mid * mid

	if mul == num || (mul - num) < 0.00001 {
		return mid
	}

	if mul < num {
		return Square(mid, j, num)
	} else {
		return Square(i, mid, num)
	}
}
