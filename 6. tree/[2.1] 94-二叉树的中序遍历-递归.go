package tree

// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//
// 示例 1：
//
// 输入：root = [1,null,2,3]
// 输出：[1,3,2]
//
// 示例 2：
//
// 输入：root = []
// 输出：[]
//
// 示例 3：
//
// 输入：root = [1]
// 输出：[1]
//
// 提示：
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
func inorderTraversal(root *TreeNode) []int {
	var res []int
	doInorderTraversal(root, &res)
	return res
}

func doInorderTraversal(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	doInorderTraversal(node.Left, res)
	*res = append(*res, node.Val)
	doInorderTraversal(node.Right, res)
}
