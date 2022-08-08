
    #Function to detect cycle in an undirected graph.
    def isCycle(self, V, adj):
        visited = [False] * V

        for i in range(V):
            if visited[i]:
                continue
            if self.bfsTraverse(adj, visited, i):
                return 1
        return 0

    def bfsTraverse(self, adj, visited, i):
        q = []
        q.append(i)

        while len(q) > 0:
            removed = q.pop(0)

            if visited[removed]:
                return True
            visited[removed] = True

            for node in adj[removed]:
                if visited[node]:
                    continue
                q.append(node)

        return False

#{
#  Driver Code Starts
if __name__ == '__main__':

    T=int(input())
    for i in range(T):
        V, E = map(int, input().split())
        adj = [[] for i in range(V)]
        for _ in range(E):
            u, v = map(int, input().split())
            adj[u].append(v)
            adj[v].append(u)
        obj = Solution()
        ans = obj.isCycle(V, adj)
        if(ans):
            print("1")
        else:
            print("0")

# } Driver Code Ends