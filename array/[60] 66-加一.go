package array

// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含
// 任何前导 0。
//
// 将大整数加 1，并返回结果的数字数组。
//
// 示例 1：
//
// 输入：digits = [1,2,3]
// 输出：[1,2,4]
// 解释：输入数组表示数字 123。
// 加 1 后得到 123 + 1 = 124。
// 因此，结果应该是 [1,2,4]。
//
// 示例 2：
//
// 输入：digits = [4,3,2,1]
// 输出：[4,3,2,2]
// 解释：输入数组表示数字 4321。
// 加 1 后得到 4321 + 1 = 4322。
// 因此，结果应该是 [4,3,2,2]。
//
// 示例 3：
//
// 输入：digits = [9]
// 输出：[1,0]
// 解释：输入数组表示数字 9。
// 加 1 得到了 9 + 1 = 10。
// 因此，结果应该是 [1,0]。
//
// 提示：
//
// 1 <= digits.length <= 100
// 0 <= digits[i] <= 9
// digits 不包含任何前导 0。
func plusOne(digits []int) []int {
	plus := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += plus
		plus = digits[i] / 10
		digits[i] = digits[i] % 10
	}
	if plus > 0 {
		digits = append([]int{plus}, digits...)
	}
	return digits
}

func plusOneII(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

/**
思路：
从后往前遍历，简单的不断计算、进位。
优化版本：
从后往前遍历，只给每一位+1。
- 无进位时即为最终结果，可提前返回；
- 有进位时当前位一定为0，且向前进一位。
循环以上操作直至再无进位。
*/
