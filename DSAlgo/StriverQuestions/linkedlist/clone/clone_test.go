package clone_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/clone"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Clone(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		list := tc["list"].([][]*int)
		node := linkedlist.CreateNodeFromList(list)
		resHead := clone.CopyRandomList(node)
		assert.Equal(t, node, resHead)
	}
}

func Test_CloneOptimized(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		list := tc["list"].([][]*int)
		node := linkedlist.CreateNodeFromList(list)
		copyNode := linkedlist.CreateNodeFromList(list)
		resHead := clone.CopyRandomListOptimized(node)
		assert.Equal(t, copyNode, resHead)
	}
}

func getTestCases() map[string]map[string]interface{} {
	num0 := 0
	num3 := 3

	num7 := 7
	num13 := 13
	num11 := 11
	num10 := 10
	num4 := 4

	num1 := 1
	num2 := 2

	return map[string]map[string]interface{}{
		"tc1": {
			"list": [][]*int{{&num7, nil}, {&num13, &num0}, {&num11, &num4}, {&num10, &num2}, {&num1, &num0}},
		},
		"tc2": {
			"list": [][]*int{{&num3, nil}, {&num3, &num0}, {&num3, nil}},
		},
		"tc3": {
			"list": [][]*int{{&num1, &num1}, {&num2, &num1}},
		},
	}
}
