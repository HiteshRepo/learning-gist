class Solution:

	def zigZag(self,arr, n):
		rel = False
		i = 0
		while i < n-1:
		    if not(rel) and arr[i] > arr[i+1]:
		        arr[i], arr[i+1] = arr[i+1], arr[i]
		    elif rel and arr[i] < arr[i+1]:
		        arr[i], arr[i+1] = arr[i+1], arr[i]
		    i = i+1
		    rel = not(rel)