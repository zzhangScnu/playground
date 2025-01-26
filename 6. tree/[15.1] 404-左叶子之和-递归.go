package tree

// 给定二叉树的根节点 root ，返回所有左叶子之和。
//
// 示例 1：
//
// 输入: root = [3,9,20,null,null,15,7]
// 输出: 24
// 解释: 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
//
// 示例 2:
//
// 输入: root = [1]
// 输出: 0
//
// 提示:
//
// 节点数在 [1, 1000] 范围内
// -1000 <= Node.val <= 1000
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left int
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		left = root.Left.Val
	} else {
		left = sumOfLeftLeaves(root.Left)
	}
	return left + sumOfLeftLeaves(root.Right)
}

/**
某一层不一定要做什么特别的事，有可能只是收集左右孩子的结果，并返回给上层而已。
这题也是用的后序遍历。
*/
