package main

import (
	"container/heap"
	"fmt"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1
const LargestPathWeight = "largestPathWt"
const SmallestPathWeight = "smallestPathWt"
const CeilPathWeight = "ceilPathWt"
const FloorPathWeight = "floorPathWt"
const LargestPath = "largestPath"
const SmallestPath = "smallestPath"
const CeilPath = "ceilPath"
const FloorPath = "floorPath"

func main() {

	// <----- BFS ----->
	//BFSTraverse()
	//isGraphCyclic()
	//isGraphBipartite()
	//spreadInfection()
	//dijkstra()
	prims()


	// <----- DFS ----->
	//hasPathTest()
	//displayConnectedVerticesTest()
	//isGraphConnected()
	//noOfIslands()
	//perfectFriends()
	//hamiltonianPathAndCycles()
	//knightsTour()
}

// BFS solutions

func BFSTraverse() {
	graph := &Graph{}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,10, false)
	graph.addEdge(1,2,10, false)
	graph.addEdge(2,3,10, false)
	graph.addEdge(0,3,10, false)
	graph.addEdge(3,4,10, false)
	graph.addEdge(4,5,10, false)
	graph.addEdge(5,6,10, false)
	graph.addEdge(4,6,10, false)

	graph.display()

	src := 2
	visited := make([]bool, 7)
	q := NewQueue()
	q.Add(&Pair{src, fmt.Sprintf("%d", src), 0})
	graph.traverseBSF(visited, q)
}

func isGraphCyclic() {
	graph := &Graph{}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,10, false)
	graph.addEdge(1,2,10, false)
	graph.addEdge(2,3,10, false)
	graph.addEdge(0,3,10, false)
	graph.addEdge(3,4,10, false)
	graph.addEdge(4,5,10, false)
	graph.addEdge(5,6,10, false)
	graph.addEdge(4,6,10, false)

	graph.display()

	graph.isGraphCyclic()

	graph = &Graph{}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,10, false)
	graph.addEdge(2,3,10, false)
	graph.addEdge(4,5,10, false)
	graph.addEdge(5,6,10, false)
	graph.addEdge(4,6,10, false)

	graph.display()

	graph.isGraphCyclic()
}

// graph is bipartite if it is possible to:
// divide vertices into - 2 mutually exclusive and exhaustive sets
// such that all edges are across sets
// solution : all non-cyclic graphs are bipartite, if cyclic then they should be of even sizes
func isGraphBipartite() {
	graph := &Graph{}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,10, false)
	graph.addEdge(1,2,10, false)
	graph.addEdge(2,3,10, false)
	graph.addEdge(0,3,10, false)
	graph.addEdge(3,4,10, false)
	graph.addEdge(4,5,10, false)
	graph.addEdge(5,6,10, false)
	graph.addEdge(4,6,10, false)

	graph.display()

	graph.isGraphBipartite()

	graph = &Graph{}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,10, false)
	graph.addEdge(1,2,10, false)
	graph.addEdge(2,3,10, false)
	graph.addEdge(3,4,10, false)

	graph.display()

	graph.isGraphBipartite()
}

func spreadInfection() {
	graph := &Graph{}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,10, false)
	graph.addEdge(1,2,10, false)
	graph.addEdge(2,3,10, false)
	graph.addEdge(0,3,10, false)
	graph.addEdge(3,4,10, false)
	graph.addEdge(4,5,10, false)
	graph.addEdge(5,6,10, false)
	graph.addEdge(4,6,10, false)

	graph.display()

	src := 6
	time := 3
	result := graph.traverseBFSAndStopAtGivenLevel(src, time)
	fmt.Println(result)
}

func dijkstra() {
	graph := &Graph{}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,10, false)
	graph.addEdge(1,2,10, false)
	graph.addEdge(2,3,10, false)
	graph.addEdge(0,3,40, false)
	graph.addEdge(3,4,2, false)
	graph.addEdge(4,5,3, false)
	graph.addEdge(5,6,3, false)
	graph.addEdge(4,6,8, false)

	graph.display()

	graph.dijkstraAlgo(0)
}

func prims() {
	graph := &Graph{}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,10, false)
	graph.addEdge(1,2,10, false)
	graph.addEdge(2,3,10, false)
	graph.addEdge(0,3,40, false)
	graph.addEdge(3,4,2, false)
	graph.addEdge(4,5,3, false)
	graph.addEdge(5,6,3, false)
	graph.addEdge(4,6,8, false)

	graph.display()

	graph.minimumSpanningTreePrims(0)
}

