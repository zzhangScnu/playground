package graph

// 给你两个 m x n 的二进制矩阵 grid1 和 grid2 ，它们只包含 0 （表示水域）和 1 （表示陆地）。一个 岛屿 是由 四个方向 （水平或者竖
// 直）上相邻的 1 组成的区域。任何矩阵以外的区域都视为水域。
//
// 如果 grid2 的一个岛屿，被 grid1 的一个岛屿 完全 包含，也就是说 grid2 中该岛屿的每一个格子都被 grid1 中同一个岛屿完全包含，那
// 么我们称 grid2 中的这个岛屿为 子岛屿 。
//
// 请你返回 grid2 中 子岛屿 的 数目 。
//
// 示例 1：
// 输入：grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]],
// grid2 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
// 输出：3
// 解释：如上图所示，左边为 grid1 ，右边为 grid2 。
// grid2 中标红的 1 区域是子岛屿，总共有 3 个子岛屿。
//
// 示例 2：
// 输入：grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]],
// grid2 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
// 输出：2
// 解释：如上图所示，左边为 grid1 ，右边为 grid2 。
// grid2 中标红的 1 区域是子岛屿，总共有 2 个子岛屿。
//
// 提示：
//
// m == grid1.length == grid2.length
// n == grid1[i].length == grid2[i].length
// 1 <= m, n <= 500
// grid1[i][j] 和 grid2[i][j] 都要么是 0 要么是 1 。
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid2), len(grid2[0])
	movements := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var traverse func(grid [][]int, x, y int)
	traverse = func(grid [][]int, x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if grid[x][y] == 0 {
			return
		}
		grid[x][y] = 0
		for _, movement := range movements {
			traverse(grid, x+movement[0], y+movement[1])
		}
	}
	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[0]); j++ {
			if grid1[i][j] == 0 && grid2[i][j] == 1 {
				traverse(grid2, i, j)
			}
		}
	}
	var count int
	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[0]); j++ {
			if grid2[i][j] == 1 {
				count++
				traverse(grid2, i, j)
			}
		}
	}
	return count
}
