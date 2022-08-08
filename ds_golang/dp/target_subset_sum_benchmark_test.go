package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"testing"
)

func BenchmarkIsTargetSumPossible(b *testing.B) {
	tc := map[string]interface{}{
		"candidates": []int{10, 1, 2, 7, 6, 1, 5},
		"target":     8,
		"expected":   [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
	}
	candidates := tc["candidates"].([]int)
	target := tc["target"].(int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dp.IsTargetSumPossible(candidates, target)
	}
	b.ReportAllocs()
}

func BenchmarkIsTargetSumPossibleOptimized(b *testing.B) {
	tc := map[string]interface{}{
		"candidates": []int{10, 1, 2, 7, 6, 1, 5},
		"target":     8,
		"expected":   [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
	}
	candidates := tc["candidates"].([]int)
	target := tc["target"].(int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dp.IsTargetSumPossibleOptimized(candidates, target)
	}
	b.ReportAllocs()
}
