package middle_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MiddleNode(t *testing.T) {
	tcs := getTestCasesForMiddleNode()

	for _, tc := range tcs {
		input := tc["input"].([]int)
		head := linkedlist.CreateFromArray(input)

		expected := tc["expected"].([]int)

		middleNode :=  MiddleNode(head)
		actual := middleNode.GetArrayFromLinkedList()

		assert.ElementsMatch(t, expected, actual)
	}
}


func Test_MiddleNodeOptimized(t *testing.T)  {
	tcs := getTestCasesForMiddleNode()

	for _, tc := range tcs {
		input := tc["input"].([]int)
		head := linkedlist.CreateFromArray(input)

		expected := tc["expected"].([]int)

		middleNode :=  MiddleNodeOptimized(head)
		actual := middleNode.GetArrayFromLinkedList()

		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestCasesForMiddleNode() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"input": []int{1,2,3,4,5},
			"expected": []int{3,4,5},
		},
		"tc2": {
			"input": []int{1,2,3,4,5,6},
			"expected": []int{4,5,6},
		},
	}
}