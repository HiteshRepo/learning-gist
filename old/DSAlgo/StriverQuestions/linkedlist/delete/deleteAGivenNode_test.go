package delete_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DeleteGivenNode(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		l1 := tc["l1"].([]int)
		node := tc["node"].(int)
		expected := tc["expected"].([]int)

		head := linkedlist.CreateFromArray(l1)

		var nodeToDelete *linkedlist.ListNode
		var tmp = head
		for tmp != nil {
			if tmp.Val == node {
				nodeToDelete = tmp
				break
			}
			tmp = tmp.Next
		}

		DeleteNode(nodeToDelete)

		actual := head.GetArrayFromLinkedList()
		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"l1": []int{4,5,1,9},
			"node": 5,
			"expected": []int{4,1,9},
		},
		"tc2": {
			"l1": []int{4,5,1,9},
			"node": 1,
			"expected": []int{4,5,9},
		},
	}
}
