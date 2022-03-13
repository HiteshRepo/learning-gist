package rotate_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/rotate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RotateList(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		list := tc["list"].([]int)
		k := tc["k"].(int)

		expected := tc["expected"].([]int)

		head := linkedlist.CreateFromArray(list)
		resHead := rotate.RotateRight(head, k)

		actual := resHead.GetArrayFromLinkedList()

		assert.Equal(t, expected, actual)
	}
}

func Test_RotateListOptimal(t *testing.T) {
	tcs := getTestCases()

	for _, tc := range tcs {
		list := tc["list"].([]int)
		k := tc["k"].(int)

		expected := tc["expected"].([]int)

		head := linkedlist.CreateFromArray(list)
		resHead := rotate.RotateRightOptimal(head, k)

		actual := resHead.GetArrayFromLinkedList()

		assert.Equal(t, expected, actual)
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"list": []int{1,2,3,4,5},
			"k": 2,
			"expected": []int{4,5,1,2,3},
		},
		"tc2": {
			"list": []int{0,1,2},
			"k": 4,
			"expected": []int{2,0,1},
		},
		"tc3": {
			"list": []int{},
			"k": 1,
			"expected": []int{},
		},
		"tc4": {
			"list": []int{1},
			"k": 1,
			"expected": []int{1},
		},
	}
}
