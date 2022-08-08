package uniqueSubsets_test

import (
	"testing"
)

func Benchmark_UniqueSubsets(b *testing.B) {
	tc := getTestcases()["tc1"]

	nums := tc["nums"].([]int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = SubsetsWithDup(nums)
	}
	b.ReportAllocs()

}
