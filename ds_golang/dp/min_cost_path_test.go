package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMinCostPath(t *testing.T) {
	matrix := getMatrix()
	expectedMinCost := 36
	actualMinCost := dp.GetMinCostPath(matrix)
	assert.Equal(t, expectedMinCost, actualMinCost)
}

func getMatrix() [][]int {
	return [][]int {
		{2,8,4,1,6,4,2},
		{6,0,9,5,3,5,8},
		{1,4,3,4,0,6,5},
		{6,4,7,2,4,6,1},
		{1,0,3,7,1,2,7},
		{1,5,3,2,3,0,9},
		{2,2,5,1,9,8,2},
	}
}
