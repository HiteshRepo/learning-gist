package rotate_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"testing"
)

func Benchmark_RotateList(b *testing.B) {
	tc := getTestCases()["tc1"]

	list := tc["list"].([]int)
	k := tc["k"].(int)

	head := linkedlist.CreateFromArray(list)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = RotateRight(head, k)
	}
	b.ReportAllocs()
}
