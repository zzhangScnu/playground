package graph

// 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' 组成，捕获 所有 被围绕的区域：
//
// 连接：一个单元格与水平或垂直方向上相邻的单元格连接。
// 区域：连接所有 'O' 的单元格来形成一个区域。
// 围绕：如果您可以用 'X' 单元格 连接这个区域，并且区域中没有任何单元格位于 board 边缘，则该区域被 'X' 单元格围绕。
//
// 通过 原地 将输入矩阵中的所有 'O' 替换为 'X' 来 捕获被围绕的区域。你不需要返回任何值。
//
// 示例 1：
//
// 输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O",
// "X","X"]]
//
// 输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
//
// 解释：
//
// 在上图中，底部的区域没有被捕获，因为它在 board 的边缘并且不能被围绕。
//
// 示例 2：
//
// 输入：board = [["X"]]
//
// 输出：[["X"]]
//
// 提示：
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 200
// board[i][j] 为 'X' 或 'O'
func solve(board [][]byte) {
	movements := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(board), len(board[0])
	var traverse func(board [][]byte, x, y int)
	traverse = func(board [][]byte, x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
			return
		}
		board[x][y] = '#'
		for _, movement := range movements {
			traverse(board, x+movement[0], y+movement[1])
		}
	}
	for i := 0; i < m; i++ {
		traverse(board, i, 0)
		traverse(board, i, n-1)
	}
	for j := 0; j < n; j++ {
		traverse(board, 0, j)
		traverse(board, m-1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
