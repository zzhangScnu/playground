package graph

// 在给定的 m x n 网格
// grid 中，每个单元格可以有以下三个值之一：
//
// 值 0 代表空单元格；
// 值 1 代表新鲜橘子；
// 值 2 代表腐烂的橘子。
//
// 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
//
// 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
//
// 示例 1：
//
// 输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
// 输出：4
//
// 示例 2：
//
// 输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
// 输出：-1
// 解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个方向上。
//
// 示例 3：
//
// 输入：grid = [[0,2]]
// 输出：0
// 解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
//
// 提示：
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 10
// grid[i][j] 仅为 0、1 或 2
func orangesRotting(grid [][]int) int {
	movements := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(grid), len(grid[0])
	var queue [][]int
	var fresh int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	if fresh == 0 {
		return 0
	}
	var time int
	for len(queue) > 0 {
		rottedFlag, size := false, len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			i, j := cur[0], cur[1]
			for _, movement := range movements {
				nextI, nextJ := i+movement[0], j+movement[1]
				if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n &&
					grid[nextI][nextJ] == 1 {
					fresh--
					rottedFlag = true
					grid[nextI][nextJ] = 2
					queue = append(queue, []int{nextI, nextJ})
				}
			}
		}
		if rottedFlag {
			time++
		}
	}
	if fresh > 0 {
		return -1
	}
	return time
}

/**
思路：
每一轮(每一分钟)，所有腐烂的橘子均向外扩散一圈，腐蚀周边的新鲜橘子。
使用BFS实现齐头并进的效果。

注意：
1. 在遍历grid时顺便统计新鲜橘子的数量，在腐蚀时递减。最终可用于判断是否已全部腐烂，节约再次遍历grid进行判断的O(n^2)时间复杂度；
2. 在每轮扩散时维护一个“是否有新鲜橘子被腐蚀”的标志位，防止全部橘子已腐烂的情况下无效累加时长；
3. 直接原地修改grid来实现visited辅助数组控制不重复访问的效果；
4. 注意控制当前所有腐烂的橘子向外腐蚀一圈时，这两种写法并不等价：
	- for i := 0; i < len(queue); i++ // 循环次数会随着对queue的操作动态变化
	- size := len(queue) // 固定循环次数
      for i := 0; i < size; i++ // 循环时与queue的操作脱钩
*/
