package trappingRainwater_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/array/trappingRainwater"
	"testing"
)

func Benchmark_TrapRainwater(b *testing.B) {
	tc := getTestcases()["tc1"]
	height := tc["height"].([]int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = trappingRainwater.TrapRainwater(height)
	}
	b.ReportAllocs()
}

func Benchmark_TrapRainwaterOptimal(b *testing.B) {
	tc := getTestcases()["tc1"]
	height := tc["height"].([]int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = trappingRainwater.TrapRainwaterOptimal(height)
	}
	b.ReportAllocs()
}
