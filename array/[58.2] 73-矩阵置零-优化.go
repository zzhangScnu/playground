package array

// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
//
// 示例 1：
//
// 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
// 输出：[[1,0,1],[0,0,0],[1,0,1]]
//
// 示例 2：
//
// 输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// 输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
//
// 提示：
//
// m == matrix.length
// n == matrix[0].length
// 1 <= m, n <= 200
// -2³¹ <= matrix[i][j] <= 2³¹ - 1
//
// 进阶：
//
// 一个直观的解决方案是使用 O(mn) 的额外空间，但这并不是一个好的解决方案。
// 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
// 你能想出一个仅使用常量空间的解决方案吗？
func setZeroesII(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row0, col0 := false, false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					row0 = true
				}
				if j == 0 {
					col0 = true
				}
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for col := 0; col < n; col++ {
			matrix[0][col] = 0
		}
	}
	if col0 {
		for row := 0; row < m; row++ {
			matrix[row][0] = 0
		}
	}
}

/**
优化：
可以用原矩阵的第0行和第0列记录矩阵中的原始0坐标，作用等同于原做法中的rows和cols。
但是需要额外的2个bool变量记录第0行和第0列本身是否有原始0。避免第0行和第0列中的原始0被记录值污染。

记录完成后，处理矩阵时，需先根据第0行和第0列的记录，处理第[1, n - 1]行和第[1, n - 1]列，
再根据bool变量处理第0行和第0列。

仅使用常量空间。
*/
