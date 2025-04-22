package tree

// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
// 示例 1：
//
// 输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// 输出：3
// 解释：和等于 8 的路径有 3 条，如图所示。
//
// 示例 2：
//
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：3
//
// 提示:
//
// 二叉树的节点个数的范围是 [0,1000]
//
// -10⁹ <= Node.val <= 10⁹
// -1000 <= targetSum <= 1000

var cnt int

func pathSumIII(root *TreeNode, targetSum int) int {
	cnt = 0
	doPathSum(root, targetSum)
	return cnt
}

func doPathSum(node *TreeNode, targetSum int) int {
	if node == nil {
		return cnt
	}
	doPathSumOfNode(node, targetSum)
	doPathSum(node.Left, targetSum)
	doPathSum(node.Right, targetSum)
	return cnt
}

func doPathSumOfNode(node *TreeNode, targetSum int) {
	if node == nil {
		return
	}
	if targetSum == node.Val {
		cnt++
	}
	doPathSumOfNode(node.Left, targetSum-node.Val)
	doPathSumOfNode(node.Right, targetSum-node.Val)
}

/**
本题需要收集任意和为targetSum的路径，而不要求【根节点 -> 叶子节点】。
之前的路径总和，说白了就是只从根节点出发，
所以在本题中，除了遍历根节点，还要分别递归遍历左右孩子。

单个节点遍历逻辑：
func doPathSumOfNode(node *TreeNode, targetSum int) {
	if node == nil {
		return
	}
	// 此处无需判断是否叶子节点，也无需return，因有负数元素，往下仍可能找到结果路径
	if targetSum == node.Val {
		cnt++
	}
	doPathSumOfNode(node.Left, targetSum-node.Val)
	doPathSumOfNode(node.Right, targetSum-node.Val)
}

多个节点组合遍历逻辑：
func doPathSum(node *TreeNode, targetSum int) int {
	if node == nil {
		return cnt
	}
	doPathSumOfNode(node, targetSum) // 从本节点开始找结果
	doPathSum(node.Left, targetSum) // 在本节点的左子树中找结果
	doPathSum(node.Right, targetSum) // 在本节点的右子树中找结果
	return cnt
}

入口函数+每次调用清空全局变量的值：
func pathSumIII(root *TreeNode, targetSum int) int {
	cnt = 0
	doPathSum(root, targetSum)
	return cnt
}
*/
