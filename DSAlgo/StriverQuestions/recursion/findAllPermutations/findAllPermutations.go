package findAllPermutations

func Permute(nums []int) [][]int {
	results := make([][]int, 0)
	getAllPermutationsRecursively(nums, &results, 0)
	return results
}

func getAllPermutationsRecursively(nums []int, results *[][]int, index int) {
	if index == len(nums) {
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		*results = append(*results, newNums)
	}

	for i:=index; i<len(nums); i++ {
		swap(nums, i, index)
		getAllPermutationsRecursively(nums, results, index+1)
		swap(nums, i, index)
	}
}

func swap(nums []int, i,j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
