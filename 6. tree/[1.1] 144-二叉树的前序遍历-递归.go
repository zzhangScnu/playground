package tree

// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
// 示例 1：
//
// 输入：root = [1,null,2,3]
//
// 输出：[1,2,3]
//
// 解释：
//
// 示例 2：
//
// 输入：root = [1,2,3,4,5,null,8,null,null,6,7,9]
//
// 输出：[1,2,4,5,6,7,3,8,9]
//
// 解释：
//
// 示例 3：
//
// 输入：root = []
//
// 输出：[]
//
// 示例 4：
//
// 输入：root = [1]
//
// 输出：[1]
//
// 提示：
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
func preorderTraversal(root *TreeNode) []int {
	var res []int
	doPreorderTraversal(root, &res)
	return res
}

func doPreorderTraversal(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	doPreorderTraversal(node.Left, res)
	doPreorderTraversal(node.Right, res)
}

/**
需要注意对结果集的传递，如果只传[]int类型，是按值传递的，
在函数内部对切片的修改不会影响到原始调用者中的切片。
需要作为指针传递。
*/
