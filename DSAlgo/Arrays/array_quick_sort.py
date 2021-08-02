def partition(a, lb, ub):
    pivot = a[lb]
    start = lb
    end = ub

    while(start < end):
        while a[start] <= pivot:
            start += 1

        while a[end] > pivot:
            end -= 1

        if start < end:
            a[start], a[end] = a[end], a[start]

    a[lb], a[end] = a[end], a[lb]

    return end


def quicksort(arr, lb, ub):
    if lb < ub:
        loc = partition(arr, lb, ub)
        quicksort(arr, lb, loc-1)
        quicksort(arr, loc + 1, ub)


arr = [7, 10, 4, 3, 20, 15]
quicksort(arr, 0, 5)
print(arr)
