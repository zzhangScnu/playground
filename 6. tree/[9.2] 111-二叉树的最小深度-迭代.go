package tree

// 给定一个二叉树，找出其最小深度。
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
// 说明：叶子节点是指没有子节点的节点。
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：2
//
// 示例 2：
//
// 输入：root = [2,null,3,null,4,null,5,null,6]
// 输出：5
//
// 提示：
//
// 树中节点数的范围在 [0, 10⁵] 内
// -1000 <= Node.val <= 1000
func minDepthIteratively(root *TreeNode) int {
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
			if p.Left == nil && p.Right == nil {
				return depth
			}
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

/**
和最大深度-迭代实现的不同在于，
当前节点如果是叶子节点，就返回其深度。
保证遇到第一个叶子节点时就返回结果，即为离根节点最近的叶子节点，所以是最小深度。
*/
