
def merge(intervals):

    if len(intervals) == 1:
        return intervals

    intervals.sort(key=lambda x: x[0])

    s = -10000
    max = -100000

    m = []

    for i in range(len(intervals)):
        a = intervals[i]
        if a[0] > max:
            if i != 0:
                m.append([s, max])
            max = a[1]
            s = a[0]
        else:
            if a[1] > max:
                max = a[1]

    if max != -100000 and [s, max] not in m:
        m.append([s, max])

    print(m)


merge([[1, 3], [2, 6], [8, 10], [15, 18]])
# merge([[1, 4], [4, 5]])
# merge([[1, 4], [5, 6]])
