package array

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
	"log"
)

func lengthOfLongestSubstringOptimized(s string) int {
	maxLen := 0
	set := utils.NewRuneSet()
	chars := []rune(s)


	l := 0
	r := 0
	for l < len(chars) && r < len(chars) && l <= r  {
		c := chars[r]
		if set.Contains(c) {
			for set.Contains(c) {
				set.Remove(chars[l])
				l += 1
			}
		}

		if !set.Contains(c) {
			currLen := r-l+1
			if currLen > maxLen {
				maxLen = currLen
			}
		}

		set.Add(c)
		r += 1
	}

	return maxLen
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	chars := []rune(s)
	subArrs := utils.GenerateSubArrayUsingKadanesN2String(chars)

	maxLen := -1

	for _, arr := range subArrs {
		if doesDuplicatesExistsInThisArray(arr) {
			continue
		}

		if len(arr) > maxLen {
			maxLen = len(arr)
		}
	}

	return maxLen
}

func doesDuplicatesExistsInThisArray(arr []rune) bool {
	charMap := make(map[rune]struct{})

	for _, c := range arr {
		if _, ok := charMap[c]; ok {
			return true
		} else {
			charMap[c] = struct{}{}
		}
	}

	return false
}

func RunTestsForLengthOfLongestSubstringOptimized() {
	tcs := getTestsForLengthOfLongestSubstring()

	for name, t := range tcs {
		s := t["s"].(string)
		expected := t["expected"].(int)

		actual := lengthOfLongestSubstringOptimized(s)

		log.Printf("Is the solution for %s correct: %v", name, expected == actual)
	}
}

func RunTestsForLengthOfLongestSubstring() {
	tcs := getTestsForLengthOfLongestSubstring()

	for name, t := range tcs {
		s := t["s"].(string)
		expected := t["expected"].(int)

		actual := lengthOfLongestSubstring(s)

		log.Printf("Is the solution for %s correct: %v", name, expected == actual)
	}
}

func getTestsForLengthOfLongestSubstring() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"s": "abcabcbb",
			"expected": 3,
		},
		"tc2": {
			"s": "bbbbb",
			"expected": 1,
		},
		"tc3": {
			"s": "pwwkew",
			"expected": 3,
		},
	}
}
