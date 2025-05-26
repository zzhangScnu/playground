package tree

func postorderTraversalN(root *NTreeNode) []int {
	var res []int
	var traverse func(node *NTreeNode)
	traverse = func(node *NTreeNode) {
		if node == nil {
			return
		}
		for _, child := range node.Children {
			traverse(child)
		}
		res = append(res, node.Val)
	}
	traverse(root)
	return res
}
