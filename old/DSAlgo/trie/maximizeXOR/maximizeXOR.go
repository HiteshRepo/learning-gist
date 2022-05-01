package maximizeXOR

import "github.com/hiteshrepo/learninggist/dsalgo/trie/trie_data_structure"

func MaximizeXor(nums []int, queries [][]int) []int {
	queryMap := make(map[int][]int)
	for _,q := range queries {
		if _,ok := queryMap[q[0]]; !ok {
			queryMap[q[0]] = make([]int, 0)
		}

		for _,n := range nums {
			if n <= q[1] {
				queryMap[q[0]] = append(queryMap[q[0]], n)
			}
		}
	}

	ans := make([]int, 0)
	for _,q := range queries {
		bitsTrie := trie_data_structure.GetBitsNodeTrie()
		bitsTrie.Insert(q[0])

		maxI := 0
		for _,n := range queryMap[q[0]] {
			maxI = max(maxI, bitsTrie.GetMax(n))
		}

		if len(queryMap[q[0]]) == 0 {
			maxI = -1
		}
		ans = append(ans, maxI)
	}

	return ans
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}