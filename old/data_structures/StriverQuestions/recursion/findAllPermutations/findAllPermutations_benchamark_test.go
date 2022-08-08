package findAllPermutations_test

import (
	"testing"
)

func Benchmark_GetAllPermutations(b *testing.B) {
	tcs := getTestcases()

	for _, tc := range tcs {
		nums := tc["nums"].([]int)

		b.ResetTimer()
		for i:=0; i<b.N; i++ {
			_ = Permute(nums)
		}
		b.ReportAllocs()
	}
}
