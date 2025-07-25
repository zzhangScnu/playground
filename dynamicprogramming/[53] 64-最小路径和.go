package dynamicprogramming

// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
// 示例 1：
//
// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。
//
// 示例 2：
//
// 输入：grid = [[1,2,3],[4,5,6]]
// 输出：12
//
// 提示：
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 200
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

/**
DP数组及下标含义：
- dp[i][j]：从(0, 0)到(i, j)的最小路径和。

递推公式：
dp[i][j] = min(
    dp[i-1][j],
    dp[i][j-1]
) + grid[i][j]

遍历方向：
i依赖于i-1推导而来，j依赖于j-1推导而来，故从左到右，从上到下遍历。

初始化：
初始化 i == 0 和 j == 0 的情况，避免后续计算时数组越界。
i == 0时，意味着路径为第一行，每个单元格的路径和为前序加总；
j == 0时，意味着路径为第一列，每个单元格的路径和为前序加总。
*/
