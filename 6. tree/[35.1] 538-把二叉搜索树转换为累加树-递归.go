package tree

// 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于
// node.val 的值之和。
//
// 提醒一下，二叉搜索树满足下列约束条件：
//
// 节点的左子树仅包含键 小于 节点键的节点。
// 节点的右子树仅包含键 大于 节点键的节点。
// 左右子树也必须是二叉搜索树。
//
// 注意：本题和 1038: https://leetcode-cn.com/problems/binary-search-tree-to-greater-
// sum-tree/ 相同
//
// 示例 1：
//
// 输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
// 输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
//
// 示例 2：
//
// 输入：root = [0,null,1]
// 输出：[1,null,1]
//
// 示例 3：
//
// 输入：root = [1,0,2]
// 输出：[3,3,2]
//
// 示例 4：
//
// 输入：root = [3,2,4,1]
// 输出：[7,9,4,10]
//
// 提示：
//
// 树中的节点数介于 0 和 10⁴ 之间。
// 每个节点的值介于 -10⁴ 和 10⁴ 之间。
// 树中的所有值 互不相同 。
// 给定的树为二叉搜索树。
func convertBST(root *TreeNode) *TreeNode {
	var pre int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Right)
		node.Val += pre
		pre = node.Val
		traverse(node.Left)
	}
	traverse(root)
	return root
}

/**
刚开始的想法是，在递归后、回溯时处理，递归是自上而下，回溯是自下而上。
这个思路，针对某条路径是正确的，然而对于整棵树来说则行不通。
如中序遍历(左->中->右)，是【自上而下遍历左 -> 自下而上回溯到中并访问 ->  自上而下遍历右】，整个过程混合了两个方向的遍历。
所以正确思路是，一边遍历，一边处理。
由于二叉搜索树的中序遍历是递增有序的，为了从最右节点开始累加，应该将遍历顺序改为【右->中->左】。
*/
