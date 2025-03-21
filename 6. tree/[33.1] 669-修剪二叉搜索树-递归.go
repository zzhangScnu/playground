package tree

// 给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树 不
// 应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在 唯一的答案 。
//
// 所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。
//
// 示例 1：
//
// 输入：root = [1,0,2], low = 1, high = 2
// 输出：[1,null,2]
//
// 示例 2：
//
// 输入：root = [3,0,4,null,2,null,null,1], low = 1, high = 3
// 输出：[3,2,null,1]
//
// 提示：
//
// 树中节点数在范围 [1, 10⁴] 内
// 0 <= Node.val <= 10⁴
// 树中每个节点的值都是 唯一 的
// 题目数据保证输入是一棵有效的二叉搜索树
// 0 <= low <= high <= 10⁴
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	} else if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

/**
单层处理逻辑：
base case：叶子节点，直接返回；
- 删除节点：
  - 节点值 < 区间左侧：说明【左子树 + 本节点】都需要移除，直接返回处理后的符合要求的右子树；
  - 节点值 > 区间左侧：说明【本节点 + 右子树】都需要移除，直接返回处理后的符合要求的左子树。
处理本节点后，再递归处理左子树和右子树，接住处理后的孩子。
*/
