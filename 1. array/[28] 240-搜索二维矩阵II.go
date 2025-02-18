package array

// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
//
// 示例 1：
//
// 输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
// ,23,26,30]], target = 5
// 输出：true
//
// 示例 2：
//
// 输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
// ,23,26,30]], target = 20
// 输出：false
//
// 提示：
//
// m == matrix.length
// n == matrix[i].length
// 1 <= n, m <= 300
// -10⁹ <= matrix[i][j] <= 10⁹
// 每行的所有元素从左到右升序排列
// 每列的所有元素从上到下升序排列
// -10⁹ <= target <= 10⁹
func searchMatrixII(matrix [][]int, target int) bool {
	rowLength, colLength := len(matrix), len(matrix[0])
	row, col := rowLength-1, 0
	for row >= 0 && col < colLength {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			row--
		} else {
			col++
		}
	}
	return false
}

/**
最笨的方法就是每行进行二分查找，但这样无法利用上"每列的所有元素从上到下升序排列"的数据特征。
如果从左下角开始，则有明确的一个递增方向和一个递减方向，可以动态地根据当前元素大小调整横坐标或纵坐标。
同理，右上角也可以。
*/
