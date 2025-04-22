package tree

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
// 一个节点也可以是它自己的祖先）。”
//
// 示例 1：
//
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出：3
// 解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
//
// 示例 2：
//
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// 输出：5
// 解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
//
// 示例 3：
//
// 输入：root = [1,2], p = 1, q = 2
// 输出：1
//
// 提示：
//
// 树中节点数目在范围 [2, 10⁵] 内。
// -10⁹ <= Node.val <= 10⁹
// 所有 Node.val 互不相同 。
// p != q
// p 和 q 均存在于给定的二叉树中。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil {
		return nil
	}
	if left != nil && right == nil {
		return left
	}
	if left == nil && right != nil {
		return right
	}
	return root
}

/**
求最近公共祖先，思路是自下而上遍历，
而正常来说二叉树只能从根节点开始向下遍历，所以这时候需要用到递归&回溯思想，
在回溯时处理最近公共祖先查找逻辑。

这里使用后序遍历，当找到包含p / q的子树时，就将节点逐层向上返回：
当【left == nil && right == nil】 -> 左右子树均不含p / q；
当【left != nil && right == nil】 -> 左子树含p / q，右子树不含；反之亦然；
当【left != nil && right != nil】 -> 左子树含p / q，右子树含p / q，由于是从下往上的，第一个满足该条件的节点即为最近公共祖先。
最终会将该节点返回至主方法。

情况1：p / q的最近公共祖先为r；
情况2：p的最近公共祖先为q；
实现能覆盖这两种情况。
尤其是情况2，在【if root == p || root == q 】这个判断处就能将q正确返回，不会进入子树遍历到p。
（对于情况2来说这样是没问题的，因为p需要遍历的情况是情况1）
可以画图模拟一下。
*/
