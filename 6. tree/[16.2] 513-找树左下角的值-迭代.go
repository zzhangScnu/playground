package tree

// 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//
// 假设二叉树中至少有一个节点。
//
// 示例 1:
//
// 输入: root = [2,1,3]
// 输出: 1
//
// 示例 2:
//
// 输入: [1,2,3,4,null,5,6,null,null,7]
// 输出: 7
//
// 提示:
//
// 二叉树的节点个数的范围是 [1,10⁴]
//
// -2³¹ <= Node.val <= 2³¹ - 1
func findBottomLeftValueIteratively(root *TreeNode) int {
	queue := NewQueue()
	queue.Push(root)
	var res int
	for !queue.IsEmpty() {
		levelSize := queue.Size()
		for i := 0; i < levelSize; i++ {
			p := queue.Pop()
			if i == 0 && p.Left == nil && p.Right == nil {
				res = p.Val
			}
			if p.Left != nil {
				queue.Push(p.Left)
			}
			if p.Right != nil {
				queue.Push(p.Right)
			}
		}
	}
	return res
}

/**
不断迭代取每层最左边的叶子节点的值，最终会取到最底层，即为结果。
*/
