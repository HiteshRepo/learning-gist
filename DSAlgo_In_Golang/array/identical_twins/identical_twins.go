package identical_twins

import "fmt"

func GetIdenticalTwinsCount(arr []int) int {
	numTracker := make(map[int][]int)
	twins := 0
	indexes := make([][]int, 0)

	for i,n := range arr {
		if _, ok := numTracker[n]; !ok {
			numTracker[n] = []int{i}
			continue
		}
		numTracker[n] = append(numTracker[n], i)
	}

	for _,pos := range numTracker {
		twins += (len(pos) * (len(pos) - 1)) / 2
	}

	for _,pos := range numTracker {
		for i,p1 := range pos {
			for j:=i+1; j<len(pos); j++ {
				indexes = append(indexes, []int{p1, pos[j]})
			}
		}
	}

	fmt.Printf("Indexes for %v: %v\n", arr, indexes)

	return twins
}
