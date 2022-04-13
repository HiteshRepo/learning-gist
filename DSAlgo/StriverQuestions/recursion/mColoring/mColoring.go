package mColoring

func MColoring(n, m, e int, edges [][]int) bool {
	graph := make(map[int][]int)
	visited := make(map[int]int)

	for _,edge := range edges {
		if _,ok := graph[edge[0]]; !ok {
			graph[edge[0]] = []int{edge[1]}
		} else {
			graph[edge[0]] = append(graph[edge[0]], edge[1])
		}

		if _,ok := graph[edge[1]]; !ok {
			graph[edge[1]] = []int{edge[0]}
		} else {
			graph[edge[1]] = append(graph[edge[1]], edge[0])
		}
	}

	return traverseGraph(visited, graph, m, n, 0)
}

func traverseGraph(visited map[int]int, graph map[int][]int, m, n, currNode int) bool {
	if currNode == n {
		return true
	}

	for col:=1; col<=m; col++ {
		if isItPossibleToColor(currNode, col, visited, graph) {
			visited[currNode] = col

			if traverseGraph(visited, graph, m, n, currNode + 1) {
				return true
			}

			visited[currNode] = 0
		}
	}

	return false
}

func isItPossibleToColor(currNode, currColor int, visited map[int]int, graph map[int][]int) bool {
	for _, nbr := range graph[currNode] {
		if nbrColor,ok := visited[nbr]; ok && nbrColor == currColor {
			return false
		}
	}

	return true
}
