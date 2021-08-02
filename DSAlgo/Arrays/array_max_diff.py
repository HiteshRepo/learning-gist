# arr = [2, 3, 10, 6, 4, 8, 1]
arr = [7, 9, 5, 6, 3, 2]


def maxDiff():
    global arr

    maxDif = arr[1] - arr[0]

    for i in range(0, len(arr)):
        for j in range(i+1, len(arr)):
            if arr[j] - arr[i] > maxDif:
                maxDif = arr[j] - arr[i]
    return maxDif


print(maxDiff())


def maxDiff_optimized():
    global arr

    maxDif = arr[1] - arr[0]
    minEl = arr[0]

    for i in range(1, len(arr)):
        if arr[i] - minEl > maxDif:
            maxDif = arr[i] - minEl

        if minEl > arr[i]:
            minEl = arr[i]

    return maxDif


print(maxDiff_optimized())
