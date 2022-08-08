package amazon

// Detect cycle in a directed graph
func DetectCycleInDirectedGraph(graph [][]int) bool {
	visited := make(map[int]struct{})
	ancestors := make(map[int]struct{})

	return traverse(ancestors, visited, graph, graph[0][0])
}

func traverse(ancestors, visited map[int]struct{}, graph [][]int, currentVertex int) bool {
	visited[currentVertex] = struct{}{}
	ancestors[currentVertex] = struct{}{}

	for _,nbr := range graph[currentVertex] {
		if _,ok := visited[nbr]; !ok {
			if traverse(ancestors, visited, graph, nbr) {
				return true
			}
		}

		if _,ok := ancestors[nbr]; ok {
			return true
		}
	}

	delete(ancestors, currentVertex)

	return false
}
