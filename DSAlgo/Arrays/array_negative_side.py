def array_negative_side(arr):

    pos = 0
    for i in range(0, len(arr)):
        if arr[i] < 0:
            arr[i], arr[pos] = arr[pos], arr[i]
            pos += 1

    return arr


# print(array_negative_side([1, -1, 2, 3, -9]))
