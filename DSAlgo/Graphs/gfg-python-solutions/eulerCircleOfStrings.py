#User function Template for python3

ansMap = {False: 0, True: 1}
class Solution:
    def isCircle(self, N, A):
        inp = [0] * 26
        out = [0] * 26
        visited = [False] * 26
        adj = [[] for i in range(26)]

        for s in A:
            first = ord(s[0].lower()) - ord('a')
            last = ord(s[-1].lower()) - ord('a')

            inp[last] += 1
            out[first] += 1
            adj[first].append(last)

        for i in range(26):
            if inp[i] != out[i]:
                return 0

        src = -1
        for i in range(26):
            if out[i]:
                src = i
                break

        if src != -1:
            self.traverse(adj, visited, src)
            for i in range(26):
                if out[i] and not visited[i]:
                    return 0
        else:
            return 0

        return 1

    def traverse(self, adj, visited, src):
        visited[src] = True
        for v in adj[src]:
            if visited[v]:
                continue
            self.traverse(adj, visited, v)

class IncorrectSolution:
    def isCircle(self, N, A):
            adj = {}

            for s in A:
                first = s[0]
                last = s[-1]

                if first in adj:
                    adj[first].append(last)
                else:
                    adj[first] = [last]

            visited = {}
            allPaths = []
            self.traverse(adj, visited, allPaths)
            return ansMap[len(allPaths) == 1 and allPaths[0][0] in  adj[allPaths[0][-1]]]

    def traverse(self, adj, visited, allPaths):
        for v in adj.keys():
            if v in visited:
                continue
            path = []
            self.connected(adj, visited, v, path)
            allPaths.append(path)

    def connected(self, adj, visited, src, path):

        visited[src] = True
        path.append(src)

        if src in adj:
            for v in adj[src]:
                if v in visited:
                    continue
                self.connected(adj, visited, v, path)


#{
#  Driver Code Starts
#Initial Template for Python 3

import sys
sys.setrecursionlimit(10**6)
if __name__ == '__main__':
#     t = int(input())
#     for _ in range (t):
#         N = int(input())
#         A = input().split()
#     N = 3
#     A = ["abc", "bcd", "cdf"]
#     N = 4
#     A = ["ab" , "bc", "cd", "da"]
#     N = 1
#     A = ["odbmrmayzr"]
    N = 4
    A = ["for", "geek", "rig", "kaf"]
    ob = Solution()
    print(ob.isCircle(N, A))
# } Driver Code Ends