package tree

import "math"

// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
// 有效 二叉搜索树定义如下：
//
// 节点的左子树只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
// 示例 1：
//
// 输入：root = [2,1,3]
// 输出：true
//
// 示例 2：
//
// 输入：root = [5,1,4,null,null,3,6]
// 输出：false
// 解释：根节点的值是 5 ，但是右子节点的值是 4 。
//
// 提示：
//
// 树中节点数目范围在[1, 10⁴] 内
// -2³¹ <= Node.val <= 2³¹ - 1
func isValidBSTII(root *TreeNode) bool {
	var doIsValidBST func(node *TreeNode) (int, int, bool)
	doIsValidBST = func(node *TreeNode) (int, int, bool) {
		if node == nil {
			return math.MaxInt64, math.MinInt64, true
		}
		lmin, lmax, lflag := doIsValidBST(node.Left)
		rmin, rmax, rflag := doIsValidBST(node.Right)
		if lflag && rflag && lmax < node.Val && node.Val < rmin {
			return min(node.Val, lmin), max(node.Val, rmax), true
		}
		return math.MaxInt64, math.MinInt64, false
	}
	_, _, flag := doIsValidBST(root)
	return flag
}

func isValidBSTIII(root *TreeNode) bool {
	var doIsValidBST func(node, min, max *TreeNode) bool
	doIsValidBST = func(node, min, max *TreeNode) bool {
		if node == nil {
			return true
		}
		if min != nil && node.Val <= min.Val {
			return false
		}
		if max != nil && node.Val >= max.Val {
			return false
		}
		return doIsValidBST(node.Left, min, node) && doIsValidBST(node.Right, node, max)
	}
	return doIsValidBST(root, nil, nil)
}

/**
思路：
根节点不能仅简单判断与其左右孩子的大小相对关系，要满足：
左子树所有节点 < 根节点 < 右子树所有节点。
所以有两种方式可以通过参数传递实现：
1. 后序遍历：通过函数返回值，令父节点感知左右子树中的极值，从而与自身值进行比对；
2. 先序遍历：通过函数参数值，限定左右子树各自的合法大小范围。
*/

/**
注意：
在后序遍历解法中，返回极值的大小需注意边界。
由题意可知，节点元素的大小范围为[-2³¹, 2³¹ - 1]，
而	MaxInt32  = 1<<31 - 1           // 2147483647
	MinInt32  = -1 << 31            // -2147483648
所以极小值和极大值都应取int64下的。

在先序遍历解法中，二叉搜索树的严格定义是左子树最大值 < 根节点 < 右子树最小值，
所以在判断条件中应带上 = 号。即node.Val <= min.Val和node.Val >= max.Val。
*/
