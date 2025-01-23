package tree

import "math"

// 给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
//
// 完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层
// 为第 h 层（从第 0 层开始），则该层包含 1~ 2ʰ 个节点。
//
// 示例 1：
//
// 输入：root = [1,2,3,4,5,6]
// 输出：6
//
// 示例 2：
//
// 输入：root = []
// 输出：0
//
// 示例 3：
//
// 输入：root = [1]
// 输出：1
//
// 提示：
//
// 树中节点的数目范围是[0, 5 * 10⁴]
// 0 <= Node.val <= 5 * 10⁴
// 题目数据保证输入的树是 完全二叉树
//
// 进阶：遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。你可以设计一个更快的算法吗？
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func countNodesByCompleteBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := root, root
	var leftDepth, rightDepth int
	for left != nil {
		leftDepth++
		left = left.Left
	}
	for right != nil {
		rightDepth++
		right = right.Right
	}
	if leftDepth == rightDepth {
		return int(math.Pow(2, float64(leftDepth))) - 1
	}
	return 1 + countNodesByCompleteBinaryTree(root.Left) + countNodesByCompleteBinaryTree(root.Right)
}

/**
countNodesByCompleteBinaryTree 结合了递归&完全二叉树特性。
满二叉树：所有位置均填满。
完全二叉树：除了最底层的叶子节点，其他位置填满。最底层节点均靠左排列。
故一颗完全二叉树可以拆解成若干棵【满二叉树】或【完全二叉子树】，
当左侧深度 == 右侧深度时，是一颗满二叉树。该树的节点个数为 2的(深度-1)次方-1

在程序中，可以看到算子树的深度时，将深度初始化为0。
所以在本树为满二叉树的情况下，本树的节点个数就可以直接用子树深度求值而不用额外减去1了。
*/
