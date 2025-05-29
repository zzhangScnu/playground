package graph

// 给定一个 row x col 的二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地， grid[i][j] = 0 表示水域。
//
// 网格中的格子 水平和垂直 方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
//
// 岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿
// 的周长。
//
// 示例 1：
//
// 输入：grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
// 输出：16
// 解释：它的周长是上面图片中的 16 个黄色的边
//
// 示例 2：
//
// 输入：grid = [[1]]
// 输出：4
//
// 示例 3：
//
// 输入：grid = [[1,0]]
// 输出：4
//
// 提示：
//
// row == grid.length
// col == grid[i].length
// 1 <= row, col <= 100
// grid[i][j] 为 0 或 1
func islandPerimeter(grid [][]int) int {
	var perimeter int
	movements := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	row, col := len(grid), len(grid[0])
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}
	var traverse func(grid [][]int, x, y int)
	traverse = func(grid [][]int, x, y int) {
		if x < 0 || x >= row || y < 0 || y >= col || grid[x][y] == 0 || visited[x][y] {
			return
		}
		visited[x][y] = true
		for _, movement := range movements {
			nx, ny := x+movement[0], y+movement[1]
			if nx < 0 || nx >= row || ny < 0 || ny >= col || grid[nx][ny] == 0 {
				perimeter++
				traverse(grid, nx, ny)
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			traverse(grid, i, j)
		}
	}
	traverse(grid, 0, 0)
	return perimeter
}

func islandPerimeterII(grid [][]int) int {
	var perimeter int
	movements := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 0 {
				continue
			}
			for _, movement := range movements {
				nx, ny := i+movement[0], j+movement[1]
				if nx < 0 || nx >= row || ny < 0 || ny >= col ||
					grid[nx][ny] == 0 {
					perimeter++
				}
			}
		}
	}
	return perimeter
}

/**
思路：
本题不需要用DFS或BFS来解，
直接遍历图中每一个陆地单元格，对于每一个陆地单元格，遍历它的四条边：
如果这条边跟【海水 / 图的边界】相邻，则可以贡献给岛屿的周长；
否则它处于在岛屿内部，不能被计入周长。
*/
