# def merge(arr1, arr2, n, m):

#     for i in range(0, n):

#         if arr1[i] > arr2[0]:
#             arr1[i], arr2[0] = arr2[0], arr1[i]
# arr2.sort()

# also a sorting technique
# first = arr2[0]
# k = 1
# while k < m and arr2[k] < first:
#     arr2[k - 1] = arr2[k]
#     k = k + 1
# arr2[k - 1] = first

# return [arr1, arr2]
def nextGap(gap):

    if (gap <= 1):
        return 0
    return (gap // 2) + (gap % 2)


def merge(arr1, arr2, n, m):

    gap = n + m
    gap = nextGap(gap)
    while gap > 0:

        # comparing elements in
        # the first array.
        i = 0
        while i + gap < n:
            if (arr1[i] > arr1[i + gap]):
                arr1[i], arr1[i + gap] = arr1[i + gap], arr1[i]

            i += 1

        # comparing elements in both arrays.
        j = gap - n if gap > n else 0
        while i < n and j < m:
            if (arr1[i] > arr2[j]):
                arr1[i], arr2[j] = arr2[j], arr1[i]

            i += 1
            j += 1

        if (j < m):

            # comparing elements in the
            # second array.
            j = 0
            while j + gap < m:
                if (arr2[j] > arr2[j + gap]):
                    arr2[j], arr2[j + gap] = arr2[j + gap], arr2[j]

                j += 1

        gap = nextGap(gap)
    return [arr1, arr2]


# merge([1, 3, 5, 7], [0, 2, 6, 8, 9], 4, 5)

for a in merge([1, 3, 5, 7], [0, 2, 6, 8, 9], 4, 5):
    for i in a:
        print(i, end=" ")
