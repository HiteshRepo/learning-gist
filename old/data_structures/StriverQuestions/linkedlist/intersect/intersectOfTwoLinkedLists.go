package intersect

import "github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"

func GetIntersectionNode(headA, headB *linkedlist.ListNode) *linkedlist.ListNode {
	tmpA := headA
	
	for tmpA != nil {
		tmpB := headB
		for tmpB != nil {
			if tmpA == tmpB {
				return tmpB
			}
			tmpB = tmpB.Next
		}
		tmpA = tmpA.Next
	}

	return nil
}

func GetIntersectionNodeByDifferenceOfLength(headA, headB *linkedlist.ListNode) *linkedlist.ListNode {
	len1 := 0
	len2 := 0

	tmp1 := headA
	tmp2 := headB

	for tmp1 != nil {
		len1 += 1
		tmp1 = tmp1.Next
	}

	for tmp2 != nil {
		len2 += 1
		tmp2 = tmp2.Next
	}

	tmp1 = headA
	tmp2 = headB
	diff := 0
	if len1 > len2 {
		diff = len1 - len2
		c := 0
		for c < diff {
			tmp1 = tmp1.Next
			c += 1
		}
	} else {
		diff = len2 - len1
		c := 0
		for c < diff {
			tmp2 = tmp2.Next
			c += 1
		}
	}


	for tmp1 != nil && tmp2 != nil {
		if tmp1 == tmp2 {
			return tmp2
		}

		tmp1 = tmp1.Next
		tmp2 = tmp2.Next
	}
	return nil
}

func GetIntersectionNodeByHashing(headA, headB *linkedlist.ListNode) *linkedlist.ListNode {
	tmpA := headA
	tmpB := headB

	llMap := map[*linkedlist.ListNode]struct{}{}

	for tmpA != nil {
		llMap[tmpA] = struct{}{}
		tmpA = tmpA.Next
	}

	for tmpB != nil {
		if _, ok := llMap[tmpB]; ok {
			return tmpB
		}
		tmpB = tmpB.Next
	}

	return nil
}