class Solution:
    def lastIndex(self, s):
        arr = list(s)
        n = len(arr)
        i = 0
        max1 = -1
        while i < n:
            if arr[i] == '1':
                max1 = i
            i = i+1
        return max1