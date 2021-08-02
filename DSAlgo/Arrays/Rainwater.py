class Solution:
    def trappingWater(self, arr,n):
        lmax = arr[0]
        rmax = arr[n-1]
        water = 0
        
        l = 0
        r = n-1
        while l<=r:
            if arr[l] > lmax:
                lmax = arr[l]
            if arr[r] > rmax:
                rmax = arr[r]
                
            if lmax < rmax:
                water = water + lmax - arr[l]
                l = l+1
            else:
                water = water + rmax - arr[r]
                r = r-1
        return water