package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountBinaryStrings(t *testing.T) {
	length := 5
	expected := 21
	actual := dp.CountBinaryStrings(length)
	assert.Equal(t, expected, actual)
}
