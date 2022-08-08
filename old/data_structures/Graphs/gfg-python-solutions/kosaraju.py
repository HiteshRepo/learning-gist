
class Solution:

    #Function to find number of strongly connected components in the graph.
    def kosaraju(self, V, adj):
        visited = [False] * V
        stack = []
        ans = 0

        for i in range(V):
            if visited[i]:
                continue
            self.traverse(adj, visited, i, stack, False)
            stack.append(i)

        #reverse the edges of graph
        rev = self.reverse(adj)

        visited = [False] * V

        while(len(stack) > 0):
            n = len(stack) - 1
            last = stack[n]
            stack = stack[:n]

            if visited[last]:
                continue

            self.traverse(rev, visited, last, stack, True)
            ans += 1

        return ans


    def traverse(self, adj, visited, src, stack, reverse):

        visited[src] = True

        for nbr in adj[src]:
            if visited[nbr]:
                continue
            self.traverse(adj, visited, nbr, stack, reverse)
            if not(reverse):
                stack.append(nbr)

    def reverse(self, adj):
        rev = [[] for i in range(len(adj))]
        for i in range(len(adj)):
            for nbr in adj[i]:
                rev[nbr].append(i)

        return rev




#{
#  Driver Code Starts
#Initial Template for Python 3

from collections import OrderedDict
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

        print(ob.kosaraju(V, adj))
# } Driver Code Ends