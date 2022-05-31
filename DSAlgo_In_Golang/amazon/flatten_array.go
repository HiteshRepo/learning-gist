package amazon

// Given an array consisting many inner arrays, flatten the array into one: example: input: [[6,4,7,[9,5,4,[2,4,8]]],[2,2,7],[9,0,7,[9,3,1,8,5]]] output: [6,4,7,9,5,4,2,4,8,2,2,7,9,0,7,9,3,1,8,5]
func Flatten(arr [][]int) []int {
	ans := make([]int, 0)

	for _, a := range arr {
		ans = append(ans, a...)
	}

	return ans
}
