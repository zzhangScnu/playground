package tree

// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并
// 返回其根节点。
//
// 示例 1:
//
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
//
// 示例 2:
//
// 输入: preorder = [-1], inorder = [-1]
// 输出: [-1]
//
// 提示:
//
// 1 <= preorder.length <= 3000
// inorder.length == preorder.length
// -3000 <= preorder[i], inorder[i] <= 3000
// preorder 和 inorder 均 无重复 元素
// inorder 均出现在 preorder
// preorder 保证 为二叉树的前序遍历序列
// inorder 保证 为二叉树的中序遍历序列

var inPos map[int]int

func buildTreeByPreIn(preorder []int, inorder []int) *TreeNode {
	inPos = make(map[int]int)
	for idx, num := range inorder {
		inPos[num] = idx
	}
	return doBuildTreeByPreIn(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func doBuildTreeByPreIn(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	val := preorder[preStart]
	inPivot := inPos[val]
	return &TreeNode{
		Val:   val,
		Left:  doBuildTreeByPreIn(preorder, preStart+1, preStart+(inPivot-inStart), inorder, inStart, inPivot-1),
		Right: doBuildTreeByPreIn(preorder, preStart+(inPivot-inStart)+1, preEnd, inorder, inPivot+1, inEnd),
	}
}
