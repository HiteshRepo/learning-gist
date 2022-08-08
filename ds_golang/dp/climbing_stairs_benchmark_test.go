package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"testing"
)

func BenchmarkGetCountOfPathsToClimbStairs(b *testing.B) {
	n := 10

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = dp.GetCountOfPathsToClimbStairs(n)
	}
	b.ReportAllocs()
}

func BenchmarkGetCountOfPathsToClimbStairsOptimized(b *testing.B) {
	n := 10

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = dp.GetCountOfPathsToClimbStairsMemoized(n)
	}
	b.ReportAllocs()
}

func BenchmarkGetCountOfPathsToClimbStairsTabulation(b *testing.B) {
	n := 10

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = dp.GetCountOfPathsToClimbStairsTabulation(n)
	}
	b.ReportAllocs()
}
