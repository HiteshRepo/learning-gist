package amazon

// Given an array of distinct elements and a number x, find if there is a pair with product equal to x.
func FindProductPair(nums []int, product int) (int, int) {
	numMap := make(map[int]int)

	for i,n := range nums {
		required := product/n
		if j,ok:= numMap[required]; ok {
			return j,i
		}
		numMap[n] = i
	}


	return -1,-1
}
