import sys
class Solution:
    ##Complete this function
    #Function to find the sum of contiguous subarray with maximum sum.
    def maxSubArraySum(self,a,size):
        currSum = 0
        globalSum = -sys.maxsize - 1
        i = 0
        while i<size:
            if a[i] > a[i]+currSum:
                currSum = a[i]
            else:
                currSum = a[i]+currSum
            if currSum > globalSum:
                globalSum = currSum
            i = i+1
        return globalSum
