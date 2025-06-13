package unionfindset

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
	m, n := len(board), len(board[0])
	dummyNode := m * n
	unionFindSet := NewUnionFindSet(m*n + 1)
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			unionFindSet.Union(j, dummyNode)
		}
		if board[m-1][j] == 'O' {
			unionFindSet.Union((m-1)*n+j, dummyNode)
		}
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			unionFindSet.Union(i*n, dummyNode)
		}
		if board[i][n-1] == 'O' {
			unionFindSet.Union(i*n+n-1, dummyNode)
		}
	}
	movements := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] != 'O' {
				continue
			}
			for _, movement := range movements {
				newI, newJ := i+movement[0], j+movement[1]
				if board[newI][newJ] != 'O' {
					continue
				}
				unionFindSet.Union(newI*n+newJ, i*n+j)
			}
		}
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if !unionFindSet.IsConnected(i*n+j, dummyNode) {
				board[i][j] = 'X'
			}
		}
	}
}
