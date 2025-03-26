package graph

// 二维矩阵 grid 由 0 （土地）和 1 （水）组成。岛是由最大的4个方向连通的 0 组成的群，封闭岛是一个 完全 由1包围（左、上、右、下）的岛。
//
// 请返回 封闭岛屿 的数目。
//
// 示例 1：
//
// 输入：grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,
// 0,1],[1,1,1,1,1,1,1,0]]
// 输出：2
// 解释：
// 灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。
//
// 示例 2：
//
// 输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
// 输出：1
//
// 示例 3：
//
// 输入：grid =      [
//
//					[1,1,1,1,1,1,1],
//		            [1,0,0,0,0,0,1],
//		            [1,0,1,1,1,0,1],
//		            [1,0,1,0,1,0,1],
//		            [1,0,1,1,1,0,1],
//		            [1,0,0,0,0,0,1],
//	      		    [1,1,1,1,1,1,1]
//	      						    ]
//
// 输出：2
//
// 提示：
//
// 1 <= grid.length, grid[0].length <= 100
// 0 <= grid[i][j] <=1
func closedIsland(grid [][]int) int {
	movements := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	m, n := len(grid), len(grid[0])
	var flood func(grid [][]int, x, y int)
	flood = func(grid [][]int, x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if grid[x][y] == 1 {
			return
		}
		grid[x][y] = 1
		for _, movement := range movements {
			flood(grid, x+movement[0], y+movement[1])
		}
	}
	for j := 0; j < n; j++ {
		flood(grid, 0, j)
		flood(grid, m-1, j)
	}
	for i := 0; i < m; i++ {
		flood(grid, i, 0)
		flood(grid, i, n-1)
	}
	var res int
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				res++
				flood(grid, i, j)
			}
		}
	}
	return res
}
