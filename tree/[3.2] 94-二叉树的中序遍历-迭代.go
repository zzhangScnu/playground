package tree

// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//
// 示例 1：
//
// 输入：root = [1,null,2,3]
// 输出：[1,3,2]
//
// 示例 2：
//
// 输入：root = []
// 输出：[]
//
// 示例 3：
//
// 输入：root = [1]
// 输出：[1]
//
// 提示：
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
func inorderTraversalIteratively(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := NewStack()
	p := root
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			stack.Push(p)
			p = p.Left
		} else {
			p = stack.Pop()
			res = append(res, p.Val)
			p = p.Right
		}
	}
	return res
}

/**
p = p.Left // 左
res = append(res, p.Val) // 中
p = p.Right // 右
*/
/**
中序遍历的【遍历顺序】和【处理顺序】是不一样的。
对于前序遍历，两者是一致的，所以处理起来比较方便；
而中序遍历，则需要用一个指针来遍历，用一个栈来处理。
*/

/**
注意
需要初始化 cur = root，且不需要加入遍历队列中。
即初始化为：
stack := NewStack()
p := root
循环条件为：
for p != nil || !stack.IsEmpty()
*/

/**
某次的写法是：
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{root}
	var cur *TreeNode
	var count int
	for len(stack) > 0 {
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for cur.Left != nil {
			cur = cur.Left
			stack = append(stack, cur)
		}
		count++
		if count == k {
			return cur.Val
		}
		cur = cur.Right
	}
	return -1
}
会导致右子树不能正常入栈。

以测试树 3 -> (1, nil, 2), 4 为例，走一遍执行流程，就能直观看到问题：
初始栈：[3]，弹出 3 → 遍历左子树到 1，栈变为 [1]；
弹出 1 → count=1（≠3）→ cur 指向 1 的右子树 2；
下一轮循环：栈为空（因为 2 没入栈）→ 直接退出循环，返回 - 1（预期应该返回 3）。

核心问题：
右子树未入栈：处理完左子树和根节点后，仅将cur指向右子树，但未把右子树节点入栈，导致右子树（如 2、3 的右子树 4）永远无法被遍历；
左子树遍历后覆盖了根节点：内层循环直接修改cur并入栈左节点，但丢失了原本的根节点（如 3），导致根节点无法被计数。
只实现了 “左子树遍历”，但没实现中序遍历的 “根→右” 完整逻辑。
*/
