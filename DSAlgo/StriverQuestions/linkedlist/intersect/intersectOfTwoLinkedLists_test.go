package intersect_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/intersect"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ForIntersectOfTwoLinkedLists(t *testing.T) {
	tcs := getTestCases()

	for _,tc := range tcs {
		l1 := tc["l1"].([]int)
		l2 := tc["l2"].([]int)

		if tc["expected"] != nil {
			expected := tc["expected"].(int)
			head1, head2 := linkedlist.CreateIntersectedFromArrays(l1, l2, expected)
			intersectNode := intersect.GetIntersectionNode(head1, head2)
			assert.Equal(t, expected, intersectNode.Val)
		} else {
			head1 := linkedlist.CreateFromArray(l1)
			head2 := linkedlist.CreateFromArray(l2)
			intersectNode := intersect.GetIntersectionNode(head1, head2)
			assert.Nil(t, intersectNode)
		}
	}
}

func Test_ForIntersectOfTwoLinkedListsByDifferenceOfLength(t *testing.T) {
	tcs := getTestCases()

	for _,tc := range tcs {
		l1 := tc["l1"].([]int)
		l2 := tc["l2"].([]int)

		if tc["expected"] != nil {
			expected := tc["expected"].(int)
			head1, head2 := linkedlist.CreateIntersectedFromArrays(l1, l2, expected)
			intersectNode := intersect.GetIntersectionNodeByDifferenceOfLength(head1, head2)
			assert.Equal(t, expected, intersectNode.Val)
		} else {
			head1 := linkedlist.CreateFromArray(l1)
			head2 := linkedlist.CreateFromArray(l2)
			intersectNode := intersect.GetIntersectionNodeByDifferenceOfLength(head1, head2)
			assert.Nil(t, intersectNode)
		}
	}
}

func Test_ForIntersectOfTwoLinkedListsByHashing(t *testing.T) {
	tcs := getTestCases()

	for _,tc := range tcs {
		l1 := tc["l1"].([]int)
		l2 := tc["l2"].([]int)

		if tc["expected"] != nil {
			expected := tc["expected"].(int)
			head1, head2 := linkedlist.CreateIntersectedFromArrays(l1, l2, expected)
			intersectNode := intersect.GetIntersectionNodeByHashing(head1, head2)
			assert.Equal(t, expected, intersectNode.Val)
		} else {
			head1 := linkedlist.CreateFromArray(l1)
			head2 := linkedlist.CreateFromArray(l2)
			intersectNode := intersect.GetIntersectionNodeByHashing(head1, head2)
			assert.Nil(t, intersectNode)
		}
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"l1": []int{4,1,8,4,5},
			"l2": []int{5,6,1,8,4,5},
			"expected": 8,
		},
		"tc2": {
			"l1": []int{1,9,1,2,4},
			"l2": []int{3,2,4},
			"expected": 2,
		},
		"tc3": {
			"l1": []int{2,6,4},
			"l2": []int{1,5},
			"expected": nil,
		},
	}
}
