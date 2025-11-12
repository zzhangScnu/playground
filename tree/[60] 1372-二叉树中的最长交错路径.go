package tree

// 给你一棵以 root 为根的二叉树，二叉树中的交错路径定义如下：
//
// 选择二叉树中 任意 节点和一个方向（左或者右）。
// 如果前进方向为右，那么移动到当前节点的的右子节点，否则移动到它的左子节点。
// 改变前进方向：左变右或者右变左。
// 重复第二步和第三步，直到你在树中无法继续移动。
//
// 交错路径的长度定义为：访问过的节点数目 - 1（单个节点的路径长度为 0 ）。
//
// 请你返回给定树中最长 交错路径 的长度。
//
// 示例 1：
//
// 输入：root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]
// 输出：3
// 解释：蓝色节点为树中最长交错路径（右 -> 左 -> 右）。
//
// 示例 2：
//
// 输入：root = [1,1,1,null,1,null,null,1,1,null,1]
// 输出：4
// 解释：蓝色节点为树中最长交错路径（左 -> 右 -> 左 -> 右）。
//
// 示例 3：
//
// 输入：root = [1]
// 输出：0
//
// 提示：
//
// 每棵树最多有 50000 个节点。
// 每个节点的值在 [1, 100] 之间。
func longestZigZag(root *TreeNode) int {
	var res int
	var traverse func(node *TreeNode, direction int) int // -1: left; 1: right
	traverse = func(node *TreeNode, direction int) int {
		if node == nil {
			return -1
		}
		if direction == -1 {
			left := traverse(node.Left, 1) + 1
			right := traverse(node.Right, -1) + 1
			res = max(res, left, right)
			return left
		} else {
			right := traverse(node.Right, -1) + 1
			left := traverse(node.Left, 1) + 1
			res = max(res, left, right)
			return right
		}
	}
	traverse(root, -1)
	traverse(root, 1)
	return res
}

/**
注意点：
1. 最长交错路径未必从根节点开始，也未必于叶子节点结束。可能是二叉树中间的一段；
2. 本题适用于后序遍历，在计算出左 / 右子树的最长交错路径长度后，择优累加本节点的长度；
3. 用入参direction表示【下一步】方向，-1 为向左，1 为向右；
4. 在每一层中都有重新计算的可能性：以所有节点为潜在起点，通过递归探索每个节点在【向左延伸】和【向右延伸】两种方向下的最长路径，实时更新全局最大值。
   本质是“覆盖所有节点和所有可能起始方向的全面探索”。
5. 每个节点在处理时，会同时做两件事。以 direction == -1 为例：
	- 【延伸路径】按当前要求的方向延伸。如要求向左，则递归左孩子，且下一次方向必须向右。
 	  即 left := traverse(node.Left, 1)
	  left 会被返回给父节点，作为子问题结果进行路径长度累加的值，用于拼接更长的交错路径。
	- 【重启路径】从当前节点反向重新开始。如要求向左，但尝试从当前节点向右开始新路径，下一次方向必须向左。
	  即 right := traverse(node.Right, -1)
	- 比较 left 和 right 的大小，实时更新全局最大值。而不是等到叶子节点再更新。
	  因为无法延续上一步的路径，所以只能作为独立路径存在，即 right 仅用于更新全局最大值。

误区：
- 从根节点出发，固定初始方向（向左或向右），尝试沿“交替方向”延伸路径，仅在叶子节点更新最长路径长度。本质是“从根节点开始的两条固定路径探索”。
*/
