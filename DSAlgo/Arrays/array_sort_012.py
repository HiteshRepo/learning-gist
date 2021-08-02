def sort012(a):
    hi = len(a) - 1
    lo = mid = 0

    while(mid <= hi):
        if a[mid] == 0:
            a[mid], a[lo] = a[lo], a[mid]
            mid += 1
            lo += 1

        elif a[mid] == 1:
            mid += 1

        elif a[mid] == 2:
            a[mid], a[hi] = a[hi], a[mid]
            hi -= 1

    return a


def sortCount012(arr):
    cnt0 = 0
    cnt1 = 0
    cnt2 = 0

    for i in range(0, len(arr)):
        if arr[i] == 0:
            cnt0 += 1
        if arr[i] == 1:
            cnt1 += 1
        if arr[i] == 2:
            cnt2 += 1

    i = 0

    while cnt0 != 0:
        arr[i] = 0
        i += 1
        cnt0 -= 1

    while cnt1 != 0:
        arr[i] = 1
        i += 1
        cnt1 -= 1

    while cnt2 != 0:
        arr[i] = 2
        i += 1
        cnt2 -= 1

    return arr


# print(sort012([1, 0, 2, 0, 1, 2, 0, 0, 0, 2, 1]))
