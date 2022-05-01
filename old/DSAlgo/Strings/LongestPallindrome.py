class Solution:
    def longestPalin(self, S):
        dp = [[0 for c in S] for c in S]
        i = 0
        maxL = 0
        s = []
        idx = -1
        while i<len(S):
            j = 0
            k = i
            while k<len(S):
                if i == 0:
                    dp[j][k] = 1
                elif i == 1:
                    if S[j] == S[k]:
                        dp[j][k] = 1
                else:
                    if S[j] == S[k] and dp[j+1][k-1] == 1:
                        dp[j][k] = 1
                if dp[j][k] == 1:
                    maxL = i+1
                    c = S[j:k+1]
                    if len(s) > 0 and len(c) > len(s[-1]):
                        idx = len(s)
                    elif len(s) == 0:
                        idx = 0
                    s.append(c)
                j = j+1
                k = k+1
            i = i+1
        return s[idx]