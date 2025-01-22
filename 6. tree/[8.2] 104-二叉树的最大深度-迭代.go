package tree

// 给定一个二叉树 root ，返回其最大深度。
//
// 二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：3
//
// 示例 2：
//
// 输入：root = [1,null,2]
// 输出：2
//
// 提示：
//
// 树中节点的数量在 [0, 10⁴] 区间内。
// -100 <= Node.val <= 100
func maxDepthIteratively(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := NewQueue()
	queue.Push(root)
	var depth int
	for !queue.IsEmpty() {
		depth++
		size := queue.Size()
		for i := 0; i < size; i++ {
			p := queue.Pop()
			if p.Left != nil {
				queue.Push(p.Left)
			}
			if p.Right != nil {
				queue.Push(p.Right)
			}
		}
	}
	return depth
}
