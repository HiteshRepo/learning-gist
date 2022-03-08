package linkedlist

import (
	"fmt"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
)

func mergeTwoListsInPlace(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		tmp := list1
		list1 = list2
		list2 = tmp
	}

	head := list1

	for list1 != nil && list2 != nil {
		var tmp *ListNode
		for list1 != nil && list1.Val <= list2.Val {
			tmp = list1
			list1 = list1.Next
		}
		tmp.Next = list2

		temp := list1
		list1 = list2
		list2 = temp
	}

	return head
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	prevNode := &ListNode{}

	for list1 != nil && list2 != nil {
		currNode := &ListNode{}
		if list1.Val > list2.Val {
			currNode.Val = list2.Val
			list2 = list2.Next
		} else if list2.Val > list1.Val {
			currNode.Val = list1.Val
			list1 = list1.Next
		} else if list1.Val == list2.Val {
			currNode2 := &ListNode{}
			currNode.Val = list2.Val
			currNode2.Val = list1.Val
			currNode.Next = currNode2
			list1 = list1.Next
			list2 = list2.Next
		}

		if head == nil {
			prevNode = currNode
			head = prevNode
		} else {
			prevNode.Next = currNode
		}

		for prevNode.Next != nil {
			prevNode = prevNode.Next
		}
	}

	for list1 != nil {
		currNode := &ListNode{}
		currNode.Val = list1.Val
		if head == nil {
			prevNode = currNode
			head = prevNode
		} else {
			prevNode.Next = currNode
		}

		for prevNode.Next != nil {
			prevNode = prevNode.Next
		}

		list1 = list1.Next
	}

	for list2 != nil {
		currNode := &ListNode{}
		currNode.Val = list2.Val
		if head == nil {
			prevNode = currNode
			head = prevNode
		} else {
			prevNode.Next = currNode
		}

		for prevNode.Next != nil {
			prevNode = prevNode.Next
		}

		list2 = list2.Next
	}

	return head
}

func RunTestsForMergeTwoSortedLinkedLists()  {
	tcs := getTestCasesForMergeTwoSortedLinkedLists()

	for name, tc := range tcs {
		list1 := tc["list1"].([]int)
		list2 := tc["list2"].([]int)
		head1 := CreateFromArray(list1)
		head2 := CreateFromArray(list2)

		expected := tc["expected"].([]int)

		mergedHead :=  mergeTwoLists(head1, head2)
		actual := mergedHead.GetArrayFromLinkedList()

		fmt.Printf("Is the solution correct for testcase %s : %v\n", name, utils.IntArrayEquals(expected, actual))
	}
}

func RunTestsForMergeTwoSortedLinkedListsInPlace()  {
	tcs := getTestCasesForMergeTwoSortedLinkedLists()

	for name, tc := range tcs {
		list1 := tc["list1"].([]int)
		list2 := tc["list2"].([]int)
		head1 := CreateFromArray(list1)
		head2 := CreateFromArray(list2)

		expected := tc["expected"].([]int)

		mergedHead :=  mergeTwoListsInPlace(head1, head2)
		actual := mergedHead.GetArrayFromLinkedList()

		fmt.Printf("Is the solution correct for testcase %s : %v\n", name, utils.IntArrayEquals(expected, actual))
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
