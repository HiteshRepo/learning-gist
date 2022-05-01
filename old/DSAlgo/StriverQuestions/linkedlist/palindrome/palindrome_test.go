package palindrome_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsPalindrome(t *testing.T) {
	tcs := getTestcases()

	for _, tc := range tcs {
		list := tc["list"].([]int)
		expected := tc["expected"].(bool)

		head := linkedlist.CreateFromArray(list)

		actual := IsPalindrome(head)
		assert.Equal(t, expected, actual)
	}
}

func Test_IsPalindromeOptimized(t *testing.T) {
	tcs := getTestcases()

	for _, tc := range tcs {
		list := tc["list"].([]int)
		expected := tc["expected"].(bool)

		head := linkedlist.CreateFromArray(list)

		actual := IsPalindromeOptimized(head)
		assert.Equal(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"list": []int{1,2,2,1},
			"expected": true,
		},
		"tc2": {
			"list": []int{1,2},
			"expected": false,
		},
	}
}
