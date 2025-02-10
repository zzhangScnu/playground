package tree

// 给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树 不
// 应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在 唯一的答案 。
//
// 所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。
//
// 示例 1：
//
// 输入：root = [1,0,2], low = 1, high = 2
// 输出：[1,null,2]
//
// 示例 2：
//
// 输入：root = [3,0,4,null,2,null,null,1], low = 1, high = 3
// 输出：[3,2,null,1]
//
// 提示：
//
// 树中节点数在范围 [1, 10⁴] 内
// 0 <= Node.val <= 10⁴
// 树中每个节点的值都是 唯一 的
// 题目数据保证输入是一棵有效的二叉搜索树
// 0 <= low <= high <= 10⁴
func trimBSTIteratively(root *TreeNode, low int, high int) *TreeNode {
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	cur := root
	for cur != nil {
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right
		}
		cur = cur.Left
	}
	cur = root
	for cur != nil {
		for cur.Right != nil && cur.Right.Val > high {
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}
	return root
}

/**
将根节点先挪到区间内，再分别修改左右子树。

cur := root
for cur != nil {
	for cur.Left != nil && cur.Left.Val < low { // 注意这里是for，需要将cur.Left调整到满足区间条件
		cur.Left = cur.Left.Right
	}
	cur = cur.Left // 再进入左孩子，开始下一轮修剪循环
}
*/
