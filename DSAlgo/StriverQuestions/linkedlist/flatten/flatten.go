package flatten

import "github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"

func Flatten(head *linkedlist.DoubleListNode) *linkedlist.DoubleListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head.Next = Flatten(head.Next)

	root := merge(head, head.Next)
	return root
}

func merge(l1, l2 *linkedlist.DoubleListNode) *linkedlist.DoubleListNode {
	tmp := &linkedlist.DoubleListNode{}
	root := tmp

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			tmp.Bottom = l2
			l2 = l2.Bottom
		} else {
			tmp.Bottom = l1
			l1 = l1.Bottom
		}

		tmp = tmp.Bottom
	}

	if l1 != nil {
		tmp.Bottom = l1
	}

	if l2 != nil {
		tmp.Bottom = l2
	}

	return root.Bottom
}
