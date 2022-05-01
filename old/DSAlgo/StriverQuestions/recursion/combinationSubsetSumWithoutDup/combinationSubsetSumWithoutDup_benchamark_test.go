package combinationSubsetSumWithoutDup_test

import (
	"testing"
)

func Benchmark_CombinationSum2(b *testing.B) {
	tc := getTestcases()["tc1"]
	candidates := tc["candidates"].([]int)
	target := tc["target"].(int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CombinationSum2(candidates, target)
	}
	b.ReportAllocs()
}
