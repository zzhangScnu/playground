package tree

// 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的
// 根节点的引用。
//
// 一般来说，删除节点可分为两个步骤：
//
// 首先找到需要删除的节点；
// 如果找到了，删除它。
//
// 示例 1:
//
// 输入：root = [5,3,6,2,4,null,7], key = 3
// 输出：[5,4,6,2,null,null,7]
// 解释：给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
// 一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
// 另一个正确答案是 [5,2,6,null,4,null,7]。
//
// 示例 2:
//
// 输入: root = [5,3,6,2,4,null,7], key = 0
// 输出: [5,3,6,2,4,null,7]
// 解释: 二叉树不包含值为 0 的节点
//
// 示例 3:
//
// 输入: root = [], key = 0
// 输出: []
//
// 提示:
//
// 节点数的范围 [0, 10⁴].
// -10⁵ <= Node.val <= 10⁵
// 节点值唯一
// root 是合法的二叉搜索树
// -10⁵ <= key <= 10⁵
//
// 进阶： 要求算法时间复杂度为 O(h)，h 为树的高度。
func deleteNodeIteratively(root *TreeNode, key int) *TreeNode {
	var pre, cur *TreeNode = nil, root
	for cur != nil {
		if cur.Val == key {
			break
		}
		pre = cur
		if cur.Val > key {
			cur = cur.Left
		} else if cur.Val < key {
			cur = cur.Right
		}
	}
	if pre == nil {
		return deleteNodeReturnRoot(cur)
	}
	if key > pre.Val {
		pre.Right = deleteNodeReturnRoot(cur)
	} else {
		pre.Left = deleteNodeReturnRoot(cur)
	}
	return root
}

func deleteNodeReturnRoot(target *TreeNode) *TreeNode {
	if target == nil || target.Left == nil && target.Right == nil {
		return nil
	}
	if target.Left == nil && target.Right != nil {
		return target.Right
	}
	if target.Left != nil && target.Right == nil {
		return target.Left
	}
	p := target.Right
	for p.Left != nil {
		p = p.Left
	}
	p.Left = target.Left
	return target.Right
}

/**
func deleteNodeIteratively(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	var pre, cur *TreeNode = nil, root
	// 这里对pre的赋值有问题，最终会导致pre == cur，而不是【pre是cur的父节点】
	for cur != nil {
		pre = cur
		if cur.Val > key {
			cur = cur.Left
		} else if cur.Val < key {
			cur = cur.Right
		} else {
			break
		}
	}
	// 漏了一种情况未考虑：当要删除的节点是根节点时，此时pre == nil，如何进行连接
	if cur == nil {
		return root
	}
	// 下面的实现，很冗余，分了五种情况来删除节点，又分了两种情况来连接节点
	// 实际上，可以将五种情况收敛到一个【删除节点&返回根节点】的方法中，再分两种情况连接pre和处理后的根节点。
	if cur.Left == nil && cur.Right == nil {
		if cur.Val > pre.Val {
			pre.Right = nil
		} else {
			pre.Left = nil
		}
	} else if cur.Left == nil && cur.Right != nil {
		if cur.Val > pre.Val {
			pre.Right = cur.Right
		} else {
			pre.Left = cur.Right
		}
	} else if cur.Left != nil && cur.Right == nil {
		if cur.Val > pre.Val {
			pre.Right = cur.Left
		} else {
			pre.Left = cur.Left
		}
	} else {
		p := cur.Right
		for p.Left != nil {
			p = p.Left
		}
		p.Left = cur.Left
		if cur.Val > pre.Val {
			pre.Right = cur.Right
		} else {
			pre.Left = cur.Right
		}
	}
	return root
}

*/
