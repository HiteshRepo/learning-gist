package kthPermutation_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/recursion/kthPermutation"
	"testing"
)

func Benchmark_KthPermutation(b *testing.B) {
	tc := getTestCases()["tc1"]
	n := tc["n"].(int)
	k := tc["k"].(int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = kthPermutation.GetPermutation(n, k)
	}
	b.ReportAllocs()
}
