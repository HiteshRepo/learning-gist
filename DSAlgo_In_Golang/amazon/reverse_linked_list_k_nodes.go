package amazon

import (
	linkedlist "github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/utils/linked_list"
)

// Given a linked list, write a function to reverse every k nodes (where k is an input to the function).
func ReverseKGroup(head *linkedlist.Node, k int) *linkedlist.Node {
	length := 0
	tmp := head

	for tmp != nil {
		length += 1
		tmp = tmp.Next()
	}

	tmp = linkedlist.GetNewNode(0)
	tmp.SetNext(head)

	pre := tmp
	var curr *linkedlist.Node
	var next *linkedlist.Node

	for length >= k {
		curr = pre.Next()
		next = curr.Next()

		for i:=1; i<k; i++ {
			curr.SetNext(next.Next())
			next.SetNext(pre.Next())
			pre.SetNext(next)
			next = curr.Next()
		}

		pre = curr
		length -= k
	}

	return tmp.Next()
}
