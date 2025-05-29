package tree

// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[20,9],[15,7]]
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
// -100 <= Node.val <= 100
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	flag := true
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var levelRes []int
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if flag {
				levelRes = append(levelRes, node.Val)
			} else {
				levelRes = append([]int{node.Val}, levelRes...)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, levelRes)
		flag = !flag
	}
	return res
}

/**
本质是层序遍历，只是收集结果时通过flag来控制放置的位置。

之前写的第一版，试图控制节点入列的顺序来实现。
但这样会导致后续孩子节点的遍历有误。
*/
