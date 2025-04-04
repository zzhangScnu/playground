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
		p = queue[0]
		queue = queue[1:]
		q = queue[0]
		queue = queue[1:]
		if !isSame(p, q) {
			return false
		}
		if p != nil && q != nil {
			queue = append(queue, p.Left)
			queue = append(queue, q.Left)
			queue = append(queue, p.Right)
			queue = append(queue, q.Right)
		}
	}
	return true
}

/**
通过队列进行对比。
注意当节点是nil时也会入列，但需判断当前节点不为nil时，才会将子节点入列。

本题对层序没有额外要求，只需将两棵树对等地入列，再出列进行比较。
所以在for循环中，并没有定义levelSize := len(queue)，嵌套一个对levelSize的for循环。
*/
