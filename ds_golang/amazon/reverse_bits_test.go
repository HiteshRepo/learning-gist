package amazon_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/amazon"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseBits(t *testing.T) {
	num := int64(3)

	want := int64(3221225472)
	got := amazon.ReverseBits(num)

	assert.Equal(t, want, got)
}
