package pallindromePartition

func Partition(s string) [][]string {
	ans := make([][]string, 0)
	subset := make([]string, 0)
	traverseStringRecursively(s, &ans, subset, 0)
	return ans
}

func traverseStringRecursively(s string, ans *[][]string, subset []string, partitionIdx int) {
	if partitionIdx == len(s) {
		newSubset := make([]string, len(subset))
		copy(newSubset, subset)
		*ans = append(*ans, newSubset)
		return
	}

	for i:=partitionIdx; i<len(s); i++ {
		currStr := s[partitionIdx:i+1]
		if len(currStr) == 1 || isStringPalindrome(currStr) {
			subset = append(subset, currStr)
			traverseStringRecursively(s, ans, subset, i+1)
			subset = subset[0:len(subset)-1]
		}
	}
}

func isStringPalindrome(s string) bool {
	l := 0
	r := len(s)-1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l += 1
		r -= 1
	}

	return true
}
