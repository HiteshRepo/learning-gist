package tree

func FindLowestCommonAncestor(val1, val2 int, tree *Tree) int {
	path1 := tree.NodeToRoot(val1)
	path2 := tree.NodeToRoot(val2)

	i := len(path1) - 1
	j := len(path2) - 1

	for i >= 0 && j >= 0 {
		if path1[i] != path2[j] {
			break
		}

		i -= 1
		j -= 1
	}

	return path1[i+1]
}
