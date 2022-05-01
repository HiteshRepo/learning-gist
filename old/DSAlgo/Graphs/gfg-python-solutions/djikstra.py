
    #Function to find the shortest distance of all the vertices
    #from the source vertex S.
    def dijkstra(self, V, adj, S):
        pathWts = []

        for v in range(V):

            wt = self.traverse(V, adj, v, S)
            if wt > -1:
                pathWts.append(wt)

        return pathWts

    def traverse(self, V, adj, src, dest):
        visited = [False] * V
        import queue
        q =  queue.PriorityQueue()
        q.put((0, src))
        lessWt = sys.maxsize

        while(not(q.empty())):
            removed = q.get()

            if removed[1] == dest and removed[0] < lessWt:
                lessWt = removed[0]

            if removed[0] > lessWt:
                break

            if visited[removed[1]]:
                continue
            visited[removed[1]] = True

            for edge in adj[removed[1]]:
                if visited[edge[0]]:
                    continue
                q.put((removed[0] + edge[1], edge[0]))

        if lessWt == sys.maxsize:
            return -1
        return lessWt

#{
#  Driver Code Starts
#Initial Template for Python 3
import atexit
import io
import sys


if __name__ == '__main__':
    test_cases = int(input())
    for cases in range(test_cases):
        V,E = map(int,input().strip().split())
        adj = [[] for i in range(V)]
        for i in range(E):
            u,v,w = map(int,input().strip().split())
            adj[u].append([v,w])
            adj[v].append([u,w])
        S=int(input())
        ob = Solution()

        res = ob.dijkstra(V,adj,S)
        for i in res:
            print(i,end=" ")
        print()
# } Driver Code Ends