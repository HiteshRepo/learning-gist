package combinationSubsetSum_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/recursion/combinationSubsetSum"
	"testing"
)

func Benchmark_CombinationSubsetSum(b *testing.B) {
	tc := getTestCases()["tc1"]
	candidates := tc["candidates"].([]int)
	target := tc["target"].(int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = combinationSubsetSum.CombinationSum(candidates, target)
	}
	b.ReportAllocs()
}
