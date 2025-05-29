package graph

// 给你一个大小为 n x n 二进制矩阵 grid 。最多 只能将一格 0 变成 1 。
//
// 返回执行此操作后，grid 中最大的岛屿面积是多少？
//
// 岛屿 由一组上、下、左、右四个方向相连的 1 形成。
//
// 示例 1:
//
// 输入: grid = [[1, 0], [0, 1]]
// 输出: 3
// 解释: 将一格0变成1，最终连通两个小岛得到面积为 3 的岛屿。
//
// 示例 2:
//
// 输入: grid = [[1, 1], [1, 0]]
// 输出: 4
// 解释: 将一格0变成1，岛屿的面积扩大为 4。
//
// 示例 3:
//
// 输入: grid = [[1, 1], [1, 1]]
// 输出: 4
// 解释: 没有0可以让我们变成1，面积依然为 4。
//
// 提示：
//
// n == grid.length
// n == grid[i].length
// 1 <= n <= 500
// grid[i][j] 为 0 或 1

func largestIsland(grid [][]int) int {
	movements := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	n := len(grid)
	index, islandArea := 2, make(map[int]int)
	var traverse func(grid [][]int, x, y int) int
	traverse = func(grid [][]int, x, y int) int {
		if x < 0 || x >= n || y < 0 || y >= n {
			return 0
		}
		if grid[x][y] != 1 {
			return 0
		}
		grid[x][y] = index
		area := 1
		for _, movement := range movements {
			area += traverse(grid, x+movement[0], y+movement[1])
		}
		return area
	}
	allIsland := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				allIsland = false
			}
			if grid[i][j] == 1 {
				islandArea[index] = traverse(grid, i, j)
				index++
			}
		}
	}
	if allIsland {
		return n * n
	}
	var maxArea int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				continue
			}
			used := make(map[int]bool)
			area := 1
			for _, movement := range movements {
				x, y := i+movement[0], j+movement[1]
				if x < 0 || x >= n || y < 0 || y >= n || grid[x][y] == 0 || used[grid[x][y]] {
					continue
				}
				area += islandArea[grid[x][y]]
				used[grid[x][y]] = true
			}
			maxArea = max(maxArea, area)
		}
	}
	return maxArea
}

/**
思路：
1. 统计现有岛屿【编号 -> 面积】；
2. 遍历现有海水，试图将每格海水变为陆地，计算联通岛屿的最大面积。

注意点：
1. 维护岛屿编号index，为了跟海水0、陆地1区分，从2开始。在计算岛屿面积的同时，将岛屿染色为index，
   目的是方便后续遍历海水时，从相邻岛屿单元格获得index，从而快速定位到该岛屿面积；
2. 在访问海水四个方向上的联通岛屿时，需使用used辅助数组，避免形如"凹"的岛屿被重复考虑。
	每一格海水都需要有独立的used数组，所以used在双层循环内定义且初始化即可；
3. 若图中全是陆地，则可直接返回图的面积；
4. 最大岛屿面积maxArea在每格海水的上下左右四个方向遍历完成后更新。
*/
