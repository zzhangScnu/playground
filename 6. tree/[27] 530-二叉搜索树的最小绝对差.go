package tree

import "math"

// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
//
// 差值是一个正数，其数值等于两值之差的绝对值。
//
// 示例 1：
//
// 输入：root = [4,2,6,1,3]
// 输出：1
//
// 示例 2：
//
// 输入：root = [1,0,48,null,null,12,49]
// 输出：1
//
// 提示：
//
// 树中节点的数目范围是 [2, 10⁴]
// 0 <= Node.val <= 10⁵
func getMinimumDifference(root *TreeNode) int {
	minimum := math.MaxInt
	var pre *TreeNode
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		if pre != nil && node.Val-pre.Val < minimum {
			minimum = node.Val - pre.Val
		}
		pre = node
		traverse(node.Right)
	}
	traverse(root)
	return minimum
}
