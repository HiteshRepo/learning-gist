
    #Function to return list containing vertices in Topological order.
    def topoSort(self, V, adj):
        visited = [False] * V
        stack = []
        topoOrder = []

        for node in range(V):
            if visited[node]:
                continue
            self.topologicalTraverse(node, visited, stack, adj)

        while len(stack) > 0:
            n = len(stack) - 1
            top = stack[n]
            topoOrder.append(top)
            stack = stack[:n]

        return topoOrder

    def topologicalTraverse(self, src, visited, stack, adj):
        visited[src] = True

        for nbr in adj[src]:
            if visited[nbr]:
                continue
            self.topologicalTraverse(nbr, visited, stack, adj)

        stack.append(src)

#{
#  Driver Code Starts
# Driver Program

import sys
sys.setrecursionlimit(10**6)

def check(graph, N, res):
    if N!=len(res):
        return False
    map=[0]*N
    for i in range(N):
        map[res[i]]=i
    for i in range(N):
        for v in graph[i]:
            if map[i] > map[v]:
                return False
    return True

if __name__=='__main__':
    t = int(input())
    for i in range(t):
        e,N = list(map(int, input().strip().split()))
        adj = [[] for i in range(N)]

        for i in range(e):
            u,v=map(int,input().split())
            adj[u].append(v)

        ob = Solution()

        res = ob.topoSort(N, adj)

        if check(adj, N, res):
            print(1)
        else:
            print(0)
# Contributed By: Harshit Sidhwa

# } Driver Code Ends