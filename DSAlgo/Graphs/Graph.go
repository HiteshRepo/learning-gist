package main

import (
	"container/heap"
	"fmt"
)

type Graph struct {
	vertices []*Vertex
	edges []*Edge
}

func (g *Graph) addVertex(key int) {
	if hasVertex(g.vertices ,key) {
		err := fmt.Errorf("%v vertex already exists", key)
		fmt.Println(err.Error())
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: key})
}

func (g *Graph) addEdge(src, dest, wt int, isDirected bool) {
	srcVertex := g.getVertex(src)
	destVertex := g.getVertex(dest)

	if srcVertex == nil || destVertex == nil {
		err := fmt.Errorf("invalid edge (%v -> %v)", src, dest)
		fmt.Println(err.Error())
		return
	}

	if hasVertex(srcVertex.nbrs, dest) {
		err := fmt.Errorf("edge (%v -> %v) already exists", src, dest)
		fmt.Println(err.Error())
		return
	}

	if !isDirected {
		if hasVertex(destVertex.nbrs, src) {
			err := fmt.Errorf("edge (%v -> %v) already exists", dest, src)
			fmt.Println(err.Error())
			return
		}
		destVertex.nbrs = append(destVertex.nbrs, srcVertex)
	}

	srcVertex.nbrs = append(srcVertex.nbrs, destVertex)
	g.edges = append(g.edges ,&Edge{to: destVertex, from: srcVertex, wt: wt, isDirected: isDirected})
}

func (g *Graph) getVertex(key int) *Vertex {
	for idx,v := range g.vertices {
		if key == v.key {
			return g.vertices[idx]
		}
	}
	return nil
}

func hasVertex(vertices []*Vertex, key int) bool {
	for _,v := range vertices {
		if key == v.key {
			return true
		}
	}
	return false
}

// BFS

func (g *Graph) traverseBSF(visited []bool, q *Queue) {

	for !q.IsEmpty() {
		// remove
		removedPair := q.RemoveFirst()

		// mark*
		if visited[removedPair.v] {
			continue
		}
		visited[removedPair.v] = true

		// work
		fmt.Printf("%d @ %s\n", removedPair.v, removedPair.psf)

		// add*
		for _,n := range g.vertices[removedPair.v].nbrs {
			if visited[n.key] {
				continue
			}
			q.Add(&Pair{n.key, fmt.Sprintf("%s%d", removedPair.psf, n.key), removedPair.lvl + 1})
		}
	}
}

func (g *Graph) isGraphCyclic() {
	visited := make([]bool, 7)
	for _,v := range g.vertices {
		if visited[v.key] {
			continue
		}
		if g.isVertexCyclic(visited, v.key) {
			fmt.Println("The graph has a cycle")
			return
		}
	}
}

func (g *Graph) isVertexCyclic(visited []bool, vertex int) bool {
	q := NewQueue()
	q.Add(&Pair{vertex, fmt.Sprintf("%d", vertex), 0})

	for !q.IsEmpty() {
		// remove
		removedPair := q.RemoveFirst()

		// mark*
		if visited[removedPair.v] {
			fmt.Printf("cycle detected : %d @ %s\n", removedPair.v, removedPair.psf)
			return true
		}
		visited[removedPair.v] = true

		// work
		// no work

		// add*
		for _,n := range g.vertices[removedPair.v].nbrs {
			if visited[n.key] {
				continue
			}
			q.Add(&Pair{n.key, fmt.Sprintf("%s%d", removedPair.psf, n.key), removedPair.lvl + 1})
		}
	}

	return false
}

func (g *Graph) isGraphBipartite() {
	visited := make([]int, len(g.vertices))
	for i:=0; i<len(visited); i++ {
		visited[i] = -1
	}

	for _,v := range g.vertices {
		if visited[v.key] != -1 {
			continue
		}
		if g.checkBipartite(visited, v.key) {
			fmt.Println("graph is bipartite")
			return
		}
	}
	fmt.Println("graph is not bipartite")
}

