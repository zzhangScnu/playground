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

func numIslands(grid [][]byte) int {
	var movements = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(grid), len(grid[0])
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	var traverse func(grid [][]byte, visited [][]int, x, y int)
	traverse = func(grid [][]byte, visited [][]int, x, y int) {
		visited[x][y] = 1
		for _, movement := range movements {
			nextX, nextY := x+movement[0], y+movement[1]
			if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n {
				continue
			}
			if grid[nextX][nextY] == '0' {
				continue
			}
			if visited[nextX][nextY] == 1 {
				continue
			}
			traverse(grid, visited, nextX, nextY)
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

/**
思路：
遍历每一个节点。
如果节点未访问 && 该节点是陆地，则表示节点是新岛屿的一部分。
以这个节点为中心，通过深度优先搜索 / 广度优先搜索，将该节点相连的所有陆地都标记为已访问，即纳入本岛屿范围。

DFS：
向上下左右四个方向递归，判断：
1. 坐标是否合法，未越界；
2. 是否陆地；
3. 是否已访问过。

base case：
在for循环中，做了预判断，不符合条件的就直接不进入下一层递归，避免无限压栈。
if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n {
	continue
}
等价于直接进入下一层递归，并增加base case：
if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n {
	return
}
*/

/**
也可以用淹没已访问岛屿的方式，加上统计岛屿时遇到海水就跳过的判断，优化掉visited辅助数组。
这种实现方式叫FloodFill。
*/
