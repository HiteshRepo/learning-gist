package middle_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"testing"
)

func Benchmark_MiddleNode(b *testing.B) {
	tcs := getTestCasesForMiddleNode()

	for _, tc := range tcs {
		input := tc["input"].([]int)
		head := linkedlist.CreateFromArray(input)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ =  MiddleNode(head)
		}
		b.ReportAllocs()
	}
}


func Benchmark_MiddleNodeOptimized(b *testing.B)  {
	tcs := getTestCasesForMiddleNode()

	for _, tc := range tcs {
		input := tc["input"].([]int)
		head := linkedlist.CreateFromArray(input)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ =  MiddleNodeOptimized(head)
		}
		b.ReportAllocs()
	}
}
