package main

import "fmt"

func main() {
	hasPathTest()
}

func hasPathTest() {
	graph := &Graph{isDirected: false}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,10)
	graph.addEdge(1,2,10)
	graph.addEdge(2,3,10)
	graph.addEdge(0,3,10)
	graph.addEdge(3,4,10)
	graph.addEdge(4,5,10)
	graph.addEdge(5,6,10)
	graph.addEdge(4,6,10)

	graph.display()

	src := 0
	dest := 6
	visited := make([]bool, 7)
	isPathPresent := graph.hasPath(src, dest, visited)
	if isPathPresent {
		fmt.Printf("Path exists between %v and %v\n", src, dest)
		return
	}
	fmt.Printf("No path exists between %v and %v", src, dest)
}
