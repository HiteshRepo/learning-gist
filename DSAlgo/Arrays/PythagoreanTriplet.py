class Solution:

	def checkTriplet(self,arr, n):
    	i = n-1
    	arr.sort()
    	while i >= 2:
    	    
    	    j = 0
    	    k = i-1
    	    while j<k:
    	        if arr[j]*arr[j] + arr[k]*arr[k] == arr[i]*arr[i]:
    	            return True
    	        if arr[j]*arr[j] + arr[k]*arr[k] > arr[i]*arr[i]:
    	            k = k-1
    	        else:
    	            j = j+1
    	    
    	    i = i-1
    	return False