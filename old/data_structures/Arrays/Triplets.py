def countTriplet(self, arr, n):
    arr.sort()
    count=0
    i = n-1
    while i > 0:
        l=0
        r=i
        while l<r:
            if arr[l]+arr[r] == arr[i]:
                count = count+1
            if arr[l]+arr[r] > arr[i]:
                r = r-1
            else:
                l = l+1
        i = i-1
    return count