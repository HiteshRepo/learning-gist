package subsetSums_test

import (
	"testing"
)

func Benchmark_SubsetSumsByPowerBit(b *testing.B) {
	tc := getTestCases()["tc1"]
	nums := tc["arr"].([]int)
	N := tc["N"].(int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = SubsetSumsByPowerBit(nums, N)
	}
	b.ReportAllocs()
}

func Benchmark_SubsetSumsByPowerSet(b *testing.B) {
	tc := getTestCases()["tc1"]
	nums := tc["arr"].([]int)
	N := tc["N"].(int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = SubsetSumsByPowerSet(nums, N)
	}
	b.ReportAllocs()
}

func Benchmark_SubsetSumsByRecursion(b *testing.B) {
	tc := getTestCases()["tc1"]
	nums := tc["arr"].([]int)
	N := tc["N"].(int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = SubsetSumsByRecursion(nums, N)
	}
	b.ReportAllocs()
}
