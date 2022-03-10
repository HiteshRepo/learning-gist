package merge_sorted_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/merge_sorted"
	"testing"
)

func Benchmark_MergeTwoListsInPlace(b *testing.B) {
	tcs := getTestCasesForMergeTwoSortedLinkedLists()

	for _, tc := range tcs {
		list1 := tc["list1"].([]int)
		list2 := tc["list2"].([]int)
		head1 := linkedlist.CreateFromArray(list1)
		head2 := linkedlist.CreateFromArray(list2)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = merge_sorted.MergeTwoListsInPlace(head1, head2)
		}
		b.ReportAllocs()
	}
}

func Benchmark_MergeTwoLists(b *testing.B) {
	tcs := getTestCasesForMergeTwoSortedLinkedLists()

	for _, tc := range tcs {
		list1 := tc["list1"].([]int)
		list2 := tc["list2"].([]int)
		head1 := linkedlist.CreateFromArray(list1)
		head2 := linkedlist.CreateFromArray(list2)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = merge_sorted.MergeTwoLists(head1, head2)
		}
		b.ReportAllocs()
	}
}
