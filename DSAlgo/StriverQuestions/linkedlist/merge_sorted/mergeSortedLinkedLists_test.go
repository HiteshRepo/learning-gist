package merge_sorted_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/merge_sorted"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MergeTwoListsInPlace(t *testing.T) {
	tcs := getTestCasesForMergeTwoSortedLinkedLists()

	for _, tc := range tcs {
		list1 := tc["list1"].([]int)
		list2 := tc["list2"].([]int)
		head1 := linkedlist.CreateFromArray(list1)
		head2 := linkedlist.CreateFromArray(list2)

		expected := tc["expected"].([]int)

		mergedHead :=  merge_sorted.MergeTwoListsInPlace(head1, head2)
		actual := mergedHead.GetArrayFromLinkedList()

		assert.ElementsMatch(t, expected, actual)
	}
}

func Test_MergeTwoLists(t *testing.T) {
	tcs := getTestCasesForMergeTwoSortedLinkedLists()

	for _, tc := range tcs {
		list1 := tc["list1"].([]int)
		list2 := tc["list2"].([]int)
		head1 := linkedlist.CreateFromArray(list1)
		head2 := linkedlist.CreateFromArray(list2)

		expected := tc["expected"].([]int)

		mergedHead :=  merge_sorted.MergeTwoLists(head1, head2)
		actual := mergedHead.GetArrayFromLinkedList()

		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestCasesForMergeTwoSortedLinkedLists() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"list1": []int{1,2,4},
			"list2": []int{1,3,4},
			"expected": []int{1,1,2,3,4,4},
		},
		"tc2": {
			"list1": []int{},
			"list2": []int{},
			"expected": []int{},
		},
		"tc3": {
			"list1": []int{},
			"list2": []int{0},
			"expected": []int{0},
		},
		"tc4": {
			"list1": []int{-9,3},
			"list2": []int{5,7},
			"expected": []int{-9,3,5,7},
		},
	}
}
