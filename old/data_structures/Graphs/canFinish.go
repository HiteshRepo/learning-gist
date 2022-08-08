package main

func canFinish(numCourses int, prerequisites [][]int) bool {

	allCourses := make([]int, 0)
	inProgressCourses := make([]int, 0)
	completedCourses := make([]int, 0)
	courseMap := make(map[int][]int)

	for _,pairs := range prerequisites {

		if len(pairs) > 2 || pairs[0] == pairs[1] {
			return false
		}

		if preReqs,ok := courseMap[pairs[0]]; ok {
			all := append(preReqs, pairs[1])
			courseMap[pairs[0]] = all
		} else {
			courseMap[pairs[0]] = []int{pairs[1]}
		}

		if !contains(allCourses, pairs[0]) {
			allCourses = append(allCourses, pairs[0])
		}

		if !contains(allCourses, pairs[1]) {
			allCourses = append(allCourses, pairs[1])
		}
	}

	for _,c := range allCourses {
		if dfs(&allCourses, &inProgressCourses, &completedCourses, c, courseMap) {
			return false
		}
	}
	return true
}

func dfs(allCourses, inProgressCourses, completedCourses *[]int, currentCourse int, courseMap map[int][]int) bool {

	moveCourse(currentCourse, allCourses, inProgressCourses)

	for _,nbr := range courseMap[currentCourse] {

		if contains(*completedCourses, nbr) {
			continue
		}

		if contains(*inProgressCourses, nbr) {
			return true
		}

		if dfs(allCourses, inProgressCourses, completedCourses, nbr, courseMap) {
			return true
		}
	}

	moveCourse(currentCourse, inProgressCourses, completedCourses)
	return false
}

func moveCourse(course int, source, destination *[]int) {
	idx := findIndex(course, *source)
	if idx != -1 {
		*source = append((*source)[0:idx], (*source)[idx+1:]...)
	}
	*destination = append(*destination, course)
}

func findIndex(course int, lst []int) int {
	idx := -1

	for i,c := range lst {
		if c == course {
			return i
		}
	}

	return idx
}
