package array

// 给你一个满足下述两条属性的 m x n 整数矩阵：
//
// 每行中的整数从左到右按非严格递增顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
//
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// 输出：true
//
// 示例 2：
//
// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// 输出：false
//
// 提示：
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -10⁴ <= matrix[i][j], target <= 10⁴
func searchMatrix(matrix [][]int, target int) bool {
	rowLength, colLength := len(matrix), len(matrix[0])
	low, high := 0, rowLength*colLength-1
	for low <= high {
		mid := low + (high-low)/2
		row, col := mid/colLength, mid%colLength
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

/**
核心思路是将二维数组打平至一维数组，在该一维数组上进行二分查找。
其中坐标的计算都是基于列的长度：
横坐标：mid/colLength
纵坐标：mid%colLength
*/
