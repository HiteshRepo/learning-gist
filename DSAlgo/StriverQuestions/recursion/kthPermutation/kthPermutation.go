package kthPermutation

import (
	"fmt"
	"sort"
)

func GetPermutationOptimal(n int, k int) string {
	numStr := ""
	res := ""

	for i:=1; i<=n; i++ {
		numStr = fmt.Sprintf("%s%s", numStr, fmt.Sprintf("%d", i))
	}

	currK := k-1
	currStr := numStr

	for true {
		if len(res) == len(numStr) {
			break
		}

		noOfBlocks := len(currStr)
		noOfSequencesInEachBlock := factorial(noOfBlocks - 1)
		blockNumber := currK / noOfSequencesInEachBlock
		res = fmt.Sprintf("%s%s", res, string(currStr[blockNumber]))
		currStr = string(append([]rune(currStr[0:blockNumber]), []rune(currStr[blockNumber+1:])...))
		currK = currK % noOfSequencesInEachBlock
	}

	return res
}

func factorial(n int) int {
	res := 1
	for n > 0 {
		res *= n
		n -= 1
	}

	return res
}

func GetPermutation(n int, k int) string {
	numStr := ""

	for i:=1; i<=n; i++ {
		numStr = fmt.Sprintf("%s%s", numStr, fmt.Sprintf("%d", i))
	}

	allPermutations := make([]string, 0)
	getAllPermutations([]rune(numStr), 0, &allPermutations)
	sort.Strings(allPermutations)

	return allPermutations[k-1]
}

func getAllPermutations(s []rune, index int, allPermutations *[]string) {
	if index == len(s) {
		*allPermutations = append(*allPermutations, string(s))
	}

	for i:=index; i<len(s); i++ {
		swap(s, i, index)
		getAllPermutations(s, index+1, allPermutations)
		swap(s, i, index)
	}
}

func swap(s []rune, i,j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}