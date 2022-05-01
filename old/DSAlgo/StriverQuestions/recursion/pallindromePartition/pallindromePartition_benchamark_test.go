package pallindromePartition_test

import (
	"testing"
)

func Benchmark_PalindromePartition(b *testing.B) {
	tcs := getTestCases()

	for _, tc := range tcs {
		s := tc["s"].(string)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Partition(s)
		}
		b.ReportAllocs()
	}
}
