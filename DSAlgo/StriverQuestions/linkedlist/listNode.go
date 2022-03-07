package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}

func CreateFromArray(arr []int) *ListNode {
	head := &ListNode{}

	var prev *ListNode
	for i, n := range arr {
		curr := &ListNode{
			Val:  n,
			Next: nil,
		}

		if i > 0 {
			prev.Next = curr
		}

		prev = curr
		if i == 0 {
			head = prev
		}
	}

	return head
}

func (l *ListNode) GetArrayFromLinkedList() []int {
	res := make([]int, 0)
	node := l
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	return res
}

func (l *ListNode) Display() {
	node := l
	for node != nil {
		fmt.Printf("%d->", node.Val)
		node = node.Next
	}
	fmt.Println()
}
