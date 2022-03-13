package rotate

import "github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"

func RotateRight(head *linkedlist.ListNode, k int) *linkedlist.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	var last *linkedlist.ListNode
	var secondLast *linkedlist.ListNode

	count := 0
	for count < k {
		tmp := first
		for tmp.Next != nil {
			if tmp.Next.Next == nil {
				secondLast = tmp
			}
			tmp = tmp.Next
		}
		last = tmp

		last.Next = first
		secondLast.Next = nil
		first = last

		count += 1
	}

	return first
}

func RotateRightOptimal(head *linkedlist.ListNode, k int) *linkedlist.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head
	length := 0

	for tmp.Next != nil {
		tmp = tmp.Next
		length += 1
	}
	tmp.Next = head
	length += 1

	k = k % length
	targetNode := length - k
	var node *linkedlist.ListNode

	count := 0
	tmp = head
	for count < targetNode {
		node = tmp
		tmp = tmp.Next
		count += 1
	}
	node.Next = nil

	return tmp
}
