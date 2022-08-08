package Tests

import (
	"sort"
	"strings"
)

func FindS1PermutationInS2(s1, s2 string) bool {
	perms := make([]string, 0)
	getAllPerms([]rune(s1), &perms, 0)

	for _,p := range perms {
		if exists(p, s2) {
			return true
		}
	}

	return false
}

func FindS1PermutationInS2Optimal(s1, s2 string) bool {
	indexes := make([]int, 0)
	freqMap := make(map[rune]int)

	for _,c := range s1 {
		if _,ok := freqMap[c]; !ok {
			freqMap[c] = 0
		}
		freqMap[c] += 1
	}

	for i,c := range s2 {
		if count,ok := freqMap[c]; ok && count > 0 {
			freqMap[c] -= 1
			indexes = append(indexes, i)
		}
	}

	if len(indexes) != len(s1) {
		return false
	}

	sort.Ints(indexes)

	prevIndex := indexes[0]
	for i:=1; i<len(indexes); i++ {
		if indexes[i] - prevIndex > 1 {
			return false
		}
		prevIndex = indexes[i]
	}

	return true
}

func exists(str1, str2 string) bool {
	splits := strings.Split(str2, str1)
	return len(splits) > 1
}

func getAllPerms(s1 []rune, perms *[]string, index int) {
	if index == len(s1) {
		*perms = append(*perms, string(s1))
	}
	for i:=index; i<len(s1); i++ {
		swap(i, index, s1)
		getAllPerms(s1, perms, index+1)
		swap(i, index, s1)
	}

}

func swap(i, j int, str []rune) {
	str[i], str[j] = str[j], str[i]
}