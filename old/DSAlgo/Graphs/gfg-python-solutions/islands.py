    #Function to find the number of islands.
    def numIslands(self, grid):
        visited = [[False for i in grid[0]] for j in grid]
        count = 0

        for i in range(len(grid)):
            for j in range(len(grid[i])):

                if visited[i][j] or int(grid[i][j]) == 0:
                    continue
                self.visitConnectedLands(grid, i, j, visited)
                count = count+1
        return count

    def visitConnectedLands(self, grid, i, j, visited):
        if i < 0 or j < 0 or i >= len(grid) or j >= len(grid[i]) or int(grid[i][j]) == 0 or visited[i][j]:
            return

        visited[i][j] = True
        self.visitConnectedLands(grid, i-1, j, visited)
        self.visitConnectedLands(grid, i-1, j+1, visited)
        self.visitConnectedLands(grid, i, j+1, visited)
        self.visitConnectedLands(grid, i+1, j+1, visited)
        self.visitConnectedLands(grid, i+1, j, visited)
        self.visitConnectedLands(grid, i+1, j-1, visited)
        self.visitConnectedLands(grid, i, j-1, visited)
        self.visitConnectedLands(grid, i-1, j-1, visited)



#{
#  Driver Code Starts
import sys
sys.setrecursionlimit(10**6)

if __name__ == '__main__':
    T=int(input())
    for i in range(T):
        n, m = map(int, input().split())
        grid = []
        for _ in range(n):
            a = list(input().split())
            grid.append(a)
        obj = Solution()
        ans = obj.numIslands(grid)
        print(ans)

# } Driver Code Ends