class Solution:	
    #Function to reverse every sub-array group of size k.
	def reverseInGroups(self, arr, N, K):
		i = 0
		while i < N:
		    till = i+K
		    if till > N:
		        till = N
		    j = till-1
		    k = i
		    while j > k:
		        arr[k], arr[j] = arr[j], arr[k]
		        j = j-1
		        k = k+1
		    i = i + K