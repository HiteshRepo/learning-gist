
def getMinMax(inp_arr):

    n = len(inp_arr)
    min = None
    max = None

    if n == 1:
        min = max = inp_arr[0]
        return (min, max)

    elif n > 1:
        if inp_arr[0] > inp_arr[1]:
            max = inp_arr[0]
            min = inp_arr[1]
        else:
            min = inp_arr[0]
            max = inp_arr[1]

        for i in range(2, n):
            if inp_arr[i] > max:
                max = inp_arr[i]
            elif inp_arr[i] < min:
                min = inp_arr[i]

        return (min, max)


def getMinMaxRecursive(low_pos, high_pos, inp_arr):
    max = inp_arr[low_pos]
    min = inp_arr[low_pos]

    if low_pos == high_pos:
        return (inp_arr[low_pos], inp_arr[low_pos])

    elif low_pos == high_pos - 1:
        if inp_arr[low_pos] > inp_arr[high_pos]:
            return (inp_arr[high_pos], inp_arr[low_pos])
        else:
            return (inp_arr[low_pos], inp_arr[high_pos])

    else:
        mid = int((low_pos + high_pos) / 2)
        min1, max1 = getMinMaxRecursive(low_pos, mid, inp_arr)
        min2, max2 = getMinMaxRecursive(mid + 1, high_pos, inp_arr)

    if min1 < min2:
        min = min1
    else:
        min = min2

    if max1 > max2:
        max = max1
    else:
        max = max2

    return (min, max)
