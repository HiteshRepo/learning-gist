package amazon_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/amazon"
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/linked_list/print_linked_list"
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/utils/linked_list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRetainRemove(t *testing.T) {
	tcs := getTestcasesForRetainRemove()

	for _, tc := range tcs {
		list := tc["list"].([]int)
		m := tc["m"].(int)
		n := tc["n"].(int)
		l := linked_list.GetNewLinkedList(list[0])
		l.AddFromArray(list[1:])

		expected := tc["expected"].([]int)

		amazon.RetainRemove(l.GetHead(), m, n)
		actual := print_linked_list.PrintLinkedList(l.GetHead())

		assert.Equal(t, expected, actual)
	}
}

func getTestcasesForRetainRemove() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"list": []int{1,2,3,4,5,6,7,8},
			"m": 2,
			"n": 2,
			"expected": []int{1,2,5,6},
		},
	}
}
