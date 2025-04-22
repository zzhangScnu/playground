package graph

// 有一个 m × n 的矩形岛屿，与 太平洋 和 大西洋 相邻。 “太平洋” 处于大陆的左边界和上边界，而 “大西洋” 处于大陆的右边界和下边界。
//
// 这个岛被分割成一个由若干方形单元格组成的网格。给定一个 m x n 的整数矩阵 heights ， heights[r][c] 表示坐标 (r, c) 上
// 单元格 高于海平面的高度 。
//
// 岛上雨水较多，如果相邻单元格的高度 小于或等于 当前单元格的高度，雨水可以直接向北、南、东、西流向相邻单元格。水可以从海洋附近的任何单元格流入海洋。
//
// 返回网格坐标 result 的 2D 列表 ，其中 result[i] = [ri, ci] 表示雨水从单元格 (ri, ci) 流动 既可流向太平洋也可
// 流向大西洋 。
//
// 示例 1：
//
// 输入: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
// 输出: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
//
// 示例 2：
//
// 输入: heights = [[2,1],[1,2]]
// 输出: [[0,0],[0,1],[1,0],[1,1]]
//
// 提示：
//
// m == heights.length
// n == heights[r].length
// 1 <= m, n <= 200
// 0 <= heights[r][c] <= 10⁵
func pacificAtlantic(heights [][]int) [][]int {
	movements := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(heights), len(heights[0])
	upleft, downright := make([][]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		upleft[i], downright[i] = make([]int, n), make([]int, n)
	}
	var traverse func(heights [][]int, visited [][]int, preHeight int, x, y int)
	traverse = func(heights [][]int, visited [][]int, preHeight int, x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if visited[x][y] == 1 {
			return
		}
		if heights[x][y] < preHeight {
			return
		}
		visited[x][y] = 1
		for _, movement := range movements {
			traverse(heights, visited, heights[x][y], x+movement[0], y+movement[1])
		}
	}
	for i := 0; i < m; i++ {
		traverse(heights, upleft, heights[i][0], i, 0)
		traverse(heights, downright, heights[i][n-1], i, n-1)
	}
	for j := 0; j < n; j++ {
		traverse(heights, upleft, heights[0][j], 0, j)
		traverse(heights, downright, heights[m-1][j], m-1, j)
	}
	var res [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if upleft[i][j] == 1 && downright[i][j] == 1 {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

/**
太平洋 -> 左边界 & 上边界
大西洋 -> 右边界 & 下边界

思路一-暴力：
遍历矩阵中的每一个元素，向四周深度搜索，看是否能找出同时到达太平洋和大西洋的路径。
是的话就将其加入结果集。
时间复杂度：O(m * n)^2
1. 遍历每一个元素，O(m * n)；
2. 基于每一个元素再遍历矩阵寻找路径，每个元素均为O(m * n)；
第二步可能会有很多重复访问和计算，容易超时。
可以尝试引入备忘录。

思路二-优化：
- 定义两个visited数组，分别代表从太平洋和大西洋出发
-> 从太平洋和大西洋边界开始向中心深度搜索
	-> 将符合海水流动方向要求的元素加入相应的visited数组中
-> 遍历矩阵中的每个元素，检查是否同时存在于太平洋的visited数组和大西洋的visited数组
	-> 意味着该元素同时符合流动到太平洋 + 流动到大西洋
	-> 加入结果集
时间复杂度：O(m * n)
*/
