package nMeetings_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/greedy/nMeetings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_nMeetings(t *testing.T) {
	tcs := getTestcases()

	for _,tc := range tcs {
		n := tc["n"].(int)
		start := tc["start"].([]int)
		end := tc["end"].([]int)

		expected := tc["expected"].(int)

		actual := nMeetings.MaxMeetings(start, end, n)

		assert.Equal(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"n": 6,
			"start": []int{1,3,0,5,8,5},
			"end": []int{2,4,6,7,9,9},
			"expected": 4,
		},
		"tc2": {
			"n": 3,
			"start": []int{10,12,20},
			"end": []int{20,25,30},
			"expected": 1,
		},
	}
}
