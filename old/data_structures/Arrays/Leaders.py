class Solution:
    #Back-end complete function Template for Python 3
    
    #Function to find the leaders in the array.
    def leaders(self, A, N):
        i = N-2
        leaders = [A[-1]]
        max = A[-1]
        while i >= 0:
            if A[i] >= max:
                leaders.append(A[i])
                max = A[i]
            i = i-1
        leaders.reverse()
        return leaders