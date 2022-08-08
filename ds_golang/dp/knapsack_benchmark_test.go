package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"testing"
)

func BenchmarkGetMaximumWeightForGivenCapacity(b *testing.B) {
	tc := map[string]interface{}{
		"capacity": 7,
		"weights":  []int{2, 5, 1, 3, 4},
		"items":    []int{15, 14, 10, 45, 30},
		"expected": 75,
	}

	capacity := tc["capacity"].(int)
	weights := tc["weights"].([]int)
	items := tc["items"].([]int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dp.GetMaximumWeightForGivenCapacity(capacity, weights, items)
	}
	b.ReportAllocs()
}

func BenchmarkGetMaximumWeightForGivenCapacityByRecursion(b *testing.B) {
	tc := map[string]interface{}{
		"capacity": 7,
		"weights":  []int{2, 5, 1, 3, 4},
		"items":    []int{15, 14, 10, 45, 30},
		"expected": 75,
	}

	capacity := tc["capacity"].(int)
	weights := tc["weights"].([]int)
	items := tc["items"].([]int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dp.GetMaximumWeightForGivenCapacityByRecursion(capacity, weights, items)
	}
	b.ReportAllocs()
}
