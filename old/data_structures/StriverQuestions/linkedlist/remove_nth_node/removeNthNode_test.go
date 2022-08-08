package remove_nth_node_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RemoveNthNode(t *testing.T) {
	tcs := getTestCasesForRemoveNthNodeFromEnd()
	for _,tc := range tcs {
		input := tc["input"].([]int)
		pos := tc["pos"].(int)
		expected := tc["expected"].([]int)

		head := linkedlist.CreateFromArray(input)

		newHead := RemoveNthFromEnd(head, pos)
		actual := newHead.GetArrayFromLinkedList()

		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestCasesForRemoveNthNodeFromEnd() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1" : {
			"input": []int{1,2,3,4,5},
			"pos": 2,
			"expected": []int{1,2,3,5},
		},
		"tc2" : {
			"input": []int{1},
			"pos": 1,
			"expected": []int{},
		},
		"tc3" : {
			"input": []int{1,2},
			"pos": 1,
			"expected": []int{1},
		},
		"tc4": {
			"input": []int{1,2},
			"pos": 2,
			"expected": []int{2},
		},
		"tc5": {
			"input": []int{5,6,7},
			"pos": 3,
			"expected": []int{6,7},
		},
	}
}