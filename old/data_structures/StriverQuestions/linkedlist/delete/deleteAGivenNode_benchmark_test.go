package delete_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"testing"
)

func Benchmark_DeleteGivenNode(b *testing.B) {
	tcs := getTestCases()

	for _, tc := range tcs {
		l1 := tc["l1"].([]int)
		node := tc["node"].(int)

		head := linkedlist.CreateFromArray(l1)

		var nodeToDelete *linkedlist.ListNode
		var tmp = head
		for tmp != nil {
			if tmp.Val == node {
				nodeToDelete = tmp
				break
			}
			tmp = tmp.Next
		}

		DeleteNode(nodeToDelete)

		b.ResetTimer()
		for i:=0; i<b.N; i++ {
			_ = head.GetArrayFromLinkedList()
		}
		b.ReportAllocs()
	}
}
