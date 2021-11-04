package main

type Vertex struct {
	key int
	nbrs []*Vertex
}

type Edge struct {
	to *Vertex
	from *Vertex
	wt int
	isDirected bool
}

func (g *Graph) getEdgeByVertices(src int, dest int) int {
	for _,e := range g.edges {
		if e.isDirected {
			if e.to.key == dest && e.from.key == src {
				return e.wt
			}
		} else {
			if (e.to.key == dest && e.from.key == src) || (e.from.key == dest && e.to.key == src) {
				return e.wt
			}
		}
	}
	return 0
}

