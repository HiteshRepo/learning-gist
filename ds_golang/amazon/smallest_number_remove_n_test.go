package amazon_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/amazon"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmallestNumberByRemovingN(t *testing.T) {
	//digits := "4325043"
	//n := 3

	//digits := "765028321"
	//n := 5

	digits := "121198"
	n := 2

	//want := "2043"
	//want := "0221"
	want := "1118"

	got := amazon.SmallestNumberByRemovingN(digits, n)

	assert.Equal(t, want, got)
}
