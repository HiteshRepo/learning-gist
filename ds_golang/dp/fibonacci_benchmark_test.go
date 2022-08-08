package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"testing"
)

func BenchmarkGetNthFibonacciNumber(b *testing.B) {
	n := 35

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = dp.GetNthFibonacciNumber(n)
	}
	b.ReportAllocs()
}

func BenchmarkGetNthFibonacciNumberOptimized(b *testing.B) {
	n := 100

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = dp.GetNthFibonacciNumberOptimized(n)
	}
	b.ReportAllocs()
}
