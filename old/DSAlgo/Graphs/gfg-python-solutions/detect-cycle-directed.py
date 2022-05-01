

class Solution:

    #Function to detect cycle in a directed graph.
    def isCyclic(self, V, adj):
        beingVisited = [False] * V
        visited = [False] * V

        for i in range(V):
            if visited[i]:
                continue
            if self.traverse(i, beingVisited, visited):
                return True
        return False

    def traverse(self, src, beingVisited, visited):
        beingVisited[src] = True

        for node in adj[src]:
            if beingVisited[node]:
                return True
            if not(visited[node]) and self.traverse(node, beingVisited, visited):
                return True

        beingVisited[src] = False
        visited[src] = True

#{
#  Driver Code Starts
#Initial Template for Python 3

import sys
sys.setrecursionlimit(10**6)

if __name__ == '__main__':
    t = int(input())
    for i in range(t):
        V,E = list(map(int, input().strip().split()))
        adj = [[] for i in range(V)]
        for i in range(E):
            a,b = map(int,input().strip().split())
            adj[a].append(b)
        ob = Solution()

        if ob.isCyclic(V, adj):
            print(1)
        else:
            print(0)
# } Driver Code Ends