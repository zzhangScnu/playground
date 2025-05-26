package tree

func maxDepthN(root *NTreeNode) int {
	var maxDepth int
	var traverse func(node *NTreeNode, depth int)
	traverse = func(node *NTreeNode, depth int) {
		if node == nil {
			return
		}
		depth++
		maxDepth = max(maxDepth, depth)
		for _, child := range node.Children {
			traverse(child, depth)
		}
		depth--
	}
	traverse(root, 0)
	return maxDepth
}

func maxDepthNII(root *NTreeNode) int {
	var traverse func(node *NTreeNode, depth int) int
	traverse = func(node *NTreeNode, depth int) int {
		if node == nil {
			return 0
		}
		var maxDepth int
		for _, child := range node.Children {
			maxDepth = max(maxDepth, traverse(child, depth))
		}
		return maxDepth + 1
	}
	return traverse(root, 0)
}
