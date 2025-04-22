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

/**
思路：
先将周边岛屿淹没 -> 矩阵中仅剩封闭岛屿；
再边计数边继续将访问过的孤岛淹没 -> 替代visited数组，避免重复访问导致无限循环。

注意，在grid[i][j]入列时就应【陆地 -> 海水】，这样grid[i][j]就不会重复入列。
否则，如果是先入列，下次出列时再标记为海水，即已访问，则在[入列, 出列]时间范围内，grid[i][j]若被再次访问，则会重复入栈，造成不必要的损耗。

其中，淹没岛屿时：
1. 处理周边岛屿时，应将矩阵四边都考虑周全；
2. flood函数中应【判断是陆地才处理】，否则如果在淹没周边岛屿时处理了海水，在向四周扩散深度搜索时，会将孤岛误淹没。
	if grid[x][y] == 1 {
		return
	}
	// ...
	if grid[nextX][nextY] == 0 {
		// ...
	}
	注意这里有2段这样的判断，一个是BFS函数的调用入口，一个是处理队列元素的循环中。
3. 入列的元素，应重新定义nextX和nextY，不能直接对x和y赋值和使用。很可能入列时，x和y已经被后续的for循环覆盖了。
*/
