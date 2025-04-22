package tree

// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
// 叶子节点 是指没有子节点的节点。
//
// 示例 1：
//
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：[[5,4,11,2],[5,8,4,5]]
//
// 示例 2：
//
// 输入：root = [1,2,3], targetSum = 5
// 输出：[]
//
// 示例 3：
//
// 输入：root = [1,2], targetSum = 0
// 输出：[]
//
// 提示：
//
// 树中节点总数在范围 [0, 5000] 内
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
func pathSum(root *TreeNode, targetSum int) [][]int {
	path, res := []int{}, [][]int{}
	if root == nil {
		return res
	}
	var doPathSum func(node *TreeNode, targetSum int)
	doPathSum = func(node *TreeNode, targetSum int) {
		if node.Left == nil && node.Right == nil && targetSum == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		if node.Left != nil {
			path = append(path, node.Left.Val)
			doPathSum(node.Left, targetSum-node.Left.Val)
			path = path[:len(path)-1]
		}
		if node.Right != nil {
			path = append(path, node.Right.Val)
			doPathSum(node.Right, targetSum-node.Right.Val)
			path = path[:len(path)-1]
		}
	}
	path = append(path, root.Val)
	doPathSum(root, targetSum-root.Val)
	return res
}

/**
递归方法是否需要返回值：
如果需要遍历整棵树，则无需返回值；
如果可以提前返回，例如判断是否存在其中一条路径等，则需要返回值。
*/

/**
pathSum：
遇到叶子节点就要return。只是如果满足条件，就收集结果；
条件：【targetSum == 0】，因为在进入这层递归前，就已经执行了相减：
  if node.Left != nil {
	path = append(path, node.Left.Val)
	doPathSum(node.Left, targetSum-node.Left.Val)
	path = path[:len(path)-1]
  }
  对左右子树分别找路径，分别做递归&回溯。
  所以在收集结果后就return，表示该路径已经遍历完成。
*/

func pathSumII(root *TreeNode, targetSum int) [][]int {
	path, res := []int{}, [][]int{}
	var doPathSum func(node *TreeNode, targetSum int)
	doPathSum = func(node *TreeNode, targetSum int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && node.Val == targetSum {
			res = append(res, append([]int{}, path...))
		}
		doPathSum(node.Left, targetSum-node.Val)
		doPathSum(node.Right, targetSum-node.Val)
		path = path[:len(path)-1]
	}
	doPathSum(root, targetSum)
	return res
}

/**
pathSumII：
遇到叶子节点后，不return。只是如果满足条件，就收集结果；
条件：【targetSum == node.Val】，因为在进入这层递归前，只执行了上层的相减。
在收集结果前，需要先将本层节点加入单次路径，否则会遗漏叶子节点。

在单层逻辑中处理本层节点，再分别处理孩子节点。
所以在收集结果后不return，避免遗漏右子树等可能的结果路径。
*/
