package linkedlist

import (
	"fmt"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
)

func middleNodeOptimized(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func middleNode(head *ListNode) *ListNode {
	tempNode := head

	countNodes := 0
	for tempNode != nil {
		countNodes += 1
		tempNode = tempNode.Next
	}

	mid := countNodes/2
	if countNodes % 2 == 1 {
		mid += 1
		countNodes = 1
	} else {
		countNodes = 0
	}

	tempNode = head
	for countNodes < mid {
		countNodes += 1
		tempNode = tempNode.Next
	}

	return tempNode
}

func RunTestsForMiddleOfLinkedList()  {
	tcs := getTestCasesForMiddleNode()

	for name, tc := range tcs {
		input := tc["input"].([]int)
		head := CreateFromArray(input)

		expected := tc["expected"].([]int)

		middleNode :=  middleNode(head)
		actual := middleNode.GetArrayFromLinkedList()

		fmt.Printf("Is the solution correct for testcase %s : %v\n", name, utils.IntArrayEquals(expected, actual))
	}
}

func RunTestsForMiddleOfLinkedListOptimized()  {
	tcs := getTestCasesForMiddleNode()

	for name, tc := range tcs {
		input := tc["input"].([]int)
		head := CreateFromArray(input)

		expected := tc["expected"].([]int)

		middleNode :=  middleNodeOptimized(head)
		actual := middleNode.GetArrayFromLinkedList()

		fmt.Printf("Is the solution correct for testcase %s : %v\n", name, utils.IntArrayEquals(expected, actual))
	}
}

func getTestCasesForMiddleNode() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"input": []int{1,2,3,4,5},
			"expected": []int{3,4,5},
		},
		"tc2": {
			"input": []int{1,2,3,4,5,6},
			"expected": []int{4,5,6},
		},
	}
}
