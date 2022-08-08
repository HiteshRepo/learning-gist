from functools import cmp_to_key
class Solution:
    def compare(self,x,y):
        if x+y<y+x:
            return 1
        elif x+y == y+x:
            return 0
        else:
            return -1
	def printLargest(self,arr):
	   # i = 0
	   # while i<len(arr):
	   #     j = i+1
	   #     while j<len(arr):
	   #         ij = str(arr[i]) + str(arr[j])
	   #         ji = str(arr[j]) + str(arr[i])
	   #         if int(ji) > int(ij):
	   #             arr[i], arr[j] = arr[j], arr[i]
	   #         j = j+1
	   #     i = i+1
	   # i = 0
	   # num = ''
	   # while i < len(arr):
	   #     num += str(arr[i])
	   #     i = i+1
	   # return num
	   for i in range(len(arr)):
	       arr[i] = str(arr[i])
       arr.sort(key=cmp_to_key(lambda x,y:self.compare(x,y)))
       return "".join(arr).lstrip("0") or "0"