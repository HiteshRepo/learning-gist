def findMaxIndex(index, a, curr):

    ans = -1
    idx = 0

    for i in range(index, len(a)):

        if a[i] > curr:
            idx = i
            if ans == -1:
                ans = curr
            else:
                if ans > a[i]:
                    ans = a[i]
    return idx


def next_perm(nums):

    found = False
    i = len(nums) - 2

    while i >= 0:
        if nums[i] < nums[i+1]:
            found = True
            break
        i -= 1

    if not found:
        nums.sort()
    else:
        m = findMaxIndex(i+1, nums, nums[i])
        nums[i], nums[m] = nums[m], nums[i]
        nums[i+1:] = nums[i+1:][::-1]
    return nums


print(next_perm([1, 2, 5, 4, 3]))
