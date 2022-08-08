package clone_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"testing"
)

func Benchmark_Clone(b *testing.B) {
	tc := getTestCases()["tc1"]

	list := tc["list"].([][]*int)
	node := linkedlist.CreateNodeFromList(list)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = CopyRandomList(node)
	}
	b.ReportAllocs()
}

func Benchmark_CloneOptimized(b *testing.B) {
	tc := getTestCases()["tc1"]

	list := tc["list"].([][]*int)
	node := linkedlist.CreateNodeFromList(list)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = CopyRandomListOptimized(node)
	}
	b.ReportAllocs()
}
