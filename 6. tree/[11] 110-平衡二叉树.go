package tree

import "math"

// 给定一个二叉树，判断它是否是 平衡二叉树
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：true
//
// 示例 2：
//
// 输入：root = [1,2,2,3,3,null,null,4,4]
// 输出：false
//
// 示例 3：
//
// 输入：root = []
// 输出：true
//
// 提示：
//
// 树中的节点数在范围 [0, 5000] 内
// -10⁴ <= Node.val <= 10⁴
func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := getHeight(node.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := getHeight(node.Right)
	if rightHeight == -1 {
		return -1
	}
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return -1
	}
	return 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
}
