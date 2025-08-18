package tree

// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
// 示例 1：
//
// 输入：root = [1,2,3,null,5,null,4]
//
// 输出：[1,3,4]
//
// 解释：
//
// 示例 2：
//
// 输入：root = [1,2,3,4,null,null,null,5]
//
// 输出：[1,3,4,5]
//
// 解释：
//
// 示例 3：
//
// 输入：root = [1,null,3]
//
// 输出：[1,3]
//
// 示例 4：
//
// 输入：root = []
//
// 输出：[]
//
// 提示:
//
// 二叉树的节点个数的范围是 [0,100]
//
// -100 <= Node.val <= 100
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			if i == size-1 {
				res = append(res, cur.Val)
			}
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return res
}
