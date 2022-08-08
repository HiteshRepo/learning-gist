package subsetSums

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
	"math"
)

func SubsetSumsByRecursion(arr []int, _ int) []int {
	subsets := make([][]int, 0)
	subset := make([]*int, len(arr))

	cloneArr := make([]*int, 0)
	for _,n := range arr {
		num := n
		cloneArr = append(cloneArr, &num)
	}

	getAllSubsetsRecursively(cloneArr, subset, 0, &subsets)
	sums := make([]int, 0)

	for _, subset := range subsets {
		sum := calcSum(subset)
		sums = append(sums, sum)
	}

	utils.MergeSort(sums, 0 , len(sums)-1)

	return sums
}

func getAllSubsetsRecursively(arr, subset []*int, idx int, subsets *[][]int) {
	if idx == len(arr) {
		*subsets = append(*subsets, removeNullFromArr(subset))
		return
	}

	subset[idx] = nil
	getAllSubsetsRecursively(arr, subset, idx+1, subsets)
	subset[idx] = arr[idx]
	getAllSubsetsRecursively(arr, subset, idx+1, subsets)
}

func removeNullFromArr(arr []*int) []int {
	resArr := make([]int, 0)

	for _, n := range arr {
		if n != nil {
			resArr = append(resArr, *n)
		}
	}

	return resArr
}

func SubsetSumsByPowerSet(arr []int, _ int) []int {
	subsets := make([][]int, 0)
	subsets = append(subsets, []int{})

	for _, n := range arr {
		for _, subset := range subsets {
			currSet := append(subset, n)
			subsets = append(subsets, currSet)
		}
	}

	sums := make([]int, 0)

	for _, subset := range subsets {
		sum := calcSum(subset)
		sums = append(sums, sum)
	}

	utils.MergeSort(sums, 0 , len(sums)-1)

	return sums
}

func SubsetSumsByPowerBit(arr []int, _ int) []int {
	subsets := getAllSubsetsByPowerBit(arr)
	sums := make([]int, 0)

	for _, subset := range subsets {
		sum := calcSum(subset)
		sums = append(sums, sum)
	}

	utils.MergeSort(sums, 0 , len(sums)-1)

	return sums
}

func getAllSubsetsByPowerBit(arr []int) [][]int {
	subsets := make([][]int, 0)
	subsetsCount := math.Pow(float64(2), float64(len(arr)))
	for i:=0; i<int(subsetsCount); i++ {
		currSet := make([]int, 0)
		for j:=0; j<len(arr); j++ {
			if getBit(i, j) == 1 {
				currSet = append(currSet, arr[j])
			}
		}
		subsets = append(subsets, currSet)
	}
	return subsets
}

func getBit(num, bit int) int {
	tmp := 1 << bit
	tmp = tmp & num
	if tmp == 0 {
		return 0
	}
	return 1
}

func calcSum(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum
}