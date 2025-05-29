package tree

// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
//
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
//
// 示例 1：
//
// 输入：root = [1,2,5,3,4,null,6]
// 输出：[1,null,2,null,3,null,4,null,5,null,6]
//
// 示例 2：
//
// 输入：root = []
// 输出：[]
//
// 示例 3：
//
// 输入：root = [0]
// 输出：[0]
//
// 提示：
//
// 树中结点数在范围 [0, 2000] 内
// -100 <= Node.val <= 100
//
// 进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
func flatten(root *TreeNode) {
	var traverse func(node *TreeNode) *TreeNode
	traverse = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		left, right := traverse(node.Left), traverse(node.Right)
		node.Left, node.Right = nil, left
		cur := node
		for cur.Right != nil {
			cur = cur.Right
		}
		cur.Right = right
		return node
	}
	traverse(root)
}

/**
思路：
对于每一个节点：
1. 拉平左右子树得到left和right；
2. 将left接到根节点右孩子位置；
3. 从根节点出发，寻找右孩子的末端；
4. 将right接到末端上。
本质是一个后序遍历。

【接左子树 -> 遍历到末端 -> 接右子树】的好处是，如果左子树为空，无需做特殊边界判断，直接接右子树即可。
之前的做法是，先将左右子树接到一起，再接到根节点上。
*/
