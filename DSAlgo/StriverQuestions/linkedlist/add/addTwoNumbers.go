package add

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
)

func AddTwoNumbers(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) *linkedlist.ListNode {
	sum := 0
	carry := 0

	tmp := &linkedlist.ListNode{}
	head := tmp

	for l1 != nil || l2 != nil || carry > 0 {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if carry > 0 {
			sum += carry
			carry = 0
		}

		newNode := &linkedlist.ListNode{}
		if sum >= 10 {
			newNode.Val = sum%10
			carry += sum/10
		} else {
			newNode.Val = sum
		}

		tmp.Next = newNode
		tmp = newNode
	}

	return head.Next
}
