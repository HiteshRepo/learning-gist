package combinationSubsetSumWithoutDup_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/recursion/combinationSubsetSumWithoutDup"
	"testing"
)

func Benchmark_CombinationSum2(b *testing.B) {
	tc := getTestcases()["tc1"]
	candidates := tc["candidates"].([]int)
	target := tc["target"].(int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = combinationSubsetSumWithoutDup.CombinationSum2(candidates, target)
	}
	b.ReportAllocs()
}
