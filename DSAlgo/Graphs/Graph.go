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
			heap.Push(pq,&Path{Name: psf, Weight: wsf, Index: pq.Len()})
		} else {
			leastPriority := heap.Pop(pq).(*Path)
			if leastPriority.Weight < wsf {
				heap.Push(pq,&Path{Name: psf, Weight: wsf, Index: pq.Len()})
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

func (g *Graph) display() {
	for _,v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _,n := range v.nbrs {
			fmt.Printf(" %v ", n.key)
		}
	}
	fmt.Printf("\n")
}