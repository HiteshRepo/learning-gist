package uniqueSubsets_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/recursion/uniqueSubsets"
	"testing"
)

func Benchmark_UniqueSubsets(b *testing.B) {
	tc := getTestcases()["tc1"]

	nums := tc["nums"].([]int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = uniqueSubsets.SubsetsWithDup(nums)
	}
	b.ReportAllocs()

}
