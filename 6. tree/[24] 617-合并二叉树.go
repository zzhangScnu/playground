package tree

// 给你两棵二叉树： root1 和 root2 。
//
// 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠
// ，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。
//
// 返回合并后的二叉树。
//
// 注意: 合并过程必须从两个树的根节点开始。
//
// 示例 1：
//
// 输入：root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// 输出：[3,4,5,5,4,null,7]
//
// 示例 2：
//
// 输入：root1 = [1], root2 = [1,2]
// 输出：[2,2]
//
// 提示：
//
// 两棵树中的节点数目在范围 [0, 2000] 内
// -10⁴ <= Node.val <= 10⁴
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	return doMergeTrees(root1, root2)
}

func doMergeTrees(node1, node2 *TreeNode) *TreeNode {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	return &TreeNode{
		Val:   node1.Val + node2.Val,
		Left:  doMergeTrees(node1.Left, node2.Left),
		Right: doMergeTrees(node1.Right, node2.Right),
	}
}
