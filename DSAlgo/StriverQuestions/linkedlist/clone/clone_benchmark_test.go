package clone_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/clone"
	"testing"
)

func Benchmark_Clone(b *testing.B) {
	tc := getTestCases()["tc1"]

	list := tc["list"].([][]*int)
	node := linkedlist.CreateNodeFromList(list)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = clone.CopyRandomList(node)
	}
	b.ReportAllocs()
}

func Benchmark_CloneOptimized(b *testing.B) {
	tc := getTestCases()["tc1"]

	list := tc["list"].([][]*int)
	node := linkedlist.CreateNodeFromList(list)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = clone.CopyRandomListOptimized(node)
	}
	b.ReportAllocs()
}
