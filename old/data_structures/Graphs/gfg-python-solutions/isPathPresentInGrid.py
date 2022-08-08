

class Solution:

    #Function to find whether a path exists from the source to destination.
    def is_Possible(self, grid):
        visited = set()
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                if grid[i][j] == 1:
                    return self.traverse(grid, visited, i, j)
        return 0

    def traverse(self, grid, visited, i, j):

        if i < 0 or j < 0 or i >= len(grid) or j >= len(grid) or grid[i][j] == 0 or (i,j) in visited:
            return 0

        if grid[i][j] == 2:
            return 1

        visited.add((i, j))

        if self.traverse(grid, visited, i-1, j) == 1:
            return 1
        if self.traverse(grid, visited, i, j+1) == 1:
            return 1
        if self.traverse(grid, visited, i+1, j) == 1:
            return 1
        if self.traverse(grid, visited, i, j-1) == 1:
            return 1



#{
#  Driver Code Starts
if __name__ == '__main__':
    T=int(input())
    for i in range(T):
        n = int(input())
        grid = []
        for _ in range(n):
            a = list(map(int, input().split()))
            grid.append(a)
        obj = Solution()
        ans = obj.is_Possible(grid)
        if(ans):
            print("1")
        else:
            print("0")

# } Driver Code Ends