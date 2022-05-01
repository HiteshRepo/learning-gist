def array_cyclic_rotateonce(arr):
    from collections import deque
    arr1 = deque(arr)
    arr1.rotate(1)
    return list(arr1)


def rotate(arr):
    n = len(arr)
    x = arr[n - 1]

    for i in range(n - 1, 0, -1):
        arr[i] = arr[i - 1]

    arr[0] = x
    return arr


# print(array_cyclic_rotateonce(
#     [1, 2, 3, 4, 5]))

# print(rotate(
#     [1, 2, 3, 4, 5]))
