package tree

// 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
//
// 示例 1：
//
// 输入：n = 3
// 输出：5
//
// 示例 2：
//
// 输入：n = 1
// 输出：1
//
// 提示：
//
// 1 <= n <= 19

var memo []int

func numTreesII(n int) int {
	memo = make([]int, n+1)
	return doNumTreesII(1, n)
}

func doNumTreesII(start, end int) int {
	if start > end {
		return 1
	}
	nodes := end - start + 1
	if memo[nodes] != 0 {
		return memo[nodes]
	}
	var count int
	for r := start; r <= end; r++ {
		left := doNumTreesII(start, r-1)
		right := doNumTreesII(r+1, end)
		count += left * right
	}
	memo[nodes] = count
	return count
}

/**
函数定义：区间内能组成的二叉搜索树种数。
由二叉搜索树的性质，固定根节点后，将问题切割为子问题从而递归构造左右子树。因为左右子树的范围固定，且根据【所有左子树的值 < 根节点的值 < 所有右子树的值】，可以枚举所有的子树形态，且不会重复（重复形态的树，不满足二叉搜索树定义）。
如果是普通二叉树，因为没有值的大小性质，会有不同的排列组合方式且可能出现重复情况。

这里引用了备忘录技术，将已经计算过的子问题的值记录下来。
不需要[start][end]int来记录，因为递增序列元素数量相同的情况下，可组成的子树种类是一样的。
数值分布和形态可能不一样，但是种数是一样的。
在这里可以用备忘录技术减少计算次数，但是在生成不同的二叉搜索树的时候就不能算是相同的场景了，不能去重。
*/
