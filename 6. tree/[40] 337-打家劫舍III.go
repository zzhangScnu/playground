package tree

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

var robMemo map[*TreeNode]int

func rob(root *TreeNode) int {
	robMemo = make(map[*TreeNode]int)
	return doRob(root)
}

func doRob(cur *TreeNode) int {
	if cur == nil {
		return 0
	}
	if res, ok := robMemo[cur]; ok {
		return res
	}
	res0 := doRob(cur.Left) + doRob(cur.Right)
	res1 := cur.Val
	if cur.Left != nil {
		res1 += doRob(cur.Left.Left) + doRob(cur.Left.Right)
	}
	if cur.Right != nil {
		res1 += doRob(cur.Right.Left) + doRob(cur.Right.Right)
	}
	res := max(res0, res1)
	robMemo[cur] = res
	return res
}

/**
思路：
后序遍历，子树遍历的结果作为本层的输入用于计算。
需要分情况讨论偷/不偷本节点，来控制返回值。
不偷本节点时：
结果为偷左孩子和右孩子的金额加总；
偷本节点时：
需先算上本节点的金额。需跳过左孩子，去偷左孩子的左右孩子；同理需跳过右孩子，去偷右孩子的左右孩子。
最后对这两种情况取最大值返回。

需要注意，doRob的定义是对本节点【不偷】或【考虑偷】的情况下，能取得的最大金额。
因为方法最终返回的是【偷/不偷】两种情况中的最大值，所以不一定会偷。

因为存在大量重复计算，需要引入备忘录技术。
*/
