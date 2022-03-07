package linkedlist

import (
	"fmt"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
)

func reverseListRecursive(head *ListNode) *ListNode {
	node := head
	return reverseList(node)
}

func reverseList(node *ListNode) *ListNode {
	if node.Next == nil {
		return node
	}

	head := reverseList(node.Next)
	q := node.Next
	q.Next = node
	node.Next = nil

	return head
}

func reverseListIterative(head *ListNode) *ListNode {
	var prevNode *ListNode
	var currNode *ListNode
	var nextNode *ListNode

	currNode = head
	nextNode = head

	for nextNode != nil {
		nextNode = currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode
	}

	head = prevNode

	return head
}

func RunTestsForReverseLinkedListIterative()  {
	tcs := getTestCases()

	for name, tc := range tcs {
		input := tc["input"].([]int)
		head := CreateFromArray(input)

		expected := tc["expected"].([]int)

		reversedHead :=  reverseListIterative(head)
		actual := reversedHead.GetArrayFromLinkedList()

		fmt.Printf("Is the solution correct for testcase %s : %v\n", name, utils.IntArrayEquals(expected, actual))
	}
}

func RunTestsForReverseLinkedListRecursive()  {
	tcs := getTestCases()

	for name, tc := range tcs {
		input := tc["input"].([]int)
		head := CreateFromArray(input)

		expected := tc["expected"].([]int)

		reversedHead :=  reverseListRecursive(head)
		actual := reversedHead.GetArrayFromLinkedList()

		fmt.Printf("Is the solution correct for testcase %s : %v\n", name, utils.IntArrayEquals(expected, actual))
	}
}

func getTestCases() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"input": []int{1,2,3,4,5},
			"expected": []int{5,4,3,2,1},
		},
		"tc2": {
			"input": []int{1,2},
			"expected": []int{2,1},
		},
	}
}
