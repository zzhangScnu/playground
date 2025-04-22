package tree

// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
// 一个节点也可以是它自己的祖先）。”
//
// 例如，给定如下二叉搜索树: root = [6,2,8,0,4,7,9,null,null,3,5]
//
// 示例 1:
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
// 输出: 6
// 解释: 节点 2 和节点 8 的最近公共祖先是 6。
//
// 示例 2:
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
// 输出: 2
// 解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
//
// 说明:
//
// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉搜索树中。
func lowestCommonAncestorInSearchTree(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorInSearchTree(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorInSearchTree(root.Right, p, q)
	} else {
		return root
	}
}

/**
在二叉搜索树中，由于【左子树最大值 < 根节点值 < 右子树最小值】的特性，
在自上而下的中序遍历过程中，遇到的第一个值在[p.Val, q.Val]区间内的节点，即为p和q的最近公共祖先。

- 假设第一个遍历到的节点，其值并不在[p.Val, q.Val]区间内，说明p、q均位于它的左子树或右子树中，
  那么并不存在2条分叉路径分别通向p和q。
- 假设值在[p.Val, q.Val]区间内的节点，并不是第一个被遍历到的节点，根据二叉搜索树的中序遍历特性来看，它应为p、q最近公共祖先的父节点。
*/

/**
遍历一条边：
if (递归函数(root.Left)) return
if (递归函数(root.Right)) return

遍历整棵树：
left = 递归函数(root.Left)
right = 递归函数(root.Right)
left与right的逻辑处理
*/
