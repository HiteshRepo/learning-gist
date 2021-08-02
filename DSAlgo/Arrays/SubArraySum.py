def subArraySum(self,arr, n, s):
    start = 0
    end  = 1
    
    sum = arr[0]
    while end <= n:
        while sum>s and start<end-1:
            sum = sum - arr[start]
            start = start+1
        if sum == s:
            return [start+1, end]
        if end < n:
            sum = sum+arr[end]
        end=end+1
               
    return [-1]