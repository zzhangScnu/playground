package dynamicprogramming

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
func numTreesII(n int) int {

}

func doNumTreesII(start, end int) int {
	if start > end {
		return 1
	}

}

/**
函数定义：区间内能组成的二叉搜索树种数
*/
