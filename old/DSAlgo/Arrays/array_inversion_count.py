# def inversion_count(arr):

#     cnt = 0
#     for i in range(len(arr)):
#         for j in range(i + 1, len(arr)):
#             if arr[i] > arr[j]:
#                 cnt += 1

#     return cnt
merged_arr = []


def merge(arr, lb, mid, ub):

    i = lb
    j = mid + 1
    cnt = 0

    while(i <= mid and j <= ub):
        if(arr[i] <= arr[j]):
            merged_arr.append(arr[i])
            i += 1
        else:
            merged_arr.append(arr[j])
            j += 1
            cnt += mid-i+1

    if i > mid:
        while j <= ub:
            merged_arr.append(arr[j])
            j += 1
    else:
        while i <= mid:
            merged_arr.append(arr[i])
            i += 1

    return cnt


def mergesort(arr, lb, ub):
    if lb < ub:

        mid = (lb + (ub - 1)) // 2
        inversion_count = 0

        inversion_count += mergesort(arr, lb, mid)
        inversion_count += mergesort(arr, mid + 1, ub)
        inversion_count += merge(arr, lb, mid, ub)

        return inversion_count

    else:
        return 0


def sort(arr):
    return mergesort(arr, 0, len(arr) - 1)


print(sort(list(map(int, '2 4 1 3 5'.split()))))
