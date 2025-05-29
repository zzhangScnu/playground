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

/**
这里需要额外维护一个isUnique变量，如果找到不同值的节点/遍历到树的叶子节点，则递归返回至主函数。
为什么不能直接给递归函数traverse增加返回值，用其表示是否单值二叉树呢？
——因为在node == nil时，无论返回true还是false，都不合适，无法真正表示二叉树是否满足要求。
*/
