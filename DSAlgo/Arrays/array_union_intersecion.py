def array_union_intersecion(arr1, arr2):

    set1 = set(arr1 + arr2)
    u_i = []

    for i in arr1:
        if i in arr2 and i not in u_i:
            u_i.append(i)

    union = len(set1)
    intersection = len(u_i)
    ui = [union, intersection]
    return ui


# print(array_union_intersecion(
#     [1, 2, 3, 4, 5], [1, 2, 3, 4]))
