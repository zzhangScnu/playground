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

/*
一维 DP 做法，针对边界条件做特化处理
*/
func minPathSumII(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n) // 注意这里是列而不是行。详情可见120题
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[j] = grid[i][j]
			} else if i == 0 {
				dp[j] = dp[j-1] + grid[i][j]
			} else if j == 0 {
				dp[j] = dp[j] + grid[i][j]
			} else {
				dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}

/*
一维 DP 做法，简洁规避边界 if 分支
*/
func minPathSumIII(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]       // 出发点，只能是当前节点的权重
	for j := 1; j < n; j++ { // 初始化第一行的 dp 值，因为第一行只能从左到右移动，不会有择优选择的过程
		dp[j] = dp[j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ { // 从第二行开始推导结果
		dp[0] = dp[0] + grid[i][0] // 初始化第一列的 dp 值，因为第一列只能从上到下移动，不会有择优选择的过程
		for j := 1; j < n; j++ {   // 从第二列开始推导结果
			dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
		}
	}
	return dp[n-1] // 结果存储在最后一行的最后一列中
}
