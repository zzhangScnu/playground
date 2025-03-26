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
func closedIslandBFS(grid [][]int) int {
	movements := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	m, n := len(grid), len(grid[0])
	var flood func(grid [][]int, x, y int)
	flood = func(grid [][]int, x, y int) {
		if grid[x][y] == 1 {
			return
		}
		grid[x][y] = 1
		queue := [][]int{{x, y}}
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, movement := range movements {
				nextX, nextY := cur[0]+movement[0], cur[1]+movement[1]
				if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n {
					continue
				}
				if grid[nextX][nextY] == 0 {
					grid[nextX][nextY] = 1
					queue = append(queue, []int{nextX, nextY})
				}
			}
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
