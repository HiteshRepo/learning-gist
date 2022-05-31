package Tests

import (
	"testing"
)

func BenchmarkFindS1PermutationInS2(b *testing.B) {
	tcs := getTestCasesForFindS1PermutationInS2()

	for _,tc := range tcs {
		s1 := tc["s1"].(string)
		s2 := tc["s2"].(string)

		b.ResetTimer()
		for i:=0; i<b.N; i++ {
			_ = FindS1PermutationInS2(s1, s2)
		}
		b.ReportAllocs()
	}
}

func BenchmarkFindS1PermutationInS2Optimal(b *testing.B) {
	tcs := getTestCasesForFindS1PermutationInS2()

	for _,tc := range tcs {
		s1 := tc["s1"].(string)
		s2 := tc["s2"].(string)

		b.ResetTimer()
		for i:=0; i<b.N; i++ {
			_ = FindS1PermutationInS2Optimal(s1, s2)
		}
		b.ReportAllocs()
	}
}
