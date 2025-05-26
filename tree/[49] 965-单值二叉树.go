package tree

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	uniqueVal := root.Val
	isUnique := true
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val != uniqueVal {
			isUnique = false
			return
		}
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(root)
	return isUnique
}
