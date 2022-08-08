package array

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
	"log"
)

func findSubArrForGivenXorOptimized(A []int , B int )  int {
	count := 0
	xorStore := make(map[int]int)

	prefixXor := 0
	for _, n := range A {
		prefixXor = utils.IntCalculateXOR(prefixXor, n)
		if prefixXor == B {
			count += 1
		}

		supplementaryXor := utils.IntCalculateXOR(B, prefixXor)
		if val, ok := xorStore[supplementaryXor]; ok {
			count += val
		}

		if val, ok := xorStore[prefixXor]; ok {
			xorStore[prefixXor] = val + 1
		} else {
			xorStore[prefixXor] = 1
		}
	}

	return count
}

func findSubArrForGivenXor(A []int , B int )  int {
	subArrs := utils.GenerateSubArrayUsingKadanesN2(A)
	results := make([][]int, 0)

	for _, arr := range subArrs {
		xor := findXorForArray(arr)
		if xor == B || (len(arr) == 1 && arr[0] == B) {
			results = append(results, arr)
		}
	}

	return len(results)
}


func findXorForArray(arr []int) int {
	i := 0
	xorUntil := 0
	for i < len(arr) {
		xorUntil = utils.IntCalculateXOR(xorUntil, arr[i])
		i += 1
	}

	return xorUntil
}

func RunTestsForFindSubArrForGivenXor() {
	tcs := getTestCasesForFindSubArrForGivenXor()

	for name, tc := range tcs {
		A := tc["A"].([]int)
		B := tc["B"].(int)
		expected := tc["expected"].(int)

		actual := findSubArrForGivenXor(A, B)

		log.Printf("Is the solution correct for test case %s : %v", name, actual == expected)
	}
}

func RunTestsForFindSubArrForGivenXorOptimized() {
	tcs := getTestCasesForFindSubArrForGivenXor()

	for name, tc := range tcs {
		A := tc["A"].([]int)
		B := tc["B"].(int)
		expected := tc["expected"].(int)

		actual := findSubArrForGivenXorOptimized(A, B)

		log.Printf("Is the solution correct for test case %s : %v", name, actual == expected)
	}
}

func getTestCasesForFindSubArrForGivenXor() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"A": []int{4, 2, 2, 6, 4},
			"B": 6,
			"expected": 4,
		},
		"tc2": {
			"A": []int{5, 6, 7, 8, 9},
			"B": 5,
			"expected": 2,
		},
	}
}
