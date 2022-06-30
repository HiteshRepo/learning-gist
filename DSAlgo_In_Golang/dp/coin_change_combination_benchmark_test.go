package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"testing"
)

func BenchmarkGetNumberOfCoinChangeCombinationPossible(b *testing.B) {
	tc := map[string]interface{}{
		"coins":    []int{2, 3, 5},
		"target":   7,
	}

	coins := tc["coins"].([]int)
	target := tc["target"].(int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = dp.GetNumberOfCoinChangeCombinationPossible(coins, target)
	}
	b.ReportAllocs()
}

func BenchmarkGetNumberOfCoinChangeCombinationPossibleByRecursion(b *testing.B) {
	tc := map[string]interface{}{
		"coins":    []int{2, 3, 5},
		"target":   7,
	}

	coins := tc["coins"].([]int)
	target := tc["target"].(int)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = dp.GetNumberOfCoinChangeCombinationPossibleByRecursion(coins, target)
	}
	b.ReportAllocs()
}
