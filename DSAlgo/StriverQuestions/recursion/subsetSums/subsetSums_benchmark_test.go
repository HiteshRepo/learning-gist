package subsetSums_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/recursion/subsetSums"
	"testing"
)

func Benchmark_SubsetSumsByPowerBit(b *testing.B) {
	tc := getTestCases()["tc1"]
	nums := tc["arr"].([]int)
	N := tc["N"].(int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = subsetSums.SubsetSumsByPowerBit(nums, N)
	}
	b.ReportAllocs()
}

func Benchmark_SubsetSumsByPowerSet(b *testing.B) {
	tc := getTestCases()["tc1"]
	nums := tc["arr"].([]int)
	N := tc["N"].(int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = subsetSums.SubsetSumsByPowerSet(nums, N)
	}
	b.ReportAllocs()
}

func Benchmark_SubsetSumsByRecursion(b *testing.B) {
	tc := getTestCases()["tc1"]
	nums := tc["arr"].([]int)
	N := tc["N"].(int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = subsetSums.SubsetSumsByRecursion(nums, N)
	}
	b.ReportAllocs()
}
