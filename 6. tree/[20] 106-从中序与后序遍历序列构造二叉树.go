package tree

// 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并
// 返回这颗 二叉树 。
//
// 示例 1:
//
// 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// 输出：[3,9,20,null,null,15,7]
//
// 示例 2:
//
// 输入：inorder = [-1], postorder = [-1]
// 输出：[-1]
//
// 提示:
//
// 1 <= inorder.length <= 3000
// postorder.length == inorder.length
// -3000 <= inorder[i], postorder[i] <= 3000
// inorder 和 postorder 都由 不同 的值组成
// postorder 中每一个值都在 inorder 中
// inorder 保证是树的中序遍历
// postorder 保证是树的后序遍历
func buildTree(inorder []int, postorder []int) *TreeNode {
	return doBuildTreeByInPost(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func doBuildTreeByInPost(inorder []int, inStart, inEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	if postStart > postEnd {
		return nil
	}
	val := postorder[postEnd]
	var pivot int
	for idx, num := range inorder {
		if num == val {
			pivot = idx
			break
		}
	}
	return &TreeNode{
		Val:   val,
		Left:  doBuildTreeByInPost(inorder, inStart, pivot-1, postorder, postStart, postEnd-(inEnd-pivot)-1),
		Right: doBuildTreeByInPost(inorder, pivot+1, inEnd, postorder, postEnd-(inEnd-pivot), postEnd-1),
	}
}
