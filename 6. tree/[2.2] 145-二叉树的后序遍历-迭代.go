package tree

// 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
//
// 示例 1：
//
// 输入：root = [1,null,2,3]
//
// 输出：[3,2,1]
//
// 解释：
//
// 示例 2：
//
// 输入：root = [1,2,3,4,5,null,8,null,null,6,7,9]
//
// 输出：[4,6,7,5,2,9,8,3,1]
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
// 树中节点的数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
func postorderTraversalIteratively(root *TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	}
	stack := NewStack()
	stack.Push(root)
	for !stack.IsEmpty() {
		p := stack.Pop()
		res = append(res, p.Val)
		if p.Left != nil {
			stack.Push(p.Left)
		}
		if p.Right != nil {
			stack.Push(p.Right)
		}
	}
	reverse(res)
	return res
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

/**
是前序遍历的tricky改写版本。
前序遍历是：中左右
在前序遍历的基础上，将左右入栈顺序调换，会变成中右左。
最后再翻转结果集，就变成左右中，即后序遍历结果了。
*/
