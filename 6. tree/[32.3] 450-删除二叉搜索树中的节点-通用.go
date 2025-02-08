package tree

func deleteNodeCommonly(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Right == nil {
			return root.Left
		}
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		root.Val, cur.Val = cur.Val, root.Val
	}
	root.Left = deleteNodeCommonly(root.Left, key)
	root.Right = deleteNodeCommonly(root.Right, key)
	return root
}

/**
整体思路：
将待删除节点的值与某个叶子节点的值做交换，再删除该叶子节点。
需要遍历整棵树，因为通用方法没有使用二叉搜索树的特性来查找值为key的节点。

当root的值 == key时：
若root.Right == nil，此时覆盖了几种场景：
- root.Left == nil：节点为叶子节点，返回root.Left，即为nil；
- root.Left != nil：返回root.Left；
若root.Right != nil，此时交换root.Right作为根节点的子树中的最左边的节点值和root的值(key)，再递归地到子树中删除值为key的节点。
*/
