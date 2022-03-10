package reverse_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/reverse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReverseLinkedListIterative(t *testing.T)  {
	tcs := getTestCases()

	for _, tc := range tcs {
		input := tc["input"].([]int)
		head := linkedlist.CreateFromArray(input)

		expected := tc["expected"].([]int)

		reversedHead :=  reverse.ReverseListIterative(head)
		actual := reversedHead.GetArrayFromLinkedList()

		assert.ElementsMatch(t, expected, actual)
	}
}

func Test_ReverseLinkedListRecursive(t *testing.T)  {
	tcs := getTestCases()

	for _, tc := range tcs {
		input := tc["input"].([]int)
		head := linkedlist.CreateFromArray(input)

		expected := tc["expected"].([]int)

		reversedHead :=  reverse.ReverseListRecursive(head)
		actual := reversedHead.GetArrayFromLinkedList()

		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"input": []int{1,2,3,4,5},
			"expected": []int{5,4,3,2,1},
		},
		"tc2": {
			"input": []int{1,2},
			"expected": []int{2,1},
		},
	}
}
