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
func getMinimumDifferenceIteratively(root *TreeNode) int {
	minimum := math.MaxInt
	var pre *TreeNode
	stack := NewStack()
	for root != nil || !stack.IsEmpty() {
		if root != nil {
			stack.Push(root)
			root = root.Left
		} else {
			root = stack.Pop()
			if pre != nil && root.Val-pre.Val < minimum {
				minimum = root.Val - pre.Val
			}
			pre = root
			root = root.Right
		}
	}
	return minimum
}

/**
本质上是二叉树的中序遍历-迭代实现，
踩坑记录：
1. 在遍历中将根节点和左子树入栈即可；
2. 当最左路径遍历到底后，需要出栈访问后，再向右路径访问，此时root = root.Right。
*/
