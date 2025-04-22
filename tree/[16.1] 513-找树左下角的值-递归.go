package tree

// 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//
// 假设二叉树中至少有一个节点。
//
// 示例 1:
//
// 输入: root = [2,1,3]
// 输出: 1
//
// 示例 2:
//
// 输入: [1,2,3,4,null,5,6,null,null,7]
// 输出: 7
//
// 提示:
//
// 二叉树的节点个数的范围是 [1,10⁴]
//
// -2³¹ <= Node.val <= 2³¹ - 1

var maximumDepth int

var bottomLeftValue int

func findBottomLeftValue(root *TreeNode) int {
	maximumDepth = 0
	bottomLeftValue = 0
	doFindBottomLeftValue(root, 1)
	return bottomLeftValue
}

func doFindBottomLeftValue(node *TreeNode, depth int) {
	if depth > maximumDepth && node.Left == nil && node.Right == nil {
		maximumDepth = depth
		bottomLeftValue = node.Val
		return
	}
	if node.Left != nil {
		doFindBottomLeftValue(node.Left, depth+1)
	}
	if node.Right != nil {
		doFindBottomLeftValue(node.Right, depth+1)
	}
}

/**
在遍历到深度更深的叶子节点时，更新结果值。
由于使用了左优先的遍历方式，会在更新深度时，仅更新一次结果值。从而保证取的是左叶子节点的值。
*/
