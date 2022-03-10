package middle

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
)

func MiddleNodeOptimized(head *linkedlist.ListNode) *linkedlist.ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func MiddleNode(head *linkedlist.ListNode) *linkedlist.ListNode {
	tempNode := head

	countNodes := 0
	for tempNode != nil {
		countNodes += 1
		tempNode = tempNode.Next
	}

	mid := countNodes/2
	if countNodes % 2 == 1 {
		mid += 1
		countNodes = 1
	} else {
		countNodes = 0
	}

	tempNode = head
	for countNodes < mid {
		countNodes += 1
		tempNode = tempNode.Next
	}

	return tempNode
}


