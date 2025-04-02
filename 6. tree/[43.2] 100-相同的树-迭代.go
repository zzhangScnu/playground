package tree

// 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
//
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
//
// 示例 1：
//
// 输入：p = [1,2,3], q = [1,2,3]
// 输出：true
//
// 示例 2：
//
// 输入：p = [1,2], q = [1,null,2]
// 输出：false
//
// 示例 3：
//
// 输入：p = [1,2,1], q = [1,1,2]
// 输出：false
//
// 提示：
//
// 两棵树上的节点数目都在范围 [0, 100] 内
// -10⁴ <= Node.val <= 10⁴
func isSameTreeIteratively(p *TreeNode, q *TreeNode) bool {
	isSame := func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		} else if p == nil || q == nil {
			return false
		} else if p.Val != q.Val {
			return false
		}
		return true
	}
	queue := []*TreeNode{p, q}
	for len(queue) > 0 {
		for i := 0; i < len(queue); i++ {
			p = queue[0]
			queue = queue[1:]
			q = queue[0]
			queue = queue[1:]
			if !isSame(p, q) {
				return false
			}
			queue = append(queue, p.Left)
			queue = append(queue, q.Left)
			queue = append(queue, p.Right)
			queue = append(queue, q.Right)
		}
	}
	return true
}
