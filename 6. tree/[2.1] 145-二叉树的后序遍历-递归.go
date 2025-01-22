package tree

// 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
//
// 示例 1：
//
// 输入：root = [1,null,2,3]
//
// 输出：[3,2,1]
//
// 解释：
//
// 示例 2：
//
// 输入：root = [1,2,3,4,5,null,8,null,null,6,7,9]
//
// 输出：[4,6,7,5,2,9,8,3,1]
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
// 树中节点的数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
func postorderTraversal(root *TreeNode) []int {
	var res []int
	doPostorderTraversal(root, &res)
	return res
}

func doPostorderTraversal(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	doPostorderTraversal(node.Left, res)
	doPostorderTraversal(node.Right, res)
	*res = append(*res, node.Val)
}
