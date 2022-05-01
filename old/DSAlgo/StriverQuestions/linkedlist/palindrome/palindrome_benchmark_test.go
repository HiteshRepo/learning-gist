package palindrome_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"testing"
)

func Benchmark_IsPalindrome(b *testing.B) {
	tc := getTestcases()["tc1"]

	list := tc["list"].([]int)

	head := linkedlist.CreateFromArray(list)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsPalindrome(head)
	}
	b.ReportAllocs()
}
