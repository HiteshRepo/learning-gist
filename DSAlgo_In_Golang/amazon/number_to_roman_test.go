package amazon_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/amazon"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberToRoman(t *testing.T) {
	num := 3949
	want := "MMMCMXLIX"
	got := amazon.NumberToRoman(num)
	assert.Equal(t, want, got)
}

func TestRomanToNumber(t *testing.T) {
	roman := "MMMCMXLIX"
	want := 3949

	//roman := "VI"
	//want := 6

	//roman := "IV"
	//want := 4

	//roman := "XIX"
	//want := 19

	got := amazon.RomanToNumber(roman)
	assert.Equal(t, want, got)
}
