import sys
class Solution:

    def findMinDiff(self, A,N,M):
        A.sort()
        i = 0
        j = M-1
        diff = sys.maxsize
        while j < N:
            if A[j] - A[i] < diff:
                diff = A[j] - A[i]
            i = i+1
            j = j+1
        return diff