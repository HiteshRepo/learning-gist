package amazon_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/amazon"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatten(t *testing.T) {
	arr := [][]int{{1, 2}, {3, 4, 5, 6}, {7, 8, 9}}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := amazon.Flatten(arr)
	assert.Equal(t, want, got)
}
