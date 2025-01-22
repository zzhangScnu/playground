package tree

// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
//
// 示例 1：
//
// 输入：root = [4,2,7,1,3,6,9]
// 输出：[4,7,2,9,6,3,1]
//
// 示例 2：
//
// 输入：root = [2,1,3]
// 输出：[2,3,1]
//
// 示例 3：
//
// 输入：root = []
// 输出：[]
//
// 提示：
//
// 树中节点数目范围在 [0, 100] 内
// -100 <= Node.val <= 100
func invertTreeLevelly(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := NewStack()
	stack.Push(root)
	for !stack.IsEmpty() {
		levelSize := stack.Size()
		for i := 0; i < levelSize; i++ {
			p := stack.Pop()
			p.Left, p.Right = p.Right, p.Left
			if p.Left != nil {
				stack.Push(p.Left)
			}
			if p.Right != nil {
				stack.Push(p.Right)
			}
		}
	}
	return root
}
