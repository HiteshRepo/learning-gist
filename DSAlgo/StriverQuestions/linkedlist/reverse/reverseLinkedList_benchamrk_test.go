package reverse_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/reverse"
	"testing"
)

func Benchmark_ReverseLinkedListIterative(b *testing.B)  {
	tcs := getTestCases()

	for _, tc := range tcs {
		input := tc["input"].([]int)
		head := linkedlist.CreateFromArray(input)

		b.ResetTimer()
		for i:=0; i< b.N; i++ {
			_ = reverse.ReverseListIterative(head)
		}
		b.ReportAllocs()

	}
}

func Benchmark_ReverseLinkedListRecursive(b *testing.B)  {
	tcs := getTestCases()

	for _, tc := range tcs {
		input := tc["input"].([]int)
		head := linkedlist.CreateFromArray(input)

		b.ResetTimer()
		for i:=0; i< b.N; i++ {
			_ = reverse.ReverseListRecursive(head)
		}
		b.ReportAllocs()
	}
}
