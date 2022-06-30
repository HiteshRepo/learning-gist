package tree

func IsTreeSymmetric(tree *Tree) bool {
	return AreTwoTreeMirror(tree, tree)
}
