package amazon_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/amazon"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSquareRoot(t *testing.T) {
	//num := 4
	num := 5
	//want := float64(2)
	want := 2.23606
	got:= amazon.FindSquareRoot(num)
	assert.Equal(t, want, got)
}
