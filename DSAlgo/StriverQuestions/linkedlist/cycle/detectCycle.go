package cycle

import "github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"

func HasCycle(head *linkedlist.ListNode) bool {
	linkMap := make(map[*linkedlist.ListNode]struct{})
	tmp := head

	for tmp != nil {
		if _, ok := linkMap[tmp]; ok {
			return true
		}
		linkMap[tmp] = struct {}{}
		tmp = tmp.Next
	}

	return false
}

func HasCycleSlowAndFastPointers(head *linkedlist.ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
