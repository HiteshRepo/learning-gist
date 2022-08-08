import sys


def decideJumps(arr, n, currPos):
    if currPos >= n-1:
        return 0

    minJump = sys.maxsize
    maxSteps = arr[currPos]

    while(maxSteps > 0):
        if minJump > 1 + decideJumps(arr, n, currPos + maxSteps):
            minJump = 1 + decideJumps(arr, n, currPos + maxSteps)
        maxSteps -= 1

    return minJump


def minJump(arr, n):

    min_jumps_array = 0
    MJA = [0]*n
    i = 1
    while(i < n and i > min_jumps_array):
        if min_jumps_array+arr[min_jumps_array] >= i:
            MJA[i] = MJA[min_jumps_array]+1
            i += 1
        else:
            min_jumps_array += 1
    if MJA[-1] == 0:
        return -1
    else:
        print(MJA)
        return MJA[-1]


def jump(arr):
    #jumps = decideJumps(arr, len(arr), 0)
    # if jumps == sys.maxsize:
    # return -1
    return minJump(arr, len(arr))


# tests = int(input())

# while(tests > 0):
#     tests -= 1

#     n = int(input())

#     arr = list(map(int, input().split()))

#     jumps = jump(arr)

#     print(jumps)

print(jump([1, 3, 5, 8, 9, 2, 6, 7, 6, 8, 9]))
# print(jump([1, 4, 3, 2, 6, 7]))
# print(jump([2, 3, 1, 1, 2, 4, 2, 0, 1, 1]))
# print(jump([0, 1, 1, 1, 1]))

# arr = '32 54 12 52 56 8 30 44 94 44 39 65 19 51 91 1 5 89 34 25 58 20 51 38 65 30 7 20 10 51 18 43 71 97 61 26 5 57 70 65 0 75 29 86 93 87 87 64 75 88 89 100 7 40 37 38 36 44 24 46 95 43 89 32 5 15 58 77 72 95 8 38 69 37 24 27 90 77 92 31 30 80 30 37'
# print(jump(list(map(int, arr.split()))))
