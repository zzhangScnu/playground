package tree

// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
// 示例 1：
//
// 输入：root = [1,null,2,3]
//
// 输出：[1,2,3]
//
// 解释：
//
// 示例 2：
//
// 输入：root = [1,2,3,4,5,null,8,null,null,6,7,9]
//
// 输出：[1,2,4,5,6,7,3,8,9]
//
// 解释：
//
// 示例 3：
//
// 输入：root = []
//
// 输出：[]
//
// 示例 4：
//
// 输入：root = [1]
//
// 输出：[1]
//
// 提示：
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
func preorderTraversalIteratively(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := NewStack()
	stack.Push(root)
	for !stack.IsEmpty() {
		p := stack.Pop()
		res = append(res, p.Val)
		if p.Right != nil {
			stack.Push(p.Right)
		}
		if p.Left != nil {
			stack.Push(p.Left)
		}
	}
	return res
}

/**
本质上是自己实现了一个栈，模拟递归时入栈出栈的过程。
队列也可以，但栈更符合计算机工作的原理。
需要注意对root的判空。
*/

/**
调用栈和手动模拟栈的区别：
1. 调用栈：
	调用顺序 -> 执行顺序，两者一致。
	调用函数时，将本层现场(变量、返回地址等)入栈暂存，并进入函数处理。
	当函数处理完后，出栈回到调用现场。
	【本质是边执行边出入栈的】
2. 手动模拟栈：
	入栈顺序与执行顺序相反。
	在遍历树时，应反着将节点入栈，才能在出栈时得到想要的顺序。
	【本质是先遍历完树入栈完毕，再出栈进行节点处理的】
*/
