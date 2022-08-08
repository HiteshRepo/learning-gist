from collections import deque

row = [-1, 0, 0, 1]
col = [0, -1, 1, 0]

class Solution:
    def shortestDistance(self,N,M,A,X,Y):
        path = self.findPath(A, X, Y, N, M)
        if path:
            return len(path) - 1
        return -1


    def findPath(self, matrix, X, Y, N, M, x=0, y=0):
        if not matrix or not len(matrix):
            return

        q = deque()
        src = Node(x, y)
        q.append(src)

        visited = set()

        key = (src.x, src.y)
        visited.add(key)

        while q:

            curr = q.popleft()
            i = curr.x
            j = curr.y

            if i == X and j == Y:
                path = []
                getPath(curr, path)
                return path

            n = matrix[i][j]

            for k in range(len(row)):
                x = i + row[k] * n
                y = j + col[k] * n

                validPath = isValid(x, y, N-1, M-1, matrix)
                if validPath:
                    next = Node(x, y, curr)
                    key = (next.x, next.y)

                    if key not in visited:
                        q.append(next)
                        visited.add(key)

        return
class Node:
    def __init__(self, x, y, parent=None):
        self.x = x
        self.y = y
        self.parent = parent

    def __repr__(self):
        return str((self.x, self.y))

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

def isValid(x, y, X, Y, matrix):
    return (0 <= x <= X) and (0 <= y <= Y) and matrix[x][y] == 1

def getPath(node, path=[]):
    if node:
        getPath(node.parent, path)
        path.append(node)

#{
#  Driver Code Starts
#Initial Template for Python 3

import math

if __name__=='__main__':
    t=int(input())
    for _ in range(t):
        N,M=map(int,input().strip().split())
        a=[]
        for i in range(N):
            s=list(map(int,input().strip().split()))
            a.append(s)
        x,y=map(int,input().strip().split())
        ob=Solution()
        print(ob.shortestDistance(N,M,a,x,y))

# } Driver Code Ends