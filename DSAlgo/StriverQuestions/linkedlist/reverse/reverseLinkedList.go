package reverse

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
)

func ReverseListRecursive(head *linkedlist.ListNode) *linkedlist.ListNode {
	node := head
	return reverseList(node)
}

func reverseList(node *linkedlist.ListNode) *linkedlist.ListNode {
	if node.Next == nil {
		return node
	}

	head := reverseList(node.Next)
	q := node.Next
	q.Next = node
	node.Next = nil

	return head
}

func ReverseListIterative(head *linkedlist.ListNode) *linkedlist.ListNode {
	var prevNode *linkedlist.ListNode
	var currNode *linkedlist.ListNode
	var nextNode *linkedlist.ListNode

	currNode = head
	nextNode = head

	for nextNode != nil {
		nextNode = currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode
	}

	head = prevNode

	return head
}
