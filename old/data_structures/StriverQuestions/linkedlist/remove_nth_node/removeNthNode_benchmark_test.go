package remove_nth_node_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"testing"
)

func Benchmark_RemoveNthNode(b *testing.B) {
	tcs := getTestCasesForRemoveNthNodeFromEnd()
	for _,tc := range tcs {
		input := tc["input"].([]int)
		pos := tc["pos"].(int)

		head := linkedlist.CreateFromArray(input)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = RemoveNthFromEnd(head, pos)
		}
		b.ReportAllocs()
	}
}