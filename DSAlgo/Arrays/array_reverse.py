
def reverse_array(input_arr):
    start = 0
    end = len(input_arr) - 1

    while start < end:
        input_arr[start], input_arr[end] = input_arr[end], input_arr[start]
        start += 1
        end -= 1

    return input_arr


def reverse_array_recursive(input_arr, start, end):

    if start >= end:
        return

    input_arr[start], input_arr[end] = input_arr[end], input_arr[start]
    reverse_array_recursive(input_arr, start+1, end-1)

    return input_arr


def reverse_array_slicing(input_arr):
    return input_arr[::-1]
