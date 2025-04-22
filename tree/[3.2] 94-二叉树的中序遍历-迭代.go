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
func inorderTraversalIteratively(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := NewStack()
	p := root
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			stack.Push(p)
			p = p.Left
		} else {
			p = stack.Pop()
			res = append(res, p.Val)
			p = p.Right
		}
	}
	return res
}

/**
p = p.Left // 左
res = append(res, p.Val) // 中
p = p.Right // 右
*/
/**
中序遍历的【遍历顺序】和【处理顺序】是不一样的。
对于前序遍历，两者是一致的，所以处理起来比较方便；
而中序遍历，则需要用一个指针来遍历，用一个栈来处理。
*/
