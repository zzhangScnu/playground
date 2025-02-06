package tree

// 给你一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。
//
// 如果树中有不止一个众数，可以按 任意顺序 返回。
//
// 假定 BST 满足如下定义：
//
// 结点左子树中所含节点的值 小于等于 当前节点的值
// 结点右子树中所含节点的值 大于等于 当前节点的值
// 左子树和右子树都是二叉搜索树
//
// 示例 1：
//
// 输入：root = [1,null,2,2]
// 输出：[2]
//
// 示例 2：
//
// 输入：root = [0]
// 输出：[0]
//
// 提示：
//
// 树中节点的数目在范围 [1, 10⁴] 内
// -10⁵ <= Node.val <= 10⁵
//
// 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
func findMode(root *TreeNode) []int {
	var modes []int
	curCount, maxCount := 0, 0
	var pre *TreeNode
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		if pre != nil && pre.Val == node.Val {
			curCount++
		} else {
			curCount = 1
		}
		pre = node
		if curCount == maxCount {
			modes = append(modes, node.Val)
		} else if curCount > maxCount {
			modes = []int{node.Val}
			maxCount = curCount
		}
		traverse(node.Right)
	}
	traverse(root)
	return modes
}

/**
在中序遍历二叉搜索树的过程中，不断比较和统计当前节点和前一个节点的元素值。
*/
