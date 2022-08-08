package print_reverse_linked_list

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/utils/linked_list"
)

func PrintReverseLinkedList(head *linked_list.Node) []int {
	ans := make([]int, 0)
	node := head

	newHead := reverseLinkedList(node)
	node = newHead
	for node != nil {
		ans = append(ans, node.GetData())
		node = node.Next()
	}

	return ans
}

func reverseLinkedList(node *linked_list.Node) *linked_list.Node {
	if node.Next() == nil {
		return node
	}

	head := reverseLinkedList(node.Next())
	next := node.Next()
	next.SetNext(node)
	node.SetNext(nil)

	return head
}
