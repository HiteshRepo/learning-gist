package trappingRainwater

func TrapRainwater(height []int) int {
	trap := 0
	for i, _ := range height {
		lmax := 0
		rmax := 0

		j := i
		for j >= 0 {
			lmax = max(lmax, height[j])
			j -= 1
		}

		j = i
		for j < len(height) {
			rmax = max(rmax, height[j])
			j += 1
		}

		trap += min(lmax, rmax) - height[i]
	}

	return trap
}

func max(a, b int ) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int ) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func TrapRainwaterOptimal(height []int) int {
	trap := 0

	lmaxArr := make([]int, len(height))
	rmaxArr := make([]int, len(height))

	lmaxArr[0] = height[0]
	for i,n := range height {
		if i == 0 {
			continue
		}
		lmaxArr[i] = max(lmaxArr[i-1], n)
	}

	rmaxArr[len(height) - 1] = height[len(height)-1]
	j := len(height) - 2
	for j >= 0 {
		rmaxArr[j] = max(rmaxArr[j+1], height[j])
		j -= 1
	}

	for i, n := range height {
		trap += min(lmaxArr[i], rmaxArr[i]) - n
	}

	return trap
}
