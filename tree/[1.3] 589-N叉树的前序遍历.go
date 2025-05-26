package tree

func preorderTraversalN(root *NTreeNode) []int {
	var res []int
	var traverse func(node *NTreeNode)
	traverse = func(node *NTreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for _, child := range node.Children {
			traverse(child)
		}
	}
	traverse(root)
	return res
}
