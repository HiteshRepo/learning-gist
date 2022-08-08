str1 = "2 5 7 8"
arr = list(map(int, str1.split()))


def meetTarget(target):
    global arr

    for i in range(0, len(arr)):
        for j in range(i+1, len(arr)):
            if arr[i] < arr[j] and arr[i] + arr[j] == target:
                return [arr[i], arr[j]]

    return [None, None]


print(meetTarget(9))
