package identical_twins_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/array/identical_twins"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetIdenticalTwinsCount(t *testing.T) {
	tcs := getTestcases()

	for _,tc := range tcs {
		arr := tc["arr"].([]int)
		expected := tc["expected"].(int)

		actual := identical_twins.GetIdenticalTwinsCount(arr)
		assert.Equal(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"arr": []int{1, 2, 3, 2, 1},
			"expected": 2,
		},
		"tc2": {
			"arr": []int{1, 2, 2, 3, 2, 1},
			"expected": 4,
		},
		"tc3": {
			"arr": []int{1, 1, 1, 1},
			"expected": 6,
		},
	}
}
