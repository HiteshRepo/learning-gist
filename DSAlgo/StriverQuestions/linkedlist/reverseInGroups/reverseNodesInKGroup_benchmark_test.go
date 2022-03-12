package reverseInGroups_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/reverseInGroups"
	"testing"
)

func Benchmark_ReverseInGroups(b *testing.B) {
	tc := getTestCases()["tc1"]

	list := tc["list"].([]int)
	groupSize := tc["groupSize"].(int)
	head := linkedlist.CreateFromArray(list)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = reverseInGroups.ReverseKGroup(head, groupSize)
	}
	b.ReportAllocs()
}
