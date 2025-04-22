package backtracking

// n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
//
// 示例 1：
//
// 输入：n = 4
// 输出：2
// 解释：如上图所示，4 皇后问题存在两个不同的解法。
//
// 示例 2：
//
// 输入：n = 1
// 输出：1
//
// 提示：
//
// 1 <= n <= 9
func totalNQueens(n int) int {
	graph := make([][]bool, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]bool, n)
	}
	valid := func(graph [][]bool, x, y int) bool {
		for i := 0; i < x; i++ {
			if graph[i][y] {
				return false
			}
		}
		for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if graph[i][j] {
				return false
			}
		}
		for i, j := x-1, y+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if graph[i][j] {
				return false
			}
		}
		return true
	}
	var res int
	var traverse func(graph [][]bool, row int)
	traverse = func(graph [][]bool, row int) {
		if row == n {
			res++
			return
		}
		for j := 0; j < n; j++ {
			if valid(graph, row, j) {
				graph[row][j] = true
				traverse(graph, row+1)
				graph[row][j] = false
			}
		}
	}
	traverse(graph, 0)
	return res
}

/**
跟N皇后基本一致，需注意的细节：
1. 因放置是对每行中的每列进行尝试，所以：
	- row参数控制当前行；
	- for循环控制当前列。
   若当前graph[row][j]可放置皇后，则是一个可能解，此时占领单元格且row++，开始尝试下一行；
   同时记得回溯逻辑，遍历完[row+1, n)的所有尝试后，逐层回退，撤销占领，
   试图尝试graph[row][j+1...n)是否能放置。即二叉树中当前层级的兄弟节点及其子树。
2. 判断是否能放置：
	- 纵向：[0, row)[col]；
	- 横向：无需判断。由单层递归逻辑控制，如果该行找到了一个放置皇后的位置，则不会再往后遍历列，而是从下一行重新开始；
	- 斜向：2次循环判断，分别处理左上 / 右上。
*/
