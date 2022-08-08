package trappingRainwater_test

import (
	"testing"
)

func Benchmark_TrapRainwater(b *testing.B) {
	tc := getTestcases()["tc1"]
	height := tc["height"].([]int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = TrapRainwater(height)
	}
	b.ReportAllocs()
}

func Benchmark_TrapRainwaterOptimal(b *testing.B) {
	tc := getTestcases()["tc1"]
	height := tc["height"].([]int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = TrapRainwaterOptimal(height)
	}
	b.ReportAllocs()
}
