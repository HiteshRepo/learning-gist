class Solution:
    def MissingNumber(self,array,n):
        # for i in range(0,n):
        #     if i+1 not in array:
        #         return i+1
        xor = 0
        for i in array:
            xor = xor^i
        for i in range(1,n+1):
            xor = xor^i
        return xor