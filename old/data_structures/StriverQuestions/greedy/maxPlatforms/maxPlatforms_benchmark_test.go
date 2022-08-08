package maxPlatforms_test

import (
	"testing"
)

func Benchmark_MaxPlatforms(b *testing.B) {
	tc := getTestcases()["tc1"]

	n := tc["n"].(int)
	arr := tc["arr"].([]int)
	dep := tc["dep"].([]int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = FindPlatforms(arr, dep, n)
	}
	b.ReportAllocs()
}

func Benchmark_MaxPlatformsOptimal(b *testing.B) {
	tc := getTestcases()["tc1"]

	n := tc["n"].(int)
	arr := tc["arr"].([]int)
	dep := tc["dep"].([]int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = FindPlatformsOptimal(arr, dep, n)
	}
	b.ReportAllocs()
}
