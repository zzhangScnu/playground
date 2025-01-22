package tree

import "math"

// 给定一个二叉树，找出其最小深度。
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
// 说明：叶子节点是指没有子节点的节点。
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：2
//
// 示例 2：
//
// 输入：root = [2,null,3,null,4,null,5,null,6]
// 输出：5
//
// 提示：
//
// 树中节点数的范围在 [0, 10⁵] 内
// -1000 <= Node.val <= 1000
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight, rightHeight := minDepth(root.Left), minDepth(root.Right)
	if root.Left == nil && root.Right != nil {
		return 1 + rightHeight
	}
	if root.Left != nil && root.Right == nil {
		return 1 + leftHeight
	}
	return 1 + int(math.Min(float64(leftHeight), float64(rightHeight)))
}

/**
跟二叉树的最大深度一样，通过后序遍历求根节点的高度来实现求深度。
其中，需要注意，
当一个节点的左子树为空时，不能取它的高度+1，因为题目要求的最小深度是【根节点到最近的叶子节点】的最短路径。
*/
