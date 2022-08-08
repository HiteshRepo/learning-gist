package tree

func AreTwoTreeSimilarInShape(tree1, tree2 *Tree) bool {
	return areSimilar(tree1.Root, tree2.Root)
}

func areSimilar(n1, n2 *Node) bool {
	if len(n1.Children) != len(n2.Children) {
		return false
	}

	for i:=0; i<len(n1.Children); i++ {
		if !areSimilar(n1.Children[i], n2.Children[i]) {
			return false
		}
	}

	return true
}
