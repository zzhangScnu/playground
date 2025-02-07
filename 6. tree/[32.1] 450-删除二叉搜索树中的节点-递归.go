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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		if root.Left != nil && root.Right == nil {
			return root.Left
		}
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		cur.Left = root.Left
		return root.Right
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

/**
函数定义：删除root为根节点的二叉搜索树中的值为key的节点，并返回根节点
删除节点的五种情况：
- 节点为nil：返回nil
- 节点为叶子节点：删除叶子节点，返回nil
- 节点的左孩子为空，右孩子不为空：直接返回右孩子（在单层处理逻辑中，会被父节点的Left / Right指针接住，继承删除节点的位置）
- 节点的左孩子不为空，右孩子为空：直接返回左孩子
- 节点的左右孩子均不为空：
  由于二叉搜索树的特性，删除根节点后，为保持树的有序性，
  可以【1. 将左子树挂靠在右子树的最小值节点，作为左孩子】，或 【2. 将右子树挂靠在左子树的最大值节点，作为右孩子】，
  这样中序遍历后依然会得到单调递增序列。
  对于1的实现方式，右子树的最小值节点就是右子树的最左叶子节点。挂靠后返回右子树的根节点即可。

对于递归方法来说，base case就是找到了值为key的节点，并进行删除操作（所以这个base case写起来很长）。
本递归是遍历一条边，而不是遍历整棵树。所以是带方向地遍历的。
*/

/**
对于第五种情况，有2种写法：
第一种：
cur.Left = root.Left
return root.Right
直接改变指针指向，并返回修改后的根节点；

第二种：
root.Val = cur.Val
root.Right = deleteNode(root.Right, cur.Val)
将删除节点的值替换为右子树中的最小值，再到右子树中删除该最小值节点。
*/
