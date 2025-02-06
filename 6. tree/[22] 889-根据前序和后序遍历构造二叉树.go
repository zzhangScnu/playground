package tree

// 给定两个整数数组，preorder 和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，postorder 是同一棵
// 树的后序遍历，重构并返回二叉树。
//
// 如果存在多个答案，您可以返回其中 任何 一个。
//
// 示例 1：
//
// 输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
// 输出：[1,2,3,4,5,6,7]
//
// 示例 2:
//
// 输入: preorder = [1], postorder = [1]
// 输出: [1]
//
// 提示：
//
// 1 <= preorder.length <= 30
// 1 <= preorder[i] <= preorder.length
// preorder 中所有值都 不同
// postorder.length == preorder.length
// 1 <= postorder[i] <= postorder.length
// postorder 中所有值都 不同
// 保证 preorder 和 postorder 是同一棵二叉树的前序遍历和后序遍历
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	postposition := make(map[int]int)
	for idx, val := range postorder {
		postposition[val] = idx
	}
	var traverse func(preorder []int, preStart, preEnd int, postorder []int, postStart, postEnd int) *TreeNode
	traverse = func(preorder []int, preStart, preEnd int, postorder []int, postStart, postEnd int) *TreeNode {
		if preStart > preEnd {
			return nil
		}
		val := preorder[preStart]
		if preStart == preEnd {
			return &TreeNode{Val: val}
		}
		pivot := postposition[preorder[preStart+1]]
		leftSize := pivot - postStart + 1
		return &TreeNode{
			Val:   val,
			Left:  traverse(preorder, preStart+1, preStart+leftSize, postorder, postStart, pivot-1),
			Right: traverse(preorder, preStart+leftSize+1, preEnd, postorder, pivot+1, postEnd-1),
		}
	}
	return traverse(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}

/**
思路：
前序的第一个元素 == 后序的最后一个元素 == 当前根节点元素
无法像前+中 / 后+中，通过根节点在中序的位置划分左右子树。
所以往前多走一步，假定前序的【根节点的后一个元素 == 左子树根节点元素】，通过左子树根节点元素在后序中划分左右子树。
要注意，递归构造左右子树时，要将postorder的右区间也缩小，排除掉此次处理的根节点元素。

根据前序+后序，无法确定唯一的一棵二叉树。
例如(1, 2, 3)，可能是左倾斜或是右倾斜的一棵树。
原因：因为在代码实现过程中，假定前序的【根节点的后一个元素 == 左子树根节点元素】。
但有可能左子树为空，实际上【根节点的后一个元素 == 右子树根节点元素】。
由于这里没法确切判断，导致了最终二叉树的不唯一。
*/

/**
	pivot := -1
	if preStart+1 <= preEnd {
		pivot = postposition[preorder[preStart+1]]
	}
	leftSize := 0
	if pivot != -1 {
		leftSize = pivot - postStart + 1
	}
一开始这样做了下对【左子树根节点元素】取值的兼容，也可以，但略显冗余。
*/