// DFS solutions
func hasPathTest() {
	graph := &Graph{}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,10, false)
	graph.addEdge(1,2,10, false)
	graph.addEdge(2,3,10, false)
	graph.addEdge(0,3,40, false)
	graph.addEdge(3,4,2, false)
	graph.addEdge(4,5,3, false)
	graph.addEdge(5,6,3, false)
	graph.addEdge(4,6,8, false)

	graph.display()

	src := 0
	dest := 6
	visited := make([]bool, 7)
	isPathPresent := graph.hasPath(src, dest, visited)
	if isPathPresent {
		fmt.Printf("Path exists between %v and %v\n", src, dest)
	} else {
		fmt.Printf("No path exists between %v and %v\n", src, dest)
	}

	allPaths := make([]string, 0)
	visited = make([]bool, 7)
	graph.getAllPaths(src, dest, visited, fmt.Sprintf("%d", src), &allPaths)
	if len(allPaths) > 0 {
		fmt.Printf("The paths b/w %d and %d are : \n", src, dest)
	}
	for _,path := range allPaths {
		fmt.Println(path)
	}

	allPathsWithWeight := make(map[string]int)
	visited = make([]bool, 7)
	graph.getAllPathsWithCumulativeWeight(src, dest, visited, fmt.Sprintf("%d", src), 0, &allPathsWithWeight)
	if len(allPathsWithWeight) > 0 {
		fmt.Printf("The paths b/w %d and %d are : \n", src, dest)
	}
	for path,wt := range allPathsWithWeight {
		fmt.Printf("%s : %d\n", path, wt)
	}

	solutions := make(map[string]interface{})
	solutions[SmallestPathWeight] = MaxInt
	solutions[LargestPathWeight] = MinInt
	solutions[CeilPathWeight] = MaxInt
	solutions[FloorPathWeight] = MinInt
	solutions[SmallestPath] = ""
	solutions[LargestPath] = ""
	solutions[CeilPath] = ""
	solutions[FloorPath] = ""

	threshold := 40
	kthLargest := 3

	priorityQueue := NewPriorityQueue()

	graph.multiSolver(src, dest, threshold, visited, fmt.Sprintf("%d", src), 0, &solutions, &priorityQueue, kthLargest)
	heap.Init(&priorityQueue)

	path := heap.Pop(&priorityQueue).(*Path)

	fmt.Printf("Smallest path between %d and %d is %s\n", src, dest, solutions[SmallestPath])
	fmt.Printf("Largest path between %d and %d is %s\n", src, dest, solutions[LargestPath])
	fmt.Printf("Ceil path between %d and %d for weight %d is %s\n", src, dest, threshold, solutions[CeilPath])
	fmt.Printf("Floor path between %d and %d for weight %d is %s\n", src, dest, threshold, solutions[FloorPath])
	fmt.Printf("3rd largest path is %s with weight %d \n", path.Name, path.Weight)
}

func displayConnectedVerticesTest() {
	graph := &Graph{}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,10, false)
	graph.addEdge(2,3,10, false)
	graph.addEdge(4,5,10, false)
	graph.addEdge(5,6,10, false)
	graph.addEdge(4,6,10, false)

	graph.display()

	visited := make([]bool, 7)
	currentPath := make([]int, 0)
	connectedPaths := make([][]int, 0)
	graph.displayAllConnectedComponents(visited, &currentPath, &connectedPaths)

	fmt.Println(connectedPaths)
}

func isGraphConnected() {
	graph := &Graph{}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,10, false)
	graph.addEdge(2,3,10, false)
	graph.addEdge(4,5,10, false)
	graph.addEdge(5,6,10, false)
	graph.addEdge(4,6,10, false)

	graph.display()
	visited := make([]bool, 7)
	currentPath := make([]int, 0)
	connectedPaths := make([][]int, 0)
	graph.displayAllConnectedComponents(visited, &currentPath, &connectedPaths)
	fmt.Printf("Is above graph connected? %v\n", len(connectedPaths) == 1)

	graph = &Graph{}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,10, false)
	graph.addEdge(1,2,10, false)
	graph.addEdge(2,3,10, false)
	graph.addEdge(0,3,40, false)
	graph.addEdge(3,4,2, false)
	graph.addEdge(4,5,3, false)
	graph.addEdge(5,6,3, false)
	graph.addEdge(4,6,8, false)

	graph.display()
	visited = make([]bool, 7)
	currentPath = make([]int, 0)
	connectedPaths = make([][]int, 0)
	graph.displayAllConnectedComponents(visited, &currentPath, &connectedPaths)
	fmt.Printf("Is above graph connected? %v\n", len(connectedPaths) == 1)
}

