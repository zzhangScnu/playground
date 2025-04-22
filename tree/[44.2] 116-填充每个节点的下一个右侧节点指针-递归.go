package tree

// 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
//
//	struct Node {
//	 int val;
//	 Node *left;
//	 Node *right;
//	 Node *next;
//	}
//
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
//
// 初始状态下，所有 next 指针都被设置为 NULL。
//
// 示例 1：
//
// 输入：root = [1,2,3,4,5,6,7]
// 输出：[1,#,2,3,#,4,5,6,7,#]
// 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由
// next 指针连接，'#' 标志着每一层的结束。
//
// 示例 2:
//
// 输入：root = []
// 输出：[]
//
// 提示：
//
// 树中节点的数量在
// [0, 2¹² - 1] 范围内
// -1000 <= node.val <= 1000
//
// 进阶：
//
// 你只能使用常量级额外空间。
// 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
func connectRecursively(root *Node) *Node {
	var traverse func(p, q *Node)
	traverse = func(p, q *Node) {
		if p == nil && q == nil {
			return
		}
		p.Next = q
		traverse(p.Left, p.Right)
		traverse(p.Right, q.Left)
		traverse(q.Left, q.Right)
	}
	if root == nil {
		return nil
	}
	traverse(root.Left, root.Right)
	return root
}

/**
传入两个兄弟节点p和q，单层逻辑中处理p指向q。
并递归处理p、q左右孩子的指向。

注意，因为题目给定的是完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点，
所以不会存在p和q只有其一为nil的情况，只有同时为nil，此时表示遍历到叶子节点的下一层级了。
*/
