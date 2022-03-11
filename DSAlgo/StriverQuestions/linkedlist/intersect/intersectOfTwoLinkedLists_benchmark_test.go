package intersect_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/intersect"
	"testing"
)

func Benchmark_ForIntersectOfTwoLinkedLists(b *testing.B) {
	tcs := getTestCases()

	tc := tcs["tc1"]
	l1 := tc["l1"].([]int)
	l2 := tc["l2"].([]int)

	inter := tc["expected"].(int)
	head1, head2 := linkedlist.CreateIntersectedFromArrays(l1, l2, inter)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = intersect.GetIntersectionNode(head1, head2)
	}
	b.ReportAllocs()
}

func Benchmark_ForIntersectOfTwoLinkedListsByDifferenceOfLengths(b *testing.B) {
	tcs := getTestCases()

	tc := tcs["tc1"]
	l1 := tc["l1"].([]int)
	l2 := tc["l2"].([]int)

	inter := tc["expected"].(int)
	head1, head2 := linkedlist.CreateIntersectedFromArrays(l1, l2, inter)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = intersect.GetIntersectionNodeByDifferenceOfLength(head1, head2)
	}
	b.ReportAllocs()
}

func Benchmark_ForIntersectOfTwoLinkedListsByHashing(b *testing.B) {
	tcs := getTestCases()

	tc := tcs["tc1"]
	l1 := tc["l1"].([]int)
	l2 := tc["l2"].([]int)

	inter := tc["expected"].(int)
	head1, head2 := linkedlist.CreateIntersectedFromArrays(l1, l2, inter)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = intersect.GetIntersectionNodeByHashing(head1, head2)
	}
	b.ReportAllocs()
}
