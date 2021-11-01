package main

import "fmt"

type Graph struct {
	vertices []*Vertex
	isDirected bool
}

func (g *Graph) addVertex(key int) {
	if hasVertex(g.vertices ,key) {
		err := fmt.Errorf("%v vertex already exists", key)
		fmt.Println(err.Error())
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: key})
}

func (g *Graph) addEdge(src, dest, wt int) {
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

	if !g.isDirected {
		if hasVertex(destVertex.nbrs, src) {
			err := fmt.Errorf("edge (%v -> %v) already exists", dest, src)
			fmt.Println(err.Error())
			return
		}
		destVertex.nbrs = append(destVertex.nbrs, srcVertex)
	}

	srcVertex.nbrs = append(srcVertex.nbrs, destVertex)
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
	if src == dest {
		return true
	}

	visited[src] = true
	for _,n := range g.vertices[src].nbrs {
		if visited[n.key] {
			continue
		}

		hasNbrPath := g.hasPath(n.key, dest, visited)
		if hasNbrPath {
			return true
		}
	}

	return false
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