package cycle_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"testing"
)

func Benchmark_HasCycle(b *testing.B)  {
	tc := getTestcases()["tc1"]

	list := tc["list"].([]int)
	pos := tc["expected"].(int)

	head := linkedlist.CreateCycleFromArray(list, pos)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = HasCycle(head)
	}
	b.ReportAllocs()
}


func Benchmark_GetCycleHead(b *testing.B)  {
	tc := getTestcases()["tc1"]

	list := tc["list"].([]int)
	pos := tc["expected"].(int)

	head := linkedlist.CreateCycleFromArray(list, pos)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = GetCycleHead(head)
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
		_ = HasCycleSlowAndFastPointers(head)
	}
	b.ReportAllocs()
}

func Benchmark_GetCycleHeadSlowAndFastPointers(b *testing.B)  {
	tc := getTestcases()["tc1"]

	list := tc["list"].([]int)
	pos := tc["expected"].(int)

	head := linkedlist.CreateCycleFromArray(list, pos)

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = GetCycleHeadUsingSlowAndFastPointers(head)
	}
	b.ReportAllocs()
}
