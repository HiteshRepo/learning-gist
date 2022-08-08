package ratInAMaze

import "fmt"

func FindAllPaths(maze [][]int) []string {
	res := make([]string, 0)
	visited := make([][]int, 0)
	for i:=0; i<len(maze); i++ {
		row := make([]int, 0)
		for j:=0; j<len(maze[i]); j++ {
			row = append(row, 0)
		}
		visited = append(visited, row)
	}
	getAllPaths(maze, &res, visited,"", 0, 0)
	return res
}


func getAllPaths(maze [][]int, paths *[]string, visited [][]int, currPath string, n, m int) {
	if n<0 || m<0 || n==len(maze) || m==len(maze) || maze[n][m]==0 || visited[n][m]==1 {
		return
	}

	if n == len(maze)-1 && m == len(maze)-1 {
		*paths = append(*paths, currPath)
		return
	}

	visited[n][m] = 1
	getAllPaths(maze, paths, visited, fmt.Sprintf("%sD", currPath), n+1, m)
	getAllPaths(maze, paths, visited, fmt.Sprintf("%sL", currPath), n, m-1)
	getAllPaths(maze, paths, visited, fmt.Sprintf("%sR", currPath), n, m+1)
	getAllPaths(maze, paths, visited, fmt.Sprintf("%sU", currPath), n-1, m)
	visited[n][m] = 0

	return
}