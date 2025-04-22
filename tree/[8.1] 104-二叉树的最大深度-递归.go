package tree

import "math"

// 给定一个二叉树 root ，返回其最大深度。
//
// 二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：3
//
// 示例 2：
//
// 输入：root = [1,null,2]
// 输出：2
//
// 提示：
//
// 树中节点的数量在 [0, 10⁴] 区间内。
// -100 <= Node.val <= 100
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth, rightDepth := maxDepth(root.Left), maxDepth(root.Right)
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}

/**
概念澄清：
高度：自底向上计数，某个节点 -> 根节点——后序遍历求解
深度：自顶向下计数，根节点 -> 叶子节点——前序遍历求解
最长简单路径的【边条数(从0开始)】或【节点个数(从1开始)】
*/
/**
同样也是后序遍历。
拆解下来，完整版应该是：
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	res := int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
	return res
}

可以看到，当节点为null(叶子节点的Left / Right)时，高度计为0。
是自底向上，一层层累加计数的。
先求左子树高度，再求右子树高度，最后取两者最大加上本层高度1，即为当前节点高度。

为什么题目要求最大深度，但用后序遍历求高度也能实现呢？
因为根节点的高度，就等于树的最大深度。
*/

// 真·求最大深度-前序遍历法
var maxDepthRes int

func maxDepthPreorder(root *TreeNode) int {
	maxDepthRes = 0
	if root == nil {
		return maxDepthRes
	}
	doMaxDepthPreorder(root, 1)
	return maxDepthRes
}

func doMaxDepthPreorder(node *TreeNode, depth int) {
	if node == nil {
		return
	}
	if depth > maxDepthRes {
		maxDepthRes = depth
	}
	if node.Left != nil {
		doMaxDepthPreorder(node.Left, depth+1)
	}
	if node.Right != nil {
		doMaxDepthPreorder(node.Right, depth+1)
	}
}

/**
体现回溯细节的完整写法：

if node.Left != nil {
	depth++
	doMaxDepthPreorder(node.Left, depth)
	depth--
}
*/
