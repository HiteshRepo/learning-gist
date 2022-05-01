package print_linked_list

import "github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/utils/linked_list"

func PrintLinkedList(head *linked_list.Node) []int {
	node := head
	ans := make([]int, 0)

	for node != nil {
		ans = append(ans, node.GetData())
		node = node.Next()
	}

	return ans
}
