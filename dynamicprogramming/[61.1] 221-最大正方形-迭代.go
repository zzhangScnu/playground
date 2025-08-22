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
func maximalSquare(matrix [][]byte) int {
	var res int
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				dp[i][j] = 1 + min(dp[i][j+1], dp[i+1][j], dp[i+1][j+1])
				res = max(res, dp[i][j])
			}
		}
	}
	return res * res
}

/**
DP数组及下标含义
- i & j：二维矩阵坐标(i, j)
- DP[i][j]：(i, j)作为左上顶点时，最大正方形的边长


递推公式
right = dp[i][j + 1]
down = dp[i + 1][j]
diagonal = dp[i + 1][j + 1]

dp[i][j] = matrix[i][j] + min(right, down, diagonal)


遍历方向
由递推公式可知，i & j 依赖于 i + 1 & j + 1推导而来，故从下到上，由右到左。


初始化
需提前处理base case。将dp初始化为 (m + 1) * (n + 1)的大小，
dp[m][0 ... n]：表示行越界的情况，初始化为0
dp[0...m][n]：表示列越界的情况，初始化为0
结果存储在dp[0][0]中
*/
