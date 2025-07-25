package dynamicprogramming

import . "code.byted.org/zhanglihua.river/playground/tree"

// 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
//
// 除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的
// 房子在同一天晚上被打劫 ，房屋将自动报警。
//
// 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
//
// 示例 1:
//
// 输入: root = [3,2,3,null,3,null,1]
// 输出: 7
// 解释:小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
//
// 示例 2:
//
// 输入: root = [3,4,5,1,3,null,1]
// 输出: 9
// 解释:小偷一晚能够盗取的最高金额 4 + 5 = 9
//
// 提示：
//
// 树的节点数在 [1, 10⁴] 范围内
// 0 <= Node.val <= 10⁴
func robIII(root *TreeNode) int {
	return max(doRobIII(root))
}

func doRobIII(cur *TreeNode) (int, int) {
	if cur == nil {
		return 0, 0
	}
	dpL0, dpL1 := doRobIII(cur.Left)
	dpR0, dpR1 := doRobIII(cur.Right)
	res0 := max(dpL0, dpL1) + max(dpR0, dpR1)
	res1 := dpL0 + dpR0 + cur.Val
	return res0, res1
}

/**
思路：
因为对树的遍历本质上是递归&回溯，所以无需维护一个dp数组，只需要在每层中维护当前的状态即计算结果。
本层的计算依赖于子树即下层的结果，且会返回给父节点即上层作为输入。
通过层层递归&回溯，在遍历二叉树的同时也完成了DP数组的推导。

DP数组及下标i的含义：
在二叉树的某层处理逻辑中，
- dpL[0]：不偷左孩子时的最大价值；
- dpL[1]：考虑偷左孩子时的最大价值；
- dpR[0]：不偷右孩子时的最大价值。
- dpR[1]：考虑偷右孩子时的最大价值。
注意，是【考虑】偷，而不是一定偷。
返回的res，0 -> 不偷，1 -> 偷。

递推公式：
分情况讨论。
偷本节点的情况，此时左右孩子都不能偷：
res1 := dpL[0] + dpR[0] + cur.Val
不偷本节点的情况，此时左右孩子都可以考虑偷。分别对左右孩子考虑偷/不偷来取最大值，再进行加总：
res2 := max(dpL0, dpL1) + max(dpL1, dpR1)
将这两种情况同时向上层返回。
直到在主方法中，再取最大值。

初始化：
使用默认值即可。

遍历方向：
由【本层的计算依赖于子树即下层的结果】可知，需要使用后序遍历。
*/

/**
如果不用后序遍历的话，就存在大量的重叠子问题，需要维护备忘录来消除冗余：

var memo map[*TreeNode]int

func rob(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if res, ok := memo[node]; ok {
		return res
	}
	// 如果抢当前节点，则跳过孩子节点去抢孙子节点
	rob := node.Val
	if node.Left != nil {
		rob += rob(node.Left.Left) + rob(node.Left.Right)
	}
	if node.Right != nil {
		rob += rob(node.Right.Left) + rob(node.Right.Right)
	}
	// 如果不抢当前节点，则去抢孩子节点
	notRob := rob(node.Left) + rob(node.Right)
	res := max(rob, notRob)
	memo[node] = res
	return res
}
*/
