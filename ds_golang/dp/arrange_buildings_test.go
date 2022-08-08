package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrangeBuildings(t *testing.T) {
	roads := 1
	blocks := 5
	expected := 441
	actual := dp.ArrangeBuildings(blocks, roads)
	assert.Equal(t, expected, actual)
}
