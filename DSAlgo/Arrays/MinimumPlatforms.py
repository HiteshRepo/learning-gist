class Solution:    
    #Function to find the minimum number of platforms required at the
    #railway station such that no train waits.
    def minimumPlatform(self,n,arr,dep):
        arr.sort()
        dep.sort()
        
        i = 1
        j = 0
        max_platforms = 1
        needed_platforms = 1
        while i<n:
            if arr[i] <= dep[j]:
                max_platforms = max_platforms + 1
                needed_platforms = max(max_platforms, needed_platforms)
                i = i+1
            else:
                max_platforms = max_platforms - 1
                j = j+1
        return needed_platforms