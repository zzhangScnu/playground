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

/**
这就是一道模拟转圈圈的题，一开始想用一个offset变量来控制四边向内收缩多少，但会有点复杂，
还不如定义四条边界，分别控制它们向中心移动。
其次，每次移动完一条边，就要判断下一轮是否合法。不能依赖最外层的for。
最开始我用res的长度和matrix的大小进行对比，来控制是否跳出循环，但有可能出现：
能进循环，但在遍历靠后的边的时候数组越界了。
*/
