package tree

func AreTwoTreeMirror(tree1, tree2 *Tree) bool {
	return areMirror(tree1.Root, tree2.Root)
}

func areMirror(n1, n2 *Node) bool {
	if len(n1.Children) != len(n2.Children) {
		return false
	}

	for i:=0; i<len(n1.Children); i++ {
		j := len(n2.Children) - i - 1
		if !areMirror(n1.Children[i], n2.Children[j]) {
			return false
		}
	}

	return true
}