package kthPermutation_test

import (
	"testing"
)

func Benchmark_KthPermutation(b *testing.B) {
	tc := getTestCases()["tc1"]
	n := tc["n"].(int)
	k := tc["k"].(int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetPermutation(n, k)
	}
	b.ReportAllocs()
}
