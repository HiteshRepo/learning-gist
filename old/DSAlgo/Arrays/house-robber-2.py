def rob_dp_btm_up_o1_circular(nums, start, end):
    n = len(nums)

    f = 0
    s = 0
    t = 0

    for i in range(end, start-1, -1):
        t = max(f, nums[i] + s)
        s = f
        f = t

    return t


def circular_steal(nums):
    n = len(nums)
    if n == 0:
        return 0
    if n == 1:
        return nums[0]

    return max(rob_dp_btm_up_o1_circular(nums, 0, n-2), rob_dp_btm_up_o1_circular(nums, 1, n-1))


# print(circular_steal([2, 3, 2])) #3
print(circular_steal([1, 2, 3, 1]))  # 4
