package reverseInGroups_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/reverseInGroups"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReverseInGroups(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		list := tc["list"].([]int)
		groupSize := tc["groupSize"].(int)
		head := linkedlist.CreateFromArray(list)

		expected := tc["expected"].([]int)

		resHead := reverseInGroups.ReverseKGroup(head, groupSize)
		actual := resHead.GetArrayFromLinkedList()

		assert.Equal(t, expected, actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"list": []int{1,2,3,4,5},
			"groupSize": 2,
			"expected": []int{2,1,4,3,5},
		},
		"tc2": {
			"list": []int{1,2,3,4,5},
			"groupSize": 3,
			"expected": []int{3,2,1,4,5},
		},
	}
}
