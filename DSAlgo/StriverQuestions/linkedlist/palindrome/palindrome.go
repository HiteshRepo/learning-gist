package palindrome

import "github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"

func IsPalindrome(head *linkedlist.ListNode) bool {
	arr := make([]int, 0)
	recursiveTraverse(head, &arr)

	i := 0
	j := len(arr)-1

	for i < j {
		if arr[i] != arr[j] {
			return false
		}
		i += 1
		j -= 1
	}

	return true
}

func IsPalindromeOptimized(head *linkedlist.ListNode) bool {
	tmp := head
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = reverseList(slow)

	for slow != nil {
		if slow.Val != tmp.Val {
			return false
		}
		slow = slow.Next
		tmp = tmp.Next
	}

	return true
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

func recursiveTraverse(head *linkedlist.ListNode, arr *[]int) {
	if head == nil {
		return
	}

	*arr = append(*arr, head.Val)
	recursiveTraverse(head.Next, arr)
}
