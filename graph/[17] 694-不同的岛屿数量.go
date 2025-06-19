package graph

import (
	"fmt"
	"strings"
)

// 给你一个由'1'(陆地)和'0'(水)组成的的二维网格,请你你计算网格中岛屿的数量。
// 岛屿总是被水包围,并且每座岛屿只能由水平方向和/或竖直直方向上相邻的陆地连接形成。此外,你可以假设
// 该网格的四条边均被水包围。
func numDistinctIslands(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	memo := make(map[string]struct{})
	sb := strings.Builder{}
	var traverse func(grid [][]int, x, y int, direction string)
	traverse = func(grid [][]int, x, y int, direction string) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if grid[x][y] == 0 {
			return
		}
		grid[x][y] = 0
		sb.WriteString(direction)
		traverse(grid, x-1, y, "1")
		traverse(grid, x+1, y, "2")
		traverse(grid, x, y-1, "3")
		traverse(grid, x, y+1, "4")
		sb.WriteString(fmt.Sprintf("-%s", direction))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				sb.Reset()
				traverse(grid, i, j, "-1")
				memo[sb.String()] = struct{}{}
			}
		}
	}
	return len(memo)
}
