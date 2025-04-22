package array

// 给你一个正整数 n ，生成一个包含 1 到 n² 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
// 示例 1：
//
// 输入：n = 3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]
//
// 示例 2：
//
// 输入：n = 1
// 输出：[[1]]
//
// 提示：
//
// 1 <= n <= 20
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	rowTop, rowBottom, colLeft, colRight := 0, n-1, 0, n-1
	val := 1
	for {
		for j := colLeft; j <= colRight; j++ {
			res[rowTop][j] = val
			val++
		}
		rowTop++
		if rowTop > rowBottom {
			break
		}
		for i := rowTop; i <= rowBottom; i++ {
			res[i][colRight] = val
			val++
		}
		colRight--
		if colRight < colLeft {
			break
		}
		for j := colRight; j >= colLeft; j-- {
			res[rowBottom][j] = val
			val++
		}
		rowBottom--
		if rowBottom < rowTop {
			break
		}
		for i := rowBottom; i >= rowTop; i-- {
			res[i][colLeft] = val
			val++
		}
		colLeft++
		if colLeft > colRight {
			break
		}
	}
	return res
}
