package amazon_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/amazon"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindProductPair(t *testing.T) {
	nums := []int{7,9,4,8}
	product := 56
	want := []int{0,3}

	n1, n2 := amazon.FindProductPair(nums, product)
	got := []int{n1, n2}

	assert.Equal(t, want, got)
}
