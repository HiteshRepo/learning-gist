package Tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeStringFreq(t *testing.T) {
	tcs := map[string]map[string]string {
		"tc1": {
			"input": "abzx",
			"expected": "1226#24#",
		},
		"tc2": {
			"input": "aabccc",
			"expected": "1(2)23(3)",
		},
		"tc3": {
			"input": "bajj",
			"expected": "2110#(2)",
		},
		"tc4": {
			"input": "wwxyzwww",
			"expected": "23#(2)24#25#26#23#(3)",
		},
	}

	for _,tc := range tcs {
		input := tc["input"]
		expected := tc["expected"]
		actual := DecodeStringFreq(input)
		assert.Equal(t, expected, actual)
	}
}
