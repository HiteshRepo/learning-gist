package print_reverse_linked_list_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/linked_list/print_reverse_linked_list"
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/utils/linked_list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PrintReverseLinkedList(t *testing.T) {
	tcs := getTestcases()

	for _,tc := range tcs {
		nums := tc["nums"].([]int)
		expected := tc["expected"].([]int)

		l := linked_list.GetNewLinkedList(nums[0])
		l.AddFromArray(nums[1:])

		actual := print_reverse_linked_list.PrintReverseLinkedList(l.GetHead())
		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"nums":     []int{1, 3, 4, 7},
			"expected": []int{7, 4, 3, 1},
		},
	}
}
