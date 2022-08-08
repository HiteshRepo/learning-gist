dict_node_val = {}
def rob(node):

    if node == null:
        return 0
    
    if node in dict_node_val.keys():
        return dict_node_val[node]
    
    rob_it = node.val + rob(node.left.left) + rob(node.left.right) + rob(node.right.left) + rob(node.right.right) 

    not_rob = rob(node.left) + rob(node.right)

    res = max(rob_it, not_rob)

    dict_node_val[node] = res

    return res