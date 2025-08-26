package dynamicprogramming

// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
//
// 示例 1：
//
// 输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
// ,["1","0","0","1","0"]]
// 输出：4
//
// 示例 2：
//
// 输入：matrix = [["0","1"],["1","0"]]
// 输出：1
//
// 示例 3：
//
// 输入：matrix = [["0"]]
// 输出：0
//
// 提示：
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 300
// matrix[i][j] 为 '0' 或 '1'
func maximalSquareRecursively(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	var traverse func(i, j int) int
	traverse = func(i, j int) int {
		if i >= m || j >= n {
			return 0
		}
		if matrix[i][j] == '0' {
			dp[i][j] = 0
			return 0
		}
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		dp[i][j] = 1 + min(traverse(i, j+1), traverse(i+1, j), traverse(i+1, j+1))
		return dp[i][j]
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			traverse(i, j)
		}
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, dp[i][j])
		}
	}
	return res * res
}

/**
思路：
不同于迭代方法的自底向上，递归是自顶向下的。
通过将问题拆解为规模更小的子问题来解决。所以必然存在大量的重复计算。
需要引入备忘录模式。
同时需要区分未计算 / 计算后结果为0的情况，故将前者通过特殊值-1进行表示。

注意，需要对每一个位置做递归计算，得到以每一个位置为左上角顶点的最大正方形边长。
最后再遍历结果集择优。
*/
