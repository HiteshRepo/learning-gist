package cycle_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/cycle"
	"testing"
)

func Benchmark_HasCycle(b *testing.B)  {
	tc := getTestcases()["tc1"]

	list := tc["list"].([]int)
	pos := tc["expected"].(int)

	head := linkedlist.CreateCycleFromArray(list, pos)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = cycle.HasCycle(head)
	}
	b.ReportAllocs()
}

func Benchmark_HasCycleSlowAndFastPointers(b *testing.B)  {
	tc := getTestcases()["tc1"]

	list := tc["list"].([]int)
	pos := tc["expected"].(int)

	head := linkedlist.CreateCycleFromArray(list, pos)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = cycle.HasCycleSlowAndFastPointers(head)
	}
	b.ReportAllocs()
}
