package bitwise

// 给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）
// 。
//
// 示例 1：
// / 输入：left = 5, right = 7
// 输出：4
//
// 示例 2：
//
// 输入：left = 0, right = 0
// 输出：0
//
// 示例 3：
//
// 输入：left = 1, right = 2147483647
// 输出：0
//
// 提示：
//
// 0 <= left <= right <= 2³¹ - 1
func rangeBitwiseAnd(left int, right int) int {
	var i int
	for left != right {
		left, right = left>>1, right>>1
		i++
	}
	return left << i
}

/**
思路：
按位与的特点：
只要多个操作数的某一位是0，则该位的结果为0。

对[left, right]闭区间中的每一个数进行按位与运算，本质上是【寻找 left 和 right 的最长公共前缀】。
因为left ... right 是连续递增的，所以一定存在一条分割线：
- 该位置左边的部分，每一位均相等。该部分即为按位与结果；
- 该位置右边的部分，存在不相等情况。意味着存在二进制进位，即一定存在0，即按位与运算结果为0。
所以计算逻辑简化为，不断消除最低位，直到找到 left 和 right 的【最长公共前缀】。
此时再补齐低位，即为运算结果。
*/
