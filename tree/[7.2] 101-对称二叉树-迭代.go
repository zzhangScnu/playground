package tree

// 给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
// 示例 1：
//
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
//
// 示例 2：
//
// 输入：root = [1,2,2,null,3,null,3]
// 输出：false
//
// 提示：
//
// 树中节点数目在范围 [1, 1000] 内
// -100 <= Node.val <= 100
//
// 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
func isSymmetricIteratively(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := NewStack()
	stack.Push(root.Right)
	stack.Push(root.Left)
	for !stack.IsEmpty() {
		levelSize := stack.Size()
		for i := 0; i < levelSize; i++ {
			left := stack.Pop()
			right := stack.Pop()
			if left == nil && right == nil {
				continue
			}
			if left == nil || right == nil {
				return false
			}
			if left.Val != right.Val {
				return false
			}
			stack.Push(right.Right)
			stack.Push(left.Left)
			stack.Push(left.Right)
			stack.Push(right.Left)
		}
	}
	return true
}
