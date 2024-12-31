package array

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//
// 示例 1：
//
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
//
// 示例 2：
//
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
// 提示：
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	var res []int
	rowTop, rowBottom, colLeft, colRight := 0, m-1, 0, n-1
	for {
		for j := colLeft; j <= colRight; j++ {
			res = append(res, matrix[rowTop][j])
		}
		rowTop++
		if rowTop > rowBottom {
			break
		}
		for i := rowTop; i <= rowBottom; i++ {
			res = append(res, matrix[i][colRight])
		}
		colRight--
		if colLeft > colRight {
			break
		}
		for j := colRight; j >= colLeft; j-- {
			res = append(res, matrix[rowBottom][j])
		}
		rowBottom--
		if rowTop > rowBottom {
			break
		}
		for i := rowBottom; i >= rowTop; i-- {
			res = append(res, matrix[i][colLeft])
		}
		colLeft++
		if colLeft > colRight {
			break
		}
	}
	return res
}
