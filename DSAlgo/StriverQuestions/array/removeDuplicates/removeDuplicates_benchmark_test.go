package removeDuplicates_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/array/removeDuplicates"
	"testing"
)

func Benchmark_RemoveDuplicates(b *testing.B) {
	tc := getTestCases()["tc1"]
	nums := tc["nums"].([]int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = removeDuplicates.RemoveDuplicates(nums)
	}
	b.ReportAllocs()
}

func Benchmark_RemoveDuplicatesOptimized(b *testing.B) {
	tc := getTestCases()["tc1"]
	nums := tc["nums"].([]int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = removeDuplicates.RemoveDuplicatesOptimized(nums)
	}
	b.ReportAllocs()
}
