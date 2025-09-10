package graph

import "slices"

// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
// 示例 1：
//
// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word =
// "ABCCED"
// 输出：true
//
// 示例 2：
//
// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word =
// "SEE"
// 输出：true
//
// 示例 3：
//
// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word =
// "ABCB"
// 输出：false
//
// 提示：
//
// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board 和 word 仅由大小写英文字母组成
//
// 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	countb, countw := make(map[byte]int), make(map[byte]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			countb[board[i][j]]++
		}
	}
	for i := 0; i < len(word); i++ {
		countw[word[i]]++
		if countw[word[i]] > countb[word[i]] {
			return false
		}
	}
	if countb[word[0]] > countb[word[len(word)-1]] {
		slices.Reverse([]byte(word))
	}
	movements := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var traverse func(x, y, targetIndex int) bool
	traverse = func(x, y, targetIndex int) bool {
		if targetIndex == len(word) {
			return true
		}
		if x < 0 || x >= m || y < 0 || y >= n {
			return false
		}
		if board[x][y] == ' ' {
			return false
		}
		if board[x][y] != word[targetIndex] {
			return false
		}
		board[x][y] = ' '
		for _, movement := range movements {
			if traverse(x+movement[0], y+movement[1], targetIndex+1) {
				return true
			}
		}
		board[x][y] = word[targetIndex]
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if traverse(i, j, 0) {
				return true
			}
		}
	}
	return false
}
