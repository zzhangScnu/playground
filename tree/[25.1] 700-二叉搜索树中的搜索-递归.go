package tree

// 给定二叉搜索树（BST）的根节点
// root 和一个整数值
// val。
//
// 你需要在 BST 中找到节点值等于 val 的节点。 返回以该节点为根的子树。 如果节点不存在，则返回
// null 。
//
// 示例 1:
//
// 输入：root = [4,2,7,1,3], val = 2
// 输出：[2,1,3]
//
// 示例 2:
//
// 输入：root = [4,2,7,1,3], val = 5
// 输出：[]
//
// 提示：
//
// 树中节点数在 [1, 5000] 范围内
// 1 <= Node.val <= 10⁷
// root 是二叉搜索树
// 1 <= val <= 10⁷
func searchBST(root *TreeNode, val int) *TreeNode {
	if root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

/**
二叉搜索树性质：
- 左子树最大值 < 根节点值 < 右节点最小值（而不是左、右孩子的值）；
-【中序遍历】方式，遍历结果是单调递增的；
- 相比起二叉树的遍历框架，增加了对BST左小右大特性的使用，引入了类似二分搜索的思想。
*/
