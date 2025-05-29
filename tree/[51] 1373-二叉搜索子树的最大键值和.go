package tree

import "math"

// 给你一棵以 root 为根的 二叉树 ，请你返回 任意 二叉搜索子树的最大键值和。
//
// 二叉搜索树的定义如下：
//
// 任意节点的左子树中的键值都 小于 此节点的键值。
// 任意节点的右子树中的键值都 大于 此节点的键值。
// 任意节点的左子树和右子树都是二叉搜索树。
//
// 示例 1：
//
// 输入：root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
// 输出：20
// 解释：键值为 3 的子树是和最大的二叉搜索树。
//
// 示例 2：
//
// 输入：root = [4,3,null,1,2]
// 输出：2
// 解释：键值为 2 的单节点子树是和最大的二叉搜索树。
//
// 示例 3：
//
// 输入：root = [-4,-2,-5]
// 输出：0
// 解释：所有节点键值都为负数，和最大的二叉搜索树为空。
//
// 示例 4：
//
// 输入：root = [2,1,3]
// 输出：6
//
// 示例 5：
//
// 输入：root = [5,4,8,3,null,6,3]
// 输出：7
//
// 提示：
//
// 每棵树有 1 到 40000 个节点。
// 每个节点的键值在 [-4 * 10^4 , 4 * 10^4] 之间。
func maxSumBST(root *TreeNode) int {
	var maxSum int
	var traverse func(node *TreeNode) (bool, int, int, int)
	traverse = func(node *TreeNode) (bool, int, int, int) {
		if node == nil {
			return true, math.MaxInt, math.MinInt, 0
		}
		lvalid, lmin, lmax, lsum := traverse(node.Left)
		rvalid, rmin, rmax, rsum := traverse(node.Right)
		if lvalid && rvalid && lmax < node.Val && node.Val < rmin {
			sum := lsum + rsum + node.Val
			maxSum = max(maxSum, sum)
			return true, min(lmin, node.Val), max(rmax, node.Val), sum
		}
		return false, -1, -1, -1
	}
	traverse(root)
	return maxSum
}

/**
在遍历过程中，需要做几件事情：
1. 判断二叉树是否二叉搜索树；
2. 如果是二叉搜索树，需要计算最大键值和。

所以递归函数 var traverse func(node *TreeNode) (bool, int, int, int) 的返回值分别表示：
1. 以node为根节点的树，是否二叉搜索树；
2. 左子树的最大值；
3. 右子树的最小值；
4. 如果以node为根节点的树为二叉搜索树，则表示该树的键值和。

1、2、3，用于根节点联动左右子树，判断当前二叉树是否二叉搜索树；
4用于与根节点进行累加，若当前二叉树为二叉搜索树时，作为键值和返回。

初始化：
因为节点值在 [-4 * 10^4 , 4 * 10^4] 之间，
所以当遍历到最末层的叶子节点时，返回的2和3应为极值，
且在父节点中处理时，需相应取min和max值。

本题使用后序遍历的思路，使得子树信息可以被父节点使用，避免重复计算和递归嵌套递归。
*/
