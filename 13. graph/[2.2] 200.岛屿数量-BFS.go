package graph

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
// 示例 1：
//
// 输入：grid = [
//
//	["1","1","1","1","0"],
//	["1","1","0","1","0"],
//	["1","1","0","0","0"],
//	["0","0","0","0","0"]
//
// ]
// 输出：1
//
// 示例 2：
//
// 输入：grid = [
//
//	["1","1","0","0","0"],
//	["1","1","0","0","0"],
//	["0","0","1","0","0"],
//	["0","0","0","1","1"]
//
// ]
// 输出：3
//
// 提示：
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'

func numIslandsBFS(grid [][]byte) int {
	var movements = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(grid), len(grid[0])
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	var queue [][]int
	var traverse func(grid [][]byte, visited [][]int, x, y int)
	traverse = func(grid [][]byte, visited [][]int, x, y int) {
		visited[x][y] = 1
		queue = append(queue, []int{x, y})
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, movement := range movements {
				nextX, nextY := cur[0]+movement[0], cur[1]+movement[1]
				if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n {
					continue
				}
				if grid[nextX][nextY] == '0' {
					continue
				}
				if visited[nextX][nextY] == 1 {
					continue
				}
				visited[nextX][nextY] = 1
				queue = append(queue, []int{nextX, nextY})
			}
		}
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] == 0 && grid[i][j] == '1' {
				res++
				traverse(grid, visited, i, j)
			}
		}
	}
	return res
}
