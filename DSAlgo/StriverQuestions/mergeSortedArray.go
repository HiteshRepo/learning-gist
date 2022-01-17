package main

import (
	"fmt"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
)

func mergeSortedArr(nums1 []int, m int, nums2 []int, n int) {

	if m == 0 {
		i := m
		for _, num := range nums2{
			nums1[i] = num
			i++
		}
		return
	}

	if n == 0 {
		return
	}

	for i:=0; i<m; i++{
		if nums1[i] > nums2[0] {
			temp := nums1[i]
			nums1[i] = nums2[0]
			nums2[0] = temp
		}

		var k int
		first := nums2[0]
		for k=1; k<n && nums2[k]<first; k++ {
			nums2[k-1] = nums2[k]
		}
		nums2[k-1] = first
	}

	i := m
	for _, num := range nums2{
		nums1[i] = num
		i++
	}
}

func mergeSortedArrWithGap(nums1 []int, m int, nums2 []int, n int)  {

	gap := (int(m+n)/2) + ((m+n) % 2)

	for ;gap > 0;{

		i := 0
		for ;i+gap < m;i++ {
			if nums1[i] > nums1[i+gap]{
				nums1[i],nums1[i+gap] = nums1[i+gap],nums1[i]
			}
		}
		j := 0
		if gap > m{
			j = gap-m
		}

		for ;i < m && j < n;{
			if nums1[i] > nums2[j]{
				nums1[i],nums2[j] = nums2[j],nums1[i]
			}
			i++
			j++
		}

		if j < n{
			j = 0
			for ;j+gap < n;j++{
				if nums2[j] > nums2[j+gap]{
					nums2[j],nums2[j+gap] = nums2[j+gap],nums2[j]
				}
			}
		}

		gap = gapFn(gap)


	}
	i := m
	for _, num := range nums2{
		nums1[i] = num
		i++
	}

}

func gapFn(prevGap int) int {
	if prevGap <= 1{
		return 0
	}

	return (prevGap / 2) + (prevGap % 2)
}

func runTestsForMergeSortedArrays() {
	testCases := map[string]map[string][]int{
		"tc1": {
			"arr1":    {1,2,3,0,0,0},
			"arr1Len": {3},
			"arr2":    {2,5,6},
			"expected": {1,2,2,3,5,6},
		},
		"tc2": {
			"arr1":    {1},
			"arr1Len": {1},
			"arr2": {},
			"expected": {1},
		},
	}

	for name, tc := range testCases {
		fmt.Printf("running test-case: %s\n", name)
		arr1 := tc["arr1"]
		arr2 := tc["arr2"]
		expected := tc["expected"]
		// mergeSortedArrWithGap(arr1, tc["arr1Len"][0], arr2, len(arr2))
		mergeSortedArr(arr1, tc["arr1Len"][0], arr2, len(arr2))
		fmt.Println("is solution correct: ", utils.IntArrayEquals(arr1, expected))
	}
}
