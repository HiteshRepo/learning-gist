package cycle_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist/cycle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_HasCycle(t *testing.T)  {
	tcs := getTestcases()
	var isCycle bool

	for _, tc := range tcs {
		list := tc["list"].([]int)
		expected := tc["expected"].(int)

		if expected == -1 {
			head := linkedlist.CreateFromArray(list)
			isCycle = cycle.HasCycle(head)
		} else {
			head := linkedlist.CreateCycleFromArray(list, expected)
			isCycle = cycle.HasCycle(head)
		}

		assert.Equal(t, expected >= 0, isCycle)
	}
}

func Test_GetCycleHead(t *testing.T)  {
	tcs := getTestcases()

	for _, tc := range tcs {
		list := tc["list"].([]int)
		expected := tc["expected"].(int)

		if expected == -1 {
			head := linkedlist.CreateFromArray(list)
			cycleHead := cycle.GetCycleHead(head)
			assert.Nil(t, cycleHead)
		} else {
			head := linkedlist.CreateCycleFromArray(list, expected)
			cycleHead := cycle.GetCycleHead(head)
			assert.Equal(t, list[expected], cycleHead.Val)
		}
	}
}

func Test_HasCycleSlowAndFastPointers(t *testing.T)  {
	tcs := getTestcases()
	var isCycle bool

	for _, tc := range tcs {
		list := tc["list"].([]int)
		expected := tc["expected"].(int)

		if expected == -1 {
			head := linkedlist.CreateFromArray(list)
			isCycle = cycle.HasCycleSlowAndFastPointers(head)
		} else {
			head := linkedlist.CreateCycleFromArray(list, expected)
			isCycle = cycle.HasCycleSlowAndFastPointers(head)
		}

		assert.Equal(t, expected >= 0, isCycle)
	}
}

func Test_GetCycleHeadSlowAndFastPointers(t *testing.T)  {
	tcs := getTestcases()

	for _, tc := range tcs {
		list := tc["list"].([]int)
		expected := tc["expected"].(int)

		if expected == -1 {
			head := linkedlist.CreateFromArray(list)
			cycleHead := cycle.GetCycleHeadUsingSlowAndFastPointers(head)
			assert.Nil(t, cycleHead)
		} else {
			head := linkedlist.CreateCycleFromArray(list, expected)
			cycleHead := cycle.GetCycleHeadUsingSlowAndFastPointers(head)
			assert.Equal(t, list[expected], cycleHead.Val)
		}
	}
}

func getTestcases() map[string]map[string]interface{} {
	return map[string]map[string]interface{} {
		"tc1": {
			"list": []int{3,2,0,-4},
			"expected": 1,
		},
		"tc2": {
			"list": []int{1,2},
			"expected": 0,
		},
		"tc3": {
			"list": []int{1},
			"expected": -1,
		},
	}
}
