import sys
def findElement( arr, n):
    leftMax = [0]*n
    leftMax[0] = -sys.maxsize-1
    pos = -1
    
    i = 1
    while i<n:
        leftMax[i] = max(leftMax[i-1], arr[i-1])
        i = i+1
    rightMin = sys.maxsize
    i = n-1
    while i>=0:
        if leftMax[i] <= arr[i] and rightMin >= arr[i] and i != n-1 and i != 0:
            pos = i
        rightMin = min(rightMin, arr[i])
        i = i-1
    if pos != -1 and arr[pos] <= arr[pos+1]:
        return arr[pos]
        
    return -1