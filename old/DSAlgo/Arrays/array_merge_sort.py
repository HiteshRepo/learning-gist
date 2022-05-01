merged_arr = []


def merge(arr, lb, mid, ub):

    i = lb
    j = mid + 1

    while(i <= mid and j <= ub):
        if(arr[i] <= arr[j]):
            merged_arr.append(arr[i])
            i += 1
        else:
            merged_arr.append(arr[j])
            j += 1

    if i > mid:
        while j <= ub:
            merged_arr.append(arr[j])
            j += 1
    else:
        while i <= mid:
            merged_arr.append(arr[i])
            i += 1


def mergesort(arr, lb, ub):
    if lb < ub:

        mid = (lb + (ub - 1)) // 2

        mergesort(arr, lb, mid)
        mergesort(arr, mid + 1, ub)
        merge(arr, lb, mid, ub)


def sort(arr):
    mergesort(arr, 0, len(arr) - 1)
    print(merged_arr)
    return merged_arr[-(len(arr)):]


print(sort(list(map(int, '2 4 1 3 5'.split()))))
