package graph

// 给你一个 m x n 的迷宫矩阵 maze （下标从 0 开始），矩阵中有空格子（用 '.' 表示）和墙（用 '+' 表示）。同时给你迷宫的入口
// entrance ，用 entrance = [entrancerow, entrancecol] 表示你一开始所在格子的行和列。
//
// 每一步操作，你可以往 上，下，左 或者 右 移动一个格子。你不能进入墙所在的格子，你也不能离开迷宫。你的目标是找到离 entrance 最近 的出口。出口
// 的含义是 maze 边界 上的 空格子。entrance 格子 不算 出口。
//
// 请你返回从 entrance 到最近出口的最短路径的 步数 ，如果不存在这样的路径，请你返回 -1 。
//
// 示例 1：
// 输入：maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], entrance =
// [1,2]
// 输出：1
// 解释：总共有 3 个出口，分别位于 (1,0)，(0,2) 和 (2,3) 。
// 一开始，你在入口格子 (1,2) 处。
// - 你可以往左移动 2 步到达 (1,0) 。
// - 你可以往上移动 1 步到达 (0,2) 。
// 从入口处没法到达 (2,3) 。
// 所以，最近的出口是 (0,2) ，距离为 1 步。
//
// 示例 2：
// 输入：maze = [["+","+","+"],[".",".","."],["+","+","+"]], entrance = [1,0]
// 输出：2
// 解释：迷宫中只有 1 个出口，在 (1,2) 处。
// (1,0) 不算出口，因为它是入口格子。
// 初始时，你在入口与格子 (1,0) 处。
// - 你可以往右移动 2 步到达 (1,2) 处。
// 所以，最近的出口为 (1,2) ，距离为 2 步。
//
// 示例 3：
// 输入：maze = [[".","+"]], entrance = [0,0]
// 输出：-1
// 解释：这个迷宫中没有出口。
//
// 提示：
//
// maze.length == m
// maze[i].length == n
// 1 <= m, n <= 100
// maze[i][j] 要么是 '.' ，要么是 '+' 。
// entrance.length == 2
// 0 <= entrancerow < m
// 0 <= entrancecol < n
// entrance 一定是空格子。
func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	movements := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := [][3]int{{entrance[0], entrance[1], 0}}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	visited[entrance[0]][entrance[1]] = true
	for len(queue) > 0 {
		x, y, step := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]
		if isExit(m, n, entrance, x, y) {
			return step
		}
		for _, movement := range movements {
			nx, ny := x+movement[0], y+movement[1]
			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}
			if maze[nx][ny] == '+' {
				continue
			}
			if visited[nx][ny] {
				continue
			}
			visited[nx][ny] = true
			queue = append(queue, [3]int{nx, ny, step + 1})
		}
	}
	return -1
}

func isExit(m, n int, entrance []int, x, y int) bool {
	return (x == 0 || y == 0 || x == m-1 || y == n-1) &&
		!(x == entrance[0] && y == entrance[1])
}
