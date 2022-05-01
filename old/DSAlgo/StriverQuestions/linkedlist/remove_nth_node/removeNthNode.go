package remove_nth_node

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
)

func RemoveNthFromEnd(head *linkedlist.ListNode, n int) *linkedlist.ListNode {
	start := &linkedlist.ListNode{Next: head}
	slow := start
	fast := start

	count := 0
	var length int
	for count <= n && fast != nil {
		fast = fast.Next
		count += 1
		if fast == nil {
			length = count
		}
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	if length > 0 {
		head = head.Next
	} else {
		slow.Next = slow.Next.Next
	}

	return head
}
