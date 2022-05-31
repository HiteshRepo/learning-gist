package amazon

// // You are given an array of integers. Return an array of the same size where the element at each index is the product of all the elements in the original array except for the element at that index. For example, an input of [1, 2, 3, 4, 5] should return [120, 60, 40, 30, 24]. You cannot use division in this problem.
func ProductArray(arr []int) []int {

	// return naive(arr)
	return efficient(arr)

}

func naive(arr []int) []int {
	l := make([]int, len(arr))
	l[0] = 1
	r := make([]int, len(arr))
	r[len(arr) - 1] = 1

	prod := make([]int, len(arr))

	for i:=1; i<len(arr); i++ {
		l[i] = l[i-1]*arr[i-1]
	}

	for i:=len(arr)-2; i>=0; i-- {
		r[i] = r[i+1]*arr[i+1]
	}

	for i:=0; i<len(arr); i++ {
		prod[i] = l[i]*r[i]
	}

	return prod
}

func efficient(arr []int) []int {
	prod := make([]int, len(arr))
	temp := 1
	for i:=0; i<len(arr); i++ {
		prod[i] = temp
		temp = temp * arr[i]
	}

	temp = 1
	for i:=len(arr)-1; i>=0; i-- {
		prod[i] *= temp
		temp *= arr[i]
	}

	return prod
}
