package tree

import "strings"

// 按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
// 示例 1：
//
// 输入：n = 4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// 解释：如上图所示，4 皇后问题存在两个不同的解法。
//
// 示例 2：
//
// 输入：n = 1
// 输出：[["Q"]]
//
// 提示：
//
// 1 <= n <= 9
func solveNQueens(n int) [][]string {
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i] = make([]int, n)
		}
	}
	var solutions [][]string
	var doSolveNQueens func(row int)
	doSolveNQueens = func(row int) {
		if row == n {
			solutions = append(solutions, convert(board))
			return
		}
		for col := 0; col < n; col++ {
			if valid(row, col, board) {
				board[row][col] = 1
				doSolveNQueens(row + 1)
				board[row][col] = 0
			}
		}
	}
	doSolveNQueens(0)
	return solutions
}

func convert(board [][]int) []string {
	var res []string
	var sb strings.Builder
	for _, row := range board {
		sb.Reset()
		for _, value := range row {
			if value == 1 {
				sb.WriteString("Q")
			} else {
				sb.WriteString(".")
			}
		}
		res = append(res, sb.String())
	}
	return res
}

func valid(row, col int, board [][]int) bool {
	length := len(board)
	for i := 0; i < length; i++ {
		if board[i][col] == 1 {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j < length; i, j = i-1, j+1 {
		if board[i][j] == 1 {
			return false
		}
	}
	return true
}
