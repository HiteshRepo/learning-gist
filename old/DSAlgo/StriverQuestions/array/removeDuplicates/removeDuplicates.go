package removeDuplicates

func RemoveDuplicates(nums []int) int {
	numMap := make(map[int]struct{})

	for _, n := range nums {
		if _, ok := numMap[n]; !ok {
			numMap[n] = struct{}{}
		}
	}

	i := 0
	for n, _ := range numMap {
		nums[i] = n
		i += 1
	}

	return len(numMap)
}

func RemoveDuplicatesOptimized(nums []int) int {
	i := 0
	for j:=1; j<len(nums); j++ {
		if nums[i] != nums[j] {
			i += 1
			nums[i] = nums[j]
		}
	}
	return i+1
}
