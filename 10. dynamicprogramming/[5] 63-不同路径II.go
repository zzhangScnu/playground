package dynamicprogramming

// 给定一个 m x n 的整数数组 grid。一个机器人初始位于 左上角（即 grid[0][0]）。机器人尝试移动到 右下角（即 grid[m - 1][
// n - 1]）。机器人每次只能向下或者向右移动一步。
//
// 网格中的障碍物和空位置分别用 1 和 0 来表示。机器人的移动路径中不能包含 任何 有障碍物的方格。
//
// 返回机器人能够到达右下角的不同路径数量。
//
// 测试用例保证答案小于等于 2 * 10⁹。
//
// 示例 1：
//
// 输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
// 输出：2
// 解释：3x3 网格的正中间有一个障碍物。
// 从左上角到右下角一共有 2 条不同的路径：
// 1. 向右 -> 向右 -> 向下 -> 向下
// 2. 向下 -> 向下 -> 向右 -> 向右
//
// 示例 2：
//
// 输入：obstacleGrid = [[0,1],[0,0]]
// 输出：1
//
// 提示：
//
// m == obstacleGrid.length
// n == obstacleGrid[i].length
// 1 <= m, n <= 100
// obstacleGrid[i][j] 为 0 或 1
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 0 {
			dp[0][i] = 1
		}
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

/**
DP数组及下标含义：
- i, j：坐标
- dp[i][j]：走到该位置共有几种不同路径
递推公式：if obstacleGrid[i][j] != 0，dp[i][j] = dp[i-1][j] + dp[i][j-1]
- 只能向右或向下走，且每次只能走一步，故走到当前位置有几种方式 -> 从这两个方向走过来有几种方式；
- 跟爬楼梯一样，求的是几种不同的方式而不是几步，所以无需+1；
- 不会有重复的情况，因为不存在回头路；
- 特殊情况：若起始位置或终止位置有障碍物，则有0种方式可以到达右下角。
初始化：
- dp[i][j]由左方和上方的值推导而来，如dp[1][1] = dp[0][1] + dp[1][0]，故需要初始化dp[0][0...col-1]和dp[0...row-1][0]；
- 从[0, 0]出发，走到第0行的任意一列位置，都只有一种方法。但如果该位置有障碍物，则只有0种方法，也就是走不到；
- 同理，走到第0列的任意一行位置，也只有一种方法。但如果该位置有障碍物，则只有0种方法，也就是走不到。
遍历顺序：某个值要由左方或上方的值推导而来，即依赖于从左到右、从上而下的遍历顺序。
*/
