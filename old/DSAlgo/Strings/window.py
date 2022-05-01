from collections import Counter


def minWindow(s, t):
    if not t or not s:
        return ""

    dict_t = Counter(t)
    ans = float("inf"), None, None
    required = len(dict_t)
    window_counts = {}
    filtered_s = []

    for i, ch in enumerate(s):
        if ch in dict_t:
            filtered_s.append((i, ch))

    l = 0
    r = 0
    formed = 0

    while r < len(filtered_s):

        ch = filtered_s[r][1]
        window_counts[ch] = window_counts.get(ch, 0) + 1

        if window_counts[ch] == dict_t[ch]:
            formed += 1

        while l <= r and formed == required:

            ch = filtered_s[l][1]

            end = filtered_s[r][0]
            start = filtered_s[l][0]

            if end - start + 1 < ans[0]:
                ans = (end - start + 1, start, end)

            window_counts[ch] -= 1
            if window_counts[ch] < dict_t[ch]:
                formed -= 1

            l += 1
        r += 1
    return "" if ans[0] == float('inf') else s[ans[1]: ans[2] + 1]


print(minWindow("ADOBECODEBANC", "ABC"))
