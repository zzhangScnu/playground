package tree

import (
	"strconv"
	"strings"
)

// 给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//
// 叶子节点 是指没有子节点的节点。
//
// 示例 1：
//
// 输入：root = [1,2,3,null,5]
// 输出：["1->2->5","1->3"]
//
// 示例 2：
//
// 输入：root = [1]
// 输出：["1"]
//
// 提示：
//
// 树中节点的数目在范围 [1, 100] 内
// -100 <= Node.val <= 100

var path []string

var res []string

func binaryTreePaths(root *TreeNode) []string {
	path = []string{}
	res = []string{}
	doBinaryTreePaths(root)
	return res
}

func doBinaryTreePaths(node *TreeNode) {
	path = append(path, strconv.Itoa(node.Val))
	if node.Left == nil && node.Right == nil {
		res = append(res, strings.Join(path, "->"))
		return
	}
	if node.Left != nil {
		doBinaryTreePaths(node.Left)
		path = path[:len(path)-1]
	}
	if node.Right != nil {
		doBinaryTreePaths(node.Right)
		path = path[:len(path)-1]
	}
}

/**
1. 确定方法入参和返回值：
	- 入参：本次要处理的节点；
	- 返回值：空，使用全局变量来控制单条路径和结果集；
2. 确定终止条件：
	- 当节点的左右孩子均为空，即节点是叶子节点时。
	  不能以node == nil来判断，这样会将【根节点 -> 只有左/右孩子的非叶子节点】的路径也收集起来；
3. 确定单层处理逻辑：
	- 先收集元素到单条路径，再来判断是否到达base case；
	- 回溯逻辑：
		if node.Left != nil {
			doBinaryTreePaths(node.Left)
			path = path[:len(path)-1]
		}
	  在处理完左孩子后，回到本层，需要将左孩子弹出。【递归-回溯】一定是成双成对的！
*/
