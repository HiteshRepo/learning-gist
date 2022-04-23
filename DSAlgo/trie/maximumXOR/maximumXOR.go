package maximumXOR

import "github.com/hiteshrepo/learninggist/dsalgo/trie/trie_data_structure"

func FindMaximumXOR(nums []int) int {
	bitsTrie := trie_data_structure.GetBitsNodeTrie()
	for _,n := range nums {
		bitsTrie.Insert(n)
	}

	maxI := 0
	for _,n := range nums {
		maxI = max(maxI, bitsTrie.GetMax(n))
	}
	return maxI
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
