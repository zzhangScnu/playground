package tree

// 给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[15,7],[9,20],[3]]
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
func levelOrderBottom(root *TreeNode) [][]int {
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
		res = append([][]int{levelRes}, res...)
	}
	return res
}

/**
层序遍历-自底向上：
跟自顶向上基本一致，只是在将每层遍历结果加入最终结果集时，
直接加在头部。
*/
