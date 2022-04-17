package jobSequencing_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/greedy/jobSequencing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_JobSequencing(t *testing.T) {
	tcs := getTestcases()

	for _,tc := range tcs {
		jobs := make([]jobSequencing.Job, 0)

		jobsArr := tc["jobs"].([][]int)
		for _,job := range jobsArr {
			jobs = append(jobs, jobSequencing.Job{Id: job[0], Deadline: job[1], Profit: job[2]})
		}

		expected := tc["expected"].([]int)

		actual := jobSequencing.JobSequencing(jobs)
		assert.ElementsMatch(t, expected, actual)
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"jobs": [][]int{{1,4,20}, {2,1,10}, {3,1,40}, {4,1,30}},
			"expected": []int{2,60},
		},
		"tc2": {
			"jobs": [][]int{{1,2,100}, {2,1,19}, {3,2,27}, {4,1,25}, {5,1,15}},
			"expected": []int{2,127},
		},
	}
}
