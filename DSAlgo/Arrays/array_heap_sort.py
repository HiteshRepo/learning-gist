def maxHeapify(A, n, i):
    largest = i
    left = (2*i) + 1
    right = (2*i) + 2

    while left <= n and A[left] > A[largest]:
        largest = left

    while right <= n and A[right] > A[largest]:
        largest = right

    if largest != i:
        A[i], A[largest] = A[largest], A[i]
        # print(A)
        maxHeapify(A, n, largest)


def heapsort(A):
    n = len(A)

    for i in range(int(n/2), 0, -1):
        maxHeapify(A, n-1, i-1)

    for i in range(n - 1, 0, -1):
        # print('start')
        A[0], A[i] = A[i], A[0]
        # print(A)
        maxHeapify(A, i-1, 0)


# A = [5, 3, 9, 1, 6]
A = [9, 8, 3, 7, 2, 6, 11, 85, 22]
heapsort(A)
print(A)
