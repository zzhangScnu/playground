package tree

// 给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
// 示例 1：
//
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
//
// 示例 2：
//
// 输入：root = [1,2,2,null,3,null,3]
// 输出：false
//
// 提示：
//
// 树中节点数目在范围 [1, 1000] 内
// -100 <= Node.val <= 100
//
// 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return doIsSymmetric(root.Left, root.Right)
}

func doIsSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val &&
		doIsSymmetric(left.Left, right.Right) && doIsSymmetric(left.Right, right.Left)
}

/**
这种写法其实不对，违背了处理本层+递归调用处理左右子树的原则，也没法实现树的外侧vs外侧+内侧vs内侧的比较。
if root == nil {
	return true
}
if root.Left == nil && root.Right == nil {
	return true
}
if root.Left == nil || root.Right == nil {
	return false
}
return root.Left.Val == root.Right.Val &&
	isSymmetric(root.Left) && isSymmetric(root.Right)
*/

/**
这里是后序遍历：
1. 收集孩子信息——判断左右子树是否镜像；
2. 再在当前节点做处理。
代码实现较为精简，更符合后序逻辑的实现如下：
func doIsSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	isOutsideSymmetric := doIsSymmetric(left.Left, right.Right)
	isInsideSymmetric := doIsSymmetric(left.Right, right.Left)
	res := isOutsideSymmetric && isInsideSymmetric // 后
	return res
}
*/
