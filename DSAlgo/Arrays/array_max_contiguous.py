
import sys

seq = []


def array_max_contiguous(arr):
    global seq
    seq.append(arr[0])

    max_ending_here = max_so_far = arr[0]

    for i in range(1, len(arr)):

        if arr[i] < max_ending_here + arr[i]:
            max_ending_here = max_ending_here + arr[i]
            seq.append(arr[i])

        else:
            max_ending_here = arr[i]

        if max_ending_here > max_so_far:
            max_so_far = max_ending_here
            seq = [arr[i]]

    return max_so_far


# print(array_max_contiguous([1, 2, -3, 4, 5]))
# print(array_max_contiguous([-1, -2, -3, -4]))
print(array_max_contiguous([-2, 5, -1]))
print(seq)
