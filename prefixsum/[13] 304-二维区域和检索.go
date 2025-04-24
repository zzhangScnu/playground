package prefixsum

// NumMatrix 给定一个二维矩阵 matrix，以下类型的多个请求：
//
// 计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
//
// 实现 NumMatrix 类：
//
// NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
// int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下
// 角 (row2, col2) 所描述的子矩阵的元素 总和 。
//
// 示例 1：
//
// 输入:
// ["NumMatrix","sumRegion","sumRegion","sumRegion"]
// [[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,
// 1,2,2],[1,2,2,4]]
// 输出:
// [null, 8, 11, 12]
//
// 解释:
// NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,
// 0,1,7],[1,0,3,0,5]]);
// numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
// numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
// numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)
//
// 提示：
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 200
//
// -10⁵ <= matrix[i][j] <= 10⁵
// 0 <= row1 <= row2 < m
// 0 <= col1 <= col2 < n
//
// 最多调用 10⁴ 次 sumRegion 方法
type NumMatrix struct {
	preSum [][]int
}

func Constructor1(matrix [][]int) NumMatrix {
	rows, cols := len(matrix)+1, len(matrix[0])+1
	preSum := make([][]int, rows)
	preSum[0] = make([]int, cols)
	for i := 1; i < rows; i++ {
		preSum[i] = make([]int, cols)
		for j := 1; j < cols; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{preSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1] + this.preSum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

// 之前的写法如下：
/**
func Constructor(matrix [][]int) NumMatrix {
	rows, cols := len(matrix), len(matrix[0])
	preSum := make([][]int, rows)
	for i := 0; i < rows; i++ {
		preSum[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			if i == 0 && j == 0 {
				preSum[i][j] = matrix[i][j]
			} else if i == 0 {
				preSum[i][j] = preSum[i][j-1] + matrix[i][j]
			} else if j == 0 {
				preSum[i][j] = preSum[i-1][j] + matrix[i][j]
			} else {
				preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + preSum[i-1][j-1] + matrix[i][j]
			}
		}
	}
	return NumMatrix{preSum}
}
*/
/**
这种写法，无论在构造前缀和矩阵，还是在计算区域和的时候，都需要大量边界处理防止数组越界。
所以可以将前缀和矩阵虚拟出第一行和第一列，期间的值全为0。
这样，preSum[i][j]存储的实际是preSum[i-1][j-1]的结果。
*/
