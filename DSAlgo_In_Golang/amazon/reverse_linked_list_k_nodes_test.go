package amazon_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/amazon"
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/linked_list/print_linked_list"
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/utils/linked_list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	tcs := getTestcases()

	for _, tc := range tcs {
		list := tc["list"].([]int)
		groupSize := tc["groupSize"].(int)
		l := linked_list.GetNewLinkedList(list[0])
		l.AddFromArray(list[1:])

		expected := tc["expected"].([]int)

		resHead := amazon.ReverseKGroup(l.GetHead(), groupSize)
		actual := print_linked_list.PrintLinkedList(resHead)

		assert.Equal(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
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
