package findAllPermutations_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/recursion/findAllPermutations"
	"testing"
)

func Benchmark_GetAllPermutations(b *testing.B) {
	tcs := getTestcases()

	for _, tc := range tcs {
		nums := tc["nums"].([]int)

		b.ResetTimer()
		for i:=0; i<b.N; i++ {
			_ = findAllPermutations.Permute(nums)
		}
		b.ReportAllocs()
	}
}
