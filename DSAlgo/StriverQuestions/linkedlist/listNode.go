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
	var head *ListNode

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

func CreateIntersectedFromArrays(arr1, arr2 []int, intersect int) (*ListNode, *ListNode) {
	var head1 *ListNode
	var head2 *ListNode

	var tmp1 *ListNode
	var tmp2 *ListNode

	tillNow1 := make([]int, 0)
	tillNow2 := make([]int, 0)
	for _, n := range arr1 {
		if n != intersect {
			tillNow1 = append(tillNow1, n)
			continue
		}

		for _, n2 := range arr2 {
			if n2 != intersect {
				tillNow2 = append(tillNow2, n2)
				continue
			}
			break
		}

		if head1 == nil {
			tmp1 = CreateFromArray(tillNow1)
			tmp2 = CreateFromArray(tillNow2)

			head1 = tmp1
			head2 = tmp2
		}

		for tmp1.Next != nil {
			tmp1 = tmp1.Next
		}

		for tmp2.Next != nil {
			tmp2 = tmp2.Next
		}

		if head1 != nil {
			curr := &ListNode{
				Val:  n,
				Next: nil,
			}

			tmp1.Next = curr
			tmp2.Next = curr

			tmp1 = tmp1.Next
			tmp2 = tmp2.Next
		}
	}

	return head1, head2
}

func CreateCycleFromArray(arr []int, pos int) *ListNode {
	var head *ListNode
	var posNode *ListNode

	var prev *ListNode
	for i, n := range arr {
		curr := &ListNode{
			Val:  n,
			Next: nil,
		}

		if i == pos {
			posNode = curr
		}

		if i > 0 {
			prev.Next = curr
		}

		prev = curr
		if i == 0 {
			head = prev
		}
	}

	prev.Next = posNode

	return head
}
