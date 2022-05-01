def rob(nums, start):
    if start >= len(nums):
        return 0

    res = max(rob(nums, start + 1), nums[start] + rob(nums, start + 2))
    return res


memo = {}


def rob_dp(nums, start):

    if start >= len(nums):
        return 0

    if start in memo:
        return memo[start]

    res = max(rob_dp(nums, start + 1), nums[start] + rob_dp(nums, start + 2))
    memo[start] = res
    return res


def rob_dp_btm_up(nums):
    n = len(nums)

    memo = [0] * (n+2)

    for i in range(n-1, -1, -1):
        memo[i] = max(memo[i+1], nums[i] + memo[i+2])

    return memo[0]


def rob_dp_btm_up_o1(nums):
    n = len(nums)

    f = 0
    s = 0
    t = 0

    for i in range(n-1, -1, -1):
        t = max(f, nums[i] + s)
        s = f
        f = t

    return t


# print(rob([1, 2, 3, 1], 0))
# print(rob([2, 7, 9, 3, 1], 0))

# print(rob_dp([2, 7, 9, 3, 1], 0))

# print(rob_dp_btm_up([2, 7, 9, 3, 1]))

print(rob_dp_btm_up_o1([2, 7, 9, 3, 1]))
