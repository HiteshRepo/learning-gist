package amazon_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/amazon"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductArray(t *testing.T) {
	arr := []int{1,2,3,4,5}

	expected := []int{120, 60, 40, 30, 24}

	actual := amazon.ProductArray(arr)

	assert.Equal(t, expected, actual)
}
