package add_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/add"
	"testing"
)

func Benchmark_AddTowNumbers(b *testing.B) {
	tcs := getTestCases()

	for _, tc := range tcs {
		l1 := tc["l1"].([]int)
		l2 := tc["l2"].([]int)

		head1 := linkedlist.CreateFromArray(l1)
		head2 := linkedlist.CreateFromArray(l2)

		b.ResetTimer()
		for i:=0; i<b.N; i++ {
			_ = add.AddTwoNumbers(head1, head2)
		}
		b.ReportAllocs()
	}
}