func (g *Graph) checkBipartite(visited []int, vertex int) bool {
	q := NewQueue()
	q.Add(&Pair{vertex, fmt.Sprintf("%d", vertex), 0})

	for !q.IsEmpty() {
		// remove
		removedPair := q.RemoveFirst()

		// mark*
		if visited[removedPair.v] != -1 {
			fmt.Printf("cycle detected : %d @ %s\n", removedPair.v, removedPair.psf)
			if removedPair.lvl != visited[removedPair.v] {
				return false
			}

		} else {
			visited[removedPair.v] = removedPair.lvl
		}

		// work
		// no work

		// add*
		for _,n := range g.vertices[removedPair.v].nbrs {
			if visited[n.key] != -1 {
				continue
			}
			q.Add(&Pair{n.key, fmt.Sprintf("%s%d", removedPair.psf, n.key), removedPair.lvl + 1})
		}
	}

	return true
}

func (g *Graph) traverseBFSAndStopAtGivenLevel(src, time int) map[int][]int {
	visited := make([]int, len(g.vertices))
	for i:=0; i<len(visited); i++ {
		visited[i] = -1
	}

	count := 0

	q := NewQueue()
	q.Add(&Pair{src, fmt.Sprintf("%d", src), 0})

	for !q.IsEmpty() {
		// remove
		removedPair := q.RemoveFirst()

		// mark*
		if visited[removedPair.v] != -1 {
			continue
		}

		// work
		if removedPair.lvl < time {
			count++
		} else {
			break
		}
		visited[removedPair.v] = removedPair.lvl

		// add*
		for _,n := range g.vertices[removedPair.v].nbrs {
			if visited[n.key] != -1 {
				continue
			}
			q.Add(&Pair{n.key, fmt.Sprintf("%s%d", removedPair.psf, n.key), removedPair.lvl + 1})
		}
	}

	return map[int][]int{count: filterVisited(visited)}
}

func (g *Graph) dijkstraAlgo(src int) {
	visited := make([]bool, len(g.vertices))
	for i:=0; i<len(visited); i++ {
		visited[i] = false
	}

	q := NewPriorityQueue()
	heap.Init(&q)
	heap.Push(&q, &Path{Name: src, Psf: fmt.Sprintf("%d",src), Weight: 0, Index: q.Len()})

	for q.Len() > 0 {
		// remove
		removed := heap.Pop(&q).(*Path)

		// mark*
		if visited[removed.Name] {
			continue
		}
		visited[removed.Name] = true

		// work
		fmt.Printf("%d via %s @ %d\n", removed.Name, removed.Psf, removed.Weight)

		// add*
		for _,n := range g.vertices[removed.Name].nbrs {
			if visited[n.key] {
				continue
			}
			resWt := removed.Weight + g.getEdgeByVertices(n.key, removed.Name)
			heap.Push(&q, &Path{Name: n.key, Psf: fmt.Sprintf("%s%d", removed.Psf, n.key), Weight: resWt, Index: q.Len()})
		}
	}
}

// DFS

func (g *Graph) hasPath(src, dest int, visited []bool) bool {
	allPaths := make([]string, 0)
	g.getAllPaths(src, dest, visited, fmt.Sprintf("%d", src), &allPaths)
	if len(allPaths) > 0 {
		return true
	}
	return false
}

func (g *Graph) getAllPaths(src, dest int, visited []bool, psf string, allPaths *[]string) {
	if src == dest {
		*allPaths = append(*allPaths, psf)
		return
	}

	visited[src] = true
	for _,n := range g.vertices[src].nbrs {
		if visited[n.key] {
			continue
		}
		g.getAllPaths(n.key, dest, visited, fmt.Sprintf("%s%d", psf, n.key), allPaths)
	}
	visited[src] = false
}

func (g *Graph) drawConnectedVerticesHamiltonian(src int, visited map[int]bool, psf string, hamiltonianPaths *map[string]bool, oSrc int) {
	if len(psf) == len(g.vertices) {
		(*hamiltonianPaths)[psf] = g.hasEdge(src, oSrc)
		return
	}

	visited[src] = true
	for _,n := range g.vertices[src].nbrs {
		if visited[n.key] {
			continue
		}
		g.drawConnectedVerticesHamiltonian(n.key, visited, fmt.Sprintf("%s%d", psf, n.key), hamiltonianPaths, oSrc)
	}
	visited[src] = false
}

