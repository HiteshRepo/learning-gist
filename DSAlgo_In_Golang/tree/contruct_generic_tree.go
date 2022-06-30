package tree

func ConstructGenericTree(values []int) *Tree {
	st := NewTreeNodeStack()

	var tree *Tree
	for _,v := range values {
		if v == -1 && st.Len() > 0 {
			st.Pop()
		} else {
			node := GetNewTreeNode(v)

			if st.Len() == 0 {
				tree = GetNewTree(node)
			} else {
				parent := st.Peek()
				parent.Add(node)
			}

			st.Push(node)
		}
	}

	return tree
}