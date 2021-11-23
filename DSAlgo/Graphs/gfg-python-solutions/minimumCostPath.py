row = [-1, 0, 0, 1]
col = [0, -1, 1, 0]

class Solution:

    #Function to return the minimum cost to react at bottom
	#right cell from top left cell.
    def minimumCostPath(self, grid):
        return self.findPath(grid, len(grid))

    def checkIfReachedDestination(self, i, j, N):
        return i == N-1 and j == N-1

    def getLeastCost(self, curr, matrix, leastCost):
        if curr[0] < leastCost:
            return curr[0]
        return leastCost

    def addToQueue(self, curr, matrix, visited, q, N):
        for k in range(len(row)):
            if isValid(curr[1].x + row[k],  curr[1].y + col[k], N-1, N-1):
                next = Node(curr[1].x + row[k], curr[1].y + col[k], curr)
                key = (curr[1].x + row[k], curr[1].y + col[k])
                cost = curr[0] + matrix[next.x][next.y]
                if key not in visited:
                    q.put((cost, next))
                    visited.add(key)

    def findPath(self, matrix, N, x=0, y=0):
	    if not matrix or not len(matrix):
	        return 0

	    import sys
	    import queue

	    leastCost = sys.maxsize
	    q = queue.PriorityQueue()
	    src = Node(0, 0)
	    q.put((matrix[0][0], src))

	    visited = set()
	    key = (src.x, src.y)
	    visited.add(key)

	    while not q.empty():
	        curr = q.get()

	        if self.checkIfReachedDestination(curr[1].x, curr[1].y, N):
	            leastCost = self.getLeastCost(curr, matrix, leastCost)
	        self.addToQueue(curr, matrix, visited, q, N)
	    return leastCost

class Node:
    def __init__(self, x, y, parent=None):
        self.x = x
        self.y = y
        self.parent = parent

    def __repr__(self):
        return str((self.x, self.y))

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

def isValid(x, y, X, Y):
    return (0 <= x <= X) and (0 <= y <= Y)

#{
#  Driver Code Starts
if __name__ == '__main__':
# 	T=int(input())
# 	for i in range(T):
# 		n = int(input())
# 		grid = []
# 		for _ in range(n):
# 			a = list(map(int, input().split()))
# 			grid.append(a)
    #grid = [[9, 4, 9, 9], [6, 7, 6, 4], [8, 3, 3, 7], [7, 4, 9, 10]]
    grid = [[4,5], [3, 7]]
    obj = Solution()
    ans = obj.minimumCostPath(grid)
    print(ans)

# } Driver Code Ends