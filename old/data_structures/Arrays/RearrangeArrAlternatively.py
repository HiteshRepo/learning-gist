class Solution:
    ##Complete this function
    #Function to rearrange  the array elements alternately.
    # def shift(self,arr,stop):
    #     i = len(arr)-1
    #     while i>stop:
    #         arr[i] = arr[i-1]
    #         i = i-1
    # def rearrange(self,arr, n): 
    #     i = 0
    #     while i<n:
    #         currMax = arr[-1]
    #         self.shift(arr, i)
    #         arr[i] = currMax
    #         i = i+2
    
    #store 2 no.s at single location
    #even indices - arr[i] = arr[i] + (arr[max-i] % max-e) * max-e
    #odd indices - arr[i] = arr[i] + (arr[min-i] % max-e) * max-e
    #max-e = maximum element + 1
    def rearrange(self,arr, n):
        maxE = arr[-1] + 1
        maxI = n-1
        minI = 0
        
        i = 0
        while i<n:
            if i%2 == 0:
                arr[i] = int(arr[i] + (arr[maxI] % maxE) * maxE)
                maxI = maxI-1
            else:
                arr[i] = int(arr[i] + (arr[minI] % maxE) * maxE)
                minI = minI+1
            
            i = i+1
        
        i = 0
        while i<n:
            arr[i] = int(arr[i]/maxE)
            i = i+1