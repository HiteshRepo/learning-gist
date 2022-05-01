package removeDuplicates_test

import (
	"testing"
)

func Benchmark_RemoveDuplicates(b *testing.B) {
	tc := getTestCases()["tc1"]
	nums := tc["nums"].([]int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = RemoveDuplicates(nums)
	}
	b.ReportAllocs()
}

func Benchmark_RemoveDuplicatesOptimized(b *testing.B) {
	tc := getTestCases()["tc1"]
	nums := tc["nums"].([]int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = RemoveDuplicatesOptimized(nums)
	}
	b.ReportAllocs()
}
