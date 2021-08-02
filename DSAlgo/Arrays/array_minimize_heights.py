
def minimize_heights(arr, k):

    if len(arr) == 1:
        return 0

    arr.sort()
    n = len(arr)
    difference = arr[n-1] - arr[0]
    min = arr[0] + k
    max = arr[n-1] - k

    if min > max:
        min, max = max, min

    for i in range(0, n):
        diff = arr[i] - k
        sum = arr[i] + k

        if(diff >= min or sum <= max):
            continue
        if(max-diff <= sum-min):
            min = diff
        else:
            max = sum

    if difference > max - min:
        return max - min
    else:
        return difference