func noOfIslands() {
	island := [][]int{
		{0,0,1,1,1,1,1,1},
		{0,0,1,1,1,1,1,1},
		{1,1,1,1,1,1,1,0},
		{1,1,0,0,0,1,1,0},
		{1,1,1,1,0,1,1,0},
		{1,1,1,1,0,1,1,0},
		{1,1,1,1,1,1,1,0},
		{1,1,1,1,1,1,1,0},
	}

	visited := [][]bool{
		{false,false,false,false,false,false,false,false},
		{false,false,false,false,false,false,false,false},
		{false,false,false,false,false,false,false,false},
		{false,false,false,false,false,false,false,false},
		{false,false,false,false,false,false,false,false},
		{false,false,false,false,false,false,false,false},
		{false,false,false,false,false,false,false,false},
		{false,false,false,false,false,false,false,false},
	}
	count := 0

	for i:=0; i<len(island); i++ {
		for j:=0; j<len(island[i]); j++ {
			if visited[i][j] || island[i][j] == 1 {
				continue
			}
			visitConnectedLands(island, i, j, visited)
			count++
		}
	}
	fmt.Printf("No. of connected islands are: %d\n", count)
}

func visitConnectedLands(island [][]int, i,j int, visited [][]bool) {
	if i < 0 || j <0 || i >= len(island) || j >= len(island[i]) || island[i][j] == 1 || visited[i][j] {
		return
	}

	visited[i][j] = true
	visitConnectedLands(island, i-1, j, visited)
	visitConnectedLands(island, i, j+1, visited)
	visitConnectedLands(island, i+1, j, visited)
	visitConnectedLands(island, i, j-1, visited)
}

func perfectFriends() {
	graph := &Graph{}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,0, false)
	graph.addEdge(2,3,0, false)
	graph.addEdge(4,5,0, false)
	graph.addEdge(5,6,0, false)
	graph.addEdge(4,6,0, false)

	graph.display()

	visited := make([]bool, 7)
	currentPath := make([]int, 0)
	connectedPaths := make([][]int, 0)
	graph.displayAllConnectedComponents(visited, &currentPath, &connectedPaths)

	fmt.Println(connectedPaths)

	pairs := 0

	for i:=0; i<len(connectedPaths); i++ {
		for j:=i+1; j<len(connectedPaths); j++ {
			pairs += len(connectedPaths[i]) * len(connectedPaths[j])
		}
	}

	fmt.Printf("No of perfect freinds: %d\n", pairs)
}

func hamiltonianPathAndCycles() {
	graph := &Graph{}

	for i:=0; i<7; i++ {
		graph.addVertex(i)
	}

	graph.addEdge(0,1,10, false)
	graph.addEdge(1,2,10, false)
	graph.addEdge(2,3,10, false)
	graph.addEdge(0,3,10, false)
	graph.addEdge(3,4,10, false)
	graph.addEdge(4,5,10, false)
	graph.addEdge(2,5,10, false)
	graph.addEdge(5,6,10, false)
	graph.addEdge(4,6,10, false)

	graph.display()

	hamiltonianPaths := graph.getAllPathsHamiltonian(0)
	for path,isCycle := range hamiltonianPaths {
		fmt.Println(path, isCycle)
	}
}

func knightsTour() {
	dimension := 5
	row := 2
	column := 2

	chessBoard := make([][]int, dimension)
	for i := range chessBoard {
		chessBoard[i] = make([]int, dimension)
	}

	printKnightsTour(chessBoard, row, column, 1)
}

func printKnightsTour(board [][]int, row, column, move int) {
	if row < 0 || column < 0 || row >= len(board) || column >= len(board) || board[row][column] > 0 {
		return
	} else if move == (len(board) * len(board)) {
		board[row][column] = move
		displayBoard(board)
		board[row][column] = 0
		return
	}

	board[row][column] = move

	printKnightsTour(board, row-2, column+1, move+1)
	printKnightsTour(board, row-1, column+2, move+1)
	printKnightsTour(board, row+1, column+2, move+1)
	printKnightsTour(board, row+2, column+1, move+1)
	printKnightsTour(board, row+2, column-1, move+1)
	printKnightsTour(board, row+1, column-2, move+1)
	printKnightsTour(board, row-1, column-2, move+1)
	printKnightsTour(board, row-2, column-1, move+1)

	board[row][column] = 0
}

func displayBoard(board [][]int) {
	for i:=0; i<len(board); i++ {
		for j:=0; j<len(board[i]); j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