func (g *Graph) getAllPathsHamiltonian(src int) map[string]bool {
	hamiltonianPaths := make(map[string]bool, 0)
	visited := make(map[int]bool, 0)
	g.drawConnectedVerticesHamiltonian(src, visited, fmt.Sprintf("%d", src), &hamiltonianPaths, src)
	// for all vertices
	//for _,v := range g.vertices {
	//	visited := make(map[int]bool, 0)
	//	g.drawConnectedVerticesHamiltonian(v.key, visited, fmt.Sprintf("%d", v.key), &hamiltonianPaths, v.key)
	//}

	return hamiltonianPaths
}

func (g *Graph) getAllPathsWithCumulativeWeight(src, dest int, visited []bool, psf string, wsf int, allPathsWeighted *map[string]int) {
	if src == dest {
		(*allPathsWeighted)[psf] = wsf
		return
	}

	visited[src] = true
	for _,n := range g.vertices[src].nbrs {
		if visited[n.key] {
			continue
		}
		wt := g.getEdgeByVertices(n.key, src)
		g.getAllPathsWithCumulativeWeight(n.key, dest, visited, fmt.Sprintf("%s%d", psf, n.key), wsf + wt ,allPathsWeighted)
	}
	visited[src] = false
}

func (g *Graph) multiSolver(src, dest, threshold int, visited []bool, psf string, wsf int, solutions *map[string]interface{}, pq *PriorityQueue, k int) {
	if src == dest {

		if wsf < (*solutions)[SmallestPathWeight].(int) {
			(*solutions)[SmallestPathWeight] = wsf
			(*solutions)[SmallestPath] = psf
		}

		if wsf > (*solutions)[LargestPathWeight].(int) {
			(*solutions)[LargestPathWeight] = wsf
			(*solutions)[LargestPath] = psf
		}

		if wsf > threshold && wsf < (*solutions)[CeilPathWeight].(int) {
			(*solutions)[CeilPathWeight] = wsf
			(*solutions)[CeilPath] = psf
		}

		if wsf < threshold && wsf > (*solutions)[FloorPathWeight].(int) {
			(*solutions)[FloorPathWeight] = wsf
			(*solutions)[FloorPath] = psf
		}

		if pq.Len() < k {
			heap.Push(pq,&Path{Psf: psf, Weight: wsf, Index: pq.Len()})
		} else {
			leastPriority := heap.Pop(pq).(*Path)
			if leastPriority.Weight < wsf {
				heap.Push(pq,&Path{Psf: psf, Weight: wsf, Index: pq.Len()})
			} else {
				heap.Push(pq,leastPriority)
			}
		}
		return
	}

	visited[src] = true
	for _,n := range g.vertices[src].nbrs {
		if visited[n.key] {
			continue
		}
		wt := g.getEdgeByVertices(n.key, src)
		g.multiSolver(n.key, dest, threshold, visited, fmt.Sprintf("%s%d", psf, n.key), wsf + wt, solutions, pq, k)
	}
	visited[src] = false
}

func (g *Graph) displayAllConnectedComponents(visited []bool, currentPath *[]int, connectedPaths *[][]int) {
	for _,n := range g.vertices {
		if visited[n.key] {
			continue
		}
		newArr := make([]int, 0)
		currentPath = &newArr
		drawConnectedVertices(g, n.key, currentPath, visited)
		*connectedPaths = append(*connectedPaths, *currentPath)
	}
}

func drawConnectedVertices(g *Graph, src int, currentPath *[]int, visited []bool) {

	visited[src] = true
	*currentPath = append(*currentPath, src)

	for _,n := range g.vertices[src].nbrs {
		if visited[n.key] {
			continue
		}
		drawConnectedVertices(g, n.key, currentPath, visited)
	}
}

func (g *Graph) display() {
	for _,v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _,n := range v.nbrs {
			fmt.Printf(" %v ", n.key)
		}
	}
	fmt.Printf("\n")
}

func areAllVisited(visited map[int]bool) bool {
	for _,v := range visited {
		if !v {
			return false
		}
	}
	return true
}

func contains(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}

	return false
}

func filterVisited(all []int) []int {
	visited := make([]int, 0)
	for i,n := range all {
		if n != -1 && !contains(visited, i) {
			visited = append(visited, i)
		}
	}
	return visited
}