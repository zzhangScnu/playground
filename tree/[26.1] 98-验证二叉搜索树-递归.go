package tree

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
func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var doIsValidBST func(node *TreeNode) bool
	doIsValidBST = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		left := doIsValidBST(node.Left)
		if !left {
			return false
		}
		if pre != nil && pre.Val >= node.Val {
			return false
		}
		pre = node
		return doIsValidBST(node.Right)
	}
	return doIsValidBST(root)
}

/**
坑：之前的做法，在单层逻辑中直接比较根节点的左右孩子的值，不满足就返回false；
但二叉搜索树的性质是，左子树最大值 < 根节点值 < 右节点最小值，
而根节点的左右孩子的值，不一定是左右子树中的极值。

二叉搜索树的【中序遍历】方式，遍历结果是单调递增的。
这道题可以利用该性质，一边中序遍历一边判断合法性。
*/

/**
var pre *TreeNode
遇到类似的需要定义为全局变量的情况，则在主方法中定义pre和递归函数，
每次调用主方法时都会初始化pre，避免重复使用、相互影响。
*/

/**
解法1：双指针。不断记录前一个节点，与当前节点值进行比较；
解法2：极值。不断记录最大值，与当前节点值进行比较。
*/

/**
pre 作为全局遍量，不会被 left/right 递归绕乱的根本原因：
执行顺序：left 递归先 “从头到尾执行完”，才会执行当前节点的 pre 校验 / 更新，最后才执行 right 递归。
left 和 right 的递归是 “串行排队”，不是 “并行抢资源”。
更新时机：pre 只在 “left 递归全完成、right 递归未开始” 时更新一次，不会被 left/right 的递归调用中途修改；
变量唯一性：pre 是外层闭包变量，全程只有一个实例，没有多份副本互相覆盖。
*/
