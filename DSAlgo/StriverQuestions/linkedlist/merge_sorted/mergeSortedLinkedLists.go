package merge_sorted

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
)

func MergeTwoListsInPlace(list1 *linkedlist.ListNode, list2 *linkedlist.ListNode) *linkedlist.ListNode {
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
		var tmp *linkedlist.ListNode
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

func MergeTwoLists(list1 *linkedlist.ListNode, list2 *linkedlist.ListNode) *linkedlist.ListNode {
	var head *linkedlist.ListNode
	prevNode := &linkedlist.ListNode{}

	for list1 != nil && list2 != nil {
		currNode := &linkedlist.ListNode{}
		if list1.Val > list2.Val {
			currNode.Val = list2.Val
			list2 = list2.Next
		} else if list2.Val > list1.Val {
			currNode.Val = list1.Val
			list1 = list1.Next
		} else if list1.Val == list2.Val {
			currNode2 := &linkedlist.ListNode{}
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
		currNode := &linkedlist.ListNode{}
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
		currNode := &linkedlist.ListNode{}
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
