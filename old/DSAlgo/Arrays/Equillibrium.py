class Solution:
    # Complete this function
    
    #Function to find equilibrium point in the array.
    def equilibriumPoint(self,A, N):
        if N == 1:
            return N
        
        if N == 0:
            return -1
        
        point = -1
        indices = []
        left = [0]*N
        right = 0
        
        i=1
        while i<N:
            left[i] = left[i-1] + A[i-1]
            i = i+1
        
        i=N-1
        while i>0:
            
            if left[i] == right:
                indices.append(i)
            
            right = right + A[i]
            
            i = i-1
        
        if len(indices) > 0:
            point = indices[-1] + 1
        return point