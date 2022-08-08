#User function Template for python3

class Solution:
    def findOrder(self,dict, N, K):
        adj = {}
        for i in range(N-1):
            w1 = dict[i]
            w2 = dict[i+1]

            minLen = len(w1)
            if minLen > len(w2):
                minLen = len(w2)

            for j in range(minLen):
                v = ord(w1[j]) - ord('a')
                w = ord(w2[j]) - ord('a')
                if v not in adj:
                    adj[v] = []
                if w not in adj[v]:
                    adj[v].append(w)

        ans = ""
        visited = [False] * k
        for i in range(k):
            if visited[i]:
                continue
            self.dfs(i, adj, visited, ans)
        print(ans)
        return ans
    def dfs(self, src, adj, visited, ans):
        visited[src] = True

        for n in adj[src]:
            if visited[n]:
                continue
            self.dfs(n, adj, visited, ans)

        currLetter = chr(src + ord('a'))
        print(currLetter)
        ans += currLetter
        print(ans)


#{
#  Driver Code Starts
#Initial Template for Python 3

class sort_by_order:
    def __init__(self,s):
        self.priority = {}
        for i in range(len(s)):
            self.priority[s[i]] = i

    def transform(self,word):
        new_word = ''
        for c in word:
            new_word += chr( ord('a') + self.priority[c] )
        return new_word

    def sort_this_list(self,lst):
        lst.sort(key = self.transform)

if __name__ == '__main__':
#     t=int(input())
#     for _ in range(t):
#         line=input().strip().split()
#         n=int(line[0])
#         k=int(line[1])
#
#         alien_dict = [x for x in input().strip().split()]
#         duplicate_dict = alien_dict.copy()
    alien_dict = ["baa","abcd","abca","cab","cad"]
    n = 5
    k = 4
    ob=Solution()
    order = ob.findOrder(alien_dict,n,k)
    s = ""
    for i in range(k):
        s += chr(97+i)
    l = list(order)
    l.sort()
    l = "".join(l)
    if s != l:
        print(0)
    else:
        x = sort_by_order(order)
        x.sort_this_list(duplicate_dict)

        if duplicate_dict == alien_dict:
            print(1)
        else:
            print(0)


# } Driver Code Ends