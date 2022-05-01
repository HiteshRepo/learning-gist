package reverseInGroups

import "github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"

func ReverseKGroup(head *linkedlist.ListNode, k int) *linkedlist.ListNode {
	tmp := head
	length := 0

	for tmp != nil {
		length += 1
		tmp = tmp.Next
	}

	tmp = &linkedlist.ListNode{Next: head}
	var pre = tmp
	var curr *linkedlist.ListNode
	var next *linkedlist.ListNode
	for length >= k {
		curr = pre.Next
		next = curr.Next
		for i:=1; i<k; i++ {
			curr.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
			next = curr.Next
		}
		pre = curr
		length -= k
	}

	return tmp.Next
}
