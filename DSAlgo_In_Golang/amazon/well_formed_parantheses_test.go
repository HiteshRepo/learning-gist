package amazon_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/amazon"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWellFormedParantheses(t *testing.T) {
	//brackets := "((()()"
	brackets := "()(()))))"
	//want := 4
	want := 6
	got := amazon.WellFormedParantheses(brackets)
	assert.Equal(t, want, got)
}
