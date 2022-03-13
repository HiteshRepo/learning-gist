package consecutiveOnes

func FindMaxConsecutiveOnes(nums []int) int {
	maxOnes := 0
	currOnes := 0

	for _, n := range nums {
		if n == 1 {
			currOnes += 1
		} else {
			currOnes = 0
		}

		if currOnes > maxOnes {
			maxOnes = currOnes
		}
	}

	return maxOnes
}
