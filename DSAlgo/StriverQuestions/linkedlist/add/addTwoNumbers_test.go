package add_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/add"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AddTowNumbers(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		l1 := tc["l1"].([]int)
		l2 := tc["l2"].([]int)

		head1 := linkedlist.CreateFromArray(l1)
		head2 := linkedlist.CreateFromArray(l2)

		expected := tc["expected"].([]int)
		resHead := add.AddTwoNumbers(head1, head2)
		actual := resHead.GetArrayFromLinkedList()

		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"l1": []int{2,4,3},
			"l2": []int{5,6,4},
			"expected": []int{7,0,8},
		},
		"tc2": {
			"l1": []int{0},
			"l2": []int{0},
			"expected": []int{0},
		},
		"tc3": {
			"l1": []int{9,9,9,9,9,9,9},
			"l2": []int{9,9,9,9},
			"expected": []int{8,9,9,9,0,0,0,1},
		},
		"tc4": {
			"l1": []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1},
			"l2": []int{5,6,4},
			"expected": []int{6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1},
		},
	}
}
