package utils

func GenerateSubArrayUsingKadanesN3(arr []int) [][]int {
	//defer Duration(Track("GenerateSubArrayUsingKadanesN3"))
	subArrs := make([][]int,0)

	for i:=0; i< len(arr); i++ {
		for j:=0; j<len(arr); j++ {
			subArr := make([]int, 0)
			for k:=i; k<=j; k++{
				subArr = append(subArr, arr[k])
				if !IntIsArrayInMatrix(subArrs, subArr) {
					subArrs = append(subArrs, subArr)
				}
			}
		}
	}

	return subArrs
}

func GenerateSubArrayUsingKadanesN2(arr []int) [][]int {
	//defer Duration(Track("GenerateSubArrayUsingKadanesN2"))
	subArrs := make([][]int,0)

	for i:=0; i< len(arr); i++ {
		subArr := make([]int, 0)
		subArr = append(subArr, arr[i])
		if !IntIsArrayInMatrix(subArrs, subArr) {
			subArrs = append(subArrs, subArr)
		}
		for j:=i+1; j<len(arr); j++ {
			subArr = append(subArr, arr[j])
			if !IntIsArrayInMatrix(subArrs, subArr) {
				subArrs = append(subArrs, subArr)
			}
		}
	}

	return subArrs
}

func GenerateSubArrayUsingKadanesN2String(arr []rune) [][]rune {
	subArrs := make([][]rune,0)

	for i:=0; i< len(arr); i++ {
		subArr := make([]rune, 0)
		subArr = append(subArr, arr[i])
		if !RuneIsArrayInMatrix(subArrs, subArr) {
			subArrs = append(subArrs, subArr)
		}
		for j:=i+1; j<len(arr); j++ {
			subArr = append(subArr, arr[j])
			if !RuneIsArrayInMatrix(subArrs, subArr) {
				subArrs = append(subArrs, subArr)
			}
		}
	}

	return subArrs
}
