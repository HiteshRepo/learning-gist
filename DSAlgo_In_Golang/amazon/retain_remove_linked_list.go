package amazon

import linkedlist "github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/utils/linked_list"

// Given a linked list and two integers M and N. Traverse the linked list such that you retain M nodes then delete next N nodes, continue the same till the end of the linked list. For example, an input of M = 2, N = 2 Linked List: 1->2->3->4->5->6->7->8 should return Linked List: 1->2->5->6
func RetainRemove(head *linkedlist.Node, m,n int) {
	tmp := head

	countM := 0
	countN := 0
	var mNode, nNode *linkedlist.Node
	for tmp != nil {

		for countM < m && tmp != nil {
			mNode = tmp
			countM += 1
			tmp = tmp.Next()
		}

		for countN < n && tmp != nil {
			nNode = tmp
			countN += 1
			tmp = tmp.Next()
		}

		if countM == m && countN == n {
			mNode.SetNext(nNode.Next())
			countM = 0
			countN = 0
		}
	}

}
