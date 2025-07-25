package dynamicprogramming

import "math"

// 给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
//
// 下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。
// 具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。
//
// 示例 1：
//
// 输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
// 输出：13
// 解释：如图所示，为和最小的两条下降路径
//
// 示例 2：
//
// 输入：matrix = [[-19,57],[-40,-5]]
// 输出：-59
// 解释：如图所示，为和最小的下降路径
//
// 提示：
//
// n == matrix.length == matrix[i].length
// 1 <= n <= 100
// -100 <= matrix[i][j] <= 100
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		if i == 0 {
			copy(dp[i], matrix[i])
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = min(dp[i-1][j],
				func() int {
					if j > 0 {
						return dp[i-1][j-1]
					}
					return math.MaxInt
				}(),
				func() int {
					if j < n-1 {
						return dp[i-1][j+1]
					}
					return math.MaxInt
				}()) + matrix[i][j]
		}
	}
	res := math.MaxInt
	for j := 0; j < n; j++ {
		res = min(res, dp[n-1][j])
	}
	return res
}

/**
局部最小 -> 整体最小

DP数组及下标含义：
- i：第i行
- j：第j列
- DP[i][j]：行为i&列为j时，下降路径的最小和


递推公式：
matrix[i][j]的元素，只能由【左上、上、右上】下降一步而来：
for i IN (0, m)
	for j IN (0, n)
		dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i-1][j+1]) + matrix[i][j]
结果收集：
在dp数组推导完成后，在最后一行收集结果：
res = min(res, dp[i][j])
而不能在推导dp的过程中收集，因为收集的是最小值，如果下降路径此时还没生成完全，res会过早取到部分路径的和，且无法再正确更新。


初始化：
将dp第一行初始化为matrix第一行。
原始做法是起一个for循环，对每个单元格一一赋值。可以直接简洁地调用copy函数。

因为dp[i-1][j-1]和dp[i-1][j+1]可能会有列越界的风险，所以一开始的做法是将dp的列分别向左和向右扩展一列，用于存储特殊值math.MaxInt作为边界，
后续计算时for循环的范围为[1, n]。
但更好的做法是使用Go语言的函数字面量，通过判断j是否越界控制返回值。


遍历方向：
因为i依赖于i-1推导而来，故由上到下。
j无特殊要求，从左往右。
*/
