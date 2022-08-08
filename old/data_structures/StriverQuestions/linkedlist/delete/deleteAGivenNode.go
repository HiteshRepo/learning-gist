package delete

import "github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"

func DeleteNode(node *linkedlist.ListNode) {
	//node = node.Next
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}