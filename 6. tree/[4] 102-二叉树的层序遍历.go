package tree

// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
//
// 示例 2：
//
// 输入：root = [1]
// 输出：[[1]]
//
// 示例 3：
//
// 输入：root = []
// 输出：[]
//
// 提示：
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := NewQueue()
	queue.Push(root)
	for !queue.IsEmpty() {
		var levelRes []int
		levelSize := queue.Size()
		for i := 0; i < levelSize; i++ {
			p := queue.Pop()
			levelRes = append(levelRes, p.Val)
			if p.Left != nil {
				queue.Push(p.Left)
			}
			if p.Right != nil {
				queue.Push(p.Right)
			}
		}
		res = append(res, levelRes)
	}
	return res
}

/**
层序遍历：
使用队列作为辅助，边遍历节点边处理节点。
每层的节点数等于当前队列的节点数。
*/
