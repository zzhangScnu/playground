package tree

// 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
//
// 每条从根节点到叶节点的路径都代表一个数字：
//
// 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
//
// 计算从根节点到叶节点生成的 所有数字之和 。
//
// 叶节点 是指没有子节点的节点。
//
// 示例 1：
//
// 输入：root = [1,2,3]
// 输出：25
// 解释：
// 从根到叶子节点路径 1->2 代表数字 12
// 从根到叶子节点路径 1->3 代表数字 13
// 因此，数字总和 = 12 + 13 = 25
//
// 示例 2：
//
// 输入：root = [4,9,0,5,1]
// 输出：1026
// 解释：
// 从根到叶子节点路径 4->9->5 代表数字 495
// 从根到叶子节点路径 4->9->1 代表数字 491
// 从根到叶子节点路径 4->0 代表数字 40
// 因此，数字总和 = 495 + 491 + 40 = 1026
//
// 提示：
//
// 树中节点的数目在范围 [1, 1000] 内
// 0 <= Node.val <= 9
// 树的深度不超过 10
func sumNumbers(root *TreeNode) int {
	var res int
	var path []int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			res += slice2int(path)
			return
		}
		if node.Left != nil {
			traverse(node.Left)
			path = path[:len(path)-1]
		}
		if node.Right != nil {
			traverse(node.Right)
			path = path[:len(path)-1]
		}
	}
	traverse(root)
	return res
}

func slice2int(slice []int) int {
	var res int
	for _, num := range slice {
		res = res*10 + num
	}
	return res
}

func sumNumbersII(root *TreeNode) int {
	var res int
	path := []int{root.Val}
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			res += slice2int(path)
			return
		}
		if node.Left != nil {
			path = append(path, node.Left.Val)
			traverse(node.Left)
			path = path[:len(path)-1]
		}
		if node.Right != nil {
			path = append(path, node.Right.Val)
			traverse(node.Right)
			path = path[:len(path)-1]
		}
	}
	traverse(root)
	return res
}

// todo
