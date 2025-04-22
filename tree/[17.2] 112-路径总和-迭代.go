package tree

import "container/list"

// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和
// targetSum 。如果存在，返回 true ；否则，返回 false 。
//
// 叶子节点 是指没有子节点的节点。
//
// 示例 1：
//
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
// 输出：true
// 解释：等于目标和的根节点到叶节点路径如上图所示。
//
// 示例 2：
//
// 输入：root = [1,2,3], targetSum = 5
// 输出：false
// 解释：树中存在两条根节点到叶子节点的路径：
// (1 --> 2): 和为 3
// (1 --> 3): 和为 4
// 不存在 sum = 5 的根节点到叶子节点的路径。
//
// 示例 3：
//
// 输入：root = [], targetSum = 0
// 输出：false
// 解释：由于树是空的，所以不存在根节点到叶子节点的路径。
//
// 提示：
//
// 树中节点的数目在范围 [0, 5000] 内
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
func hasPathSumIteratively(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		element := stack.Back()
		stack.Remove(element)
		node := element.Value.(*TreeNode)
		if node.Left == nil && node.Right == nil && node.Val == targetSum {
			return true
		}
		if node.Left != nil {
			node.Left.Val += node.Val
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			node.Right.Val += node.Val
			stack.PushBack(node.Right)
		}
	}
	return false
}

/**
go标准库中的list提供了双向链表的实现，因其提供了双端的出/入方法，可作为栈使用。
迭代的实现：手动模拟栈。
在遍历树的过程中，将父节点的值累加到左右孩子节点上，从而通过判断叶子节点值和目标值的大小，就可以得出是否有路径存在。
*/
