package tree

// 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
//
// 示例 1：
//
// 输入：n = 3
// 输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
//
// 示例 2：
//
// 输入：n = 1
// 输出：[[1]]
//
// 提示：
//
// 1 <= n <= 8
func generateTrees(n int) []*TreeNode {
	return doGenerateTrees(1, n)
}

func doGenerateTrees(start, end int) []*TreeNode {
	var nodes []*TreeNode
	if start > end {
		nodes = append(nodes, nil)
		return nodes
	}
	for r := start; r <= end; r++ {
		leftNodes := doGenerateTrees(start, r-1)
		rightNodes := doGenerateTrees(r+1, end)
		for _, leftNode := range leftNodes {
			for _, rightNode := range rightNodes {
				nodes = append(nodes, &TreeNode{
					Val:   r,
					Left:  leftNode,
					Right: rightNode,
				})
			}
		}
	}
	return nodes
}

/**
在每一层中，固定根节点，递归生成左右节点的候选集，再进行左右子树的排列组合；
再将所有的根节点作为子树候选集，返回给更上一层做类似的计算。
用的是后序遍历的方式。
*/
