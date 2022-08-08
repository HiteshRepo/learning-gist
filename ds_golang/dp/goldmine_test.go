package dp_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMaxGoldPath(t *testing.T) {
	matrix := getGoldMineMatrix()
	expectedMinCost := 33
	actualMinCost := dp.GetMaxGoldPath(matrix)
	assert.Equal(t, expectedMinCost, actualMinCost)
}


func getGoldMineMatrix() [][]int {
	return [][]int {
		{0,1,4,2,8,2},
		{4,3,6,5,0,4},
		{1,2,4,1,4,6},
		{2,0,7,3,2,2},
		{3,1,5,9,2,4},
		{2,7,0,8,5,1},
	}
}