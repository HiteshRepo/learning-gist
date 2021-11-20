
class Solution:
    
    #Function to return a list containing the DFS traversal of the graph.
    def dfsOfGraph(self, V, adj):
        visited = [False] * V
        path = []
        
        self.dfsTraverse(0, visited, adj, path)
        
        return path
    
    def dfsTraverse(self, src, visited, adj, path):
        
        if visited[src]:
            return
        visited[src] = True
        
        path.append(src)
        
        for node in adj[src]:
            if visited[node]:
                continue
            self.dfsTraverse(node, visited, adj, path)

#{ 
#  Driver Code Starts

if __name__ == '__main__':
    T=int(input())
    while T>0:
        V,E=map(int,input().split())
        adj=[[] for i in range(V+1)]
        for i in range(E):
            u,v=map(int,input().split())
            adj[u].append(v)
            adj[v].append(u)
        ob=Solution()
        ans=ob.dfsOfGraph(V,adj)
        for i in range(len(ans)):
            print(ans[i],end=" ")
        print()
        T-=1
# } Driver Code Ends