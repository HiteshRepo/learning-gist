#User function Template for python3

class Solution:

    #Function to return Breadth First Traversal of given graph.
    def bfsOfGraph(self, V, adj):
        visited = [False] * V
        path = []

        self.bfsTraverse(adj, visited, path)

        return path

    def bfsTraverse(self, adj, visited, path):
        q = []
        q.append(0)

        while len(q) > 0:
            removed = q.pop(0)

            if visited[removed]:
                continue
            else:
                visited[removed] = True

            path.append(removed)

            for node in adj[removed]:
                if visited[node]:
                    continue
                q.append(node)




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
		ob = Solution()
		ans = ob.bfsOfGraph(V, adj)
		for i in range(len(ans)):
		    print(ans[i], end = " ")
		print()


# } Driver Code Ends