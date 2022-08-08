package main

import "fmt"

func isBipartite(graph [][]int) bool {

	visited := make([]int, len(graph))
	for i:=0; i<len(graph); i++ {
		visited[i] = -1
	}

	for v,_ := range graph {

		if visited[v] != -1 {
			continue
		}

		if !checkBipartite(visited, v, graph) {
			return false
		}
	}
	return true
}

func checkBipartite(visited []int, vertex int, graph [][]int) bool {
	q := NewQueue()
	q.Add(&Pair{vertex, fmt.Sprintf("%d", vertex), 0})

	for !q.IsEmpty() {
		removed := q.RemoveFirst()

		if visited[removed.v] != -1 {
			if removed.lvl != visited[removed.v] {
				return false
			}
		} else {
			visited[removed.v] = removed.lvl
		}

		for _,nbr := range graph[removed.v] {
			if visited[nbr] != -1 {
				continue
			}
			q.Add(&Pair{nbr, fmt.Sprintf("%s%d", removed.psf, nbr), removed.lvl+1})
		}
	}
	return true
}
