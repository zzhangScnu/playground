package binarysearch

// 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
//
// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
//
// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
//
// 示例 1：
//
// 输入：x = 4
// 输出：2
//
// 示例 2：
//
// 输入：x = 8
// 输出：2
// 解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
//
// 提示：
//
// 0 <= x <= 2³¹ - 1
func mySqrt(x int) int {
	low, high, mid := 1, x, 0
	for low <= high {
		mid = low + (high-low)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}

/**
思路：
使用二分搜索，用 mid * mid 无限逼近 x。
因为结束条件是 low > high，如果未命中 mid*mid == x 而提前返回，此时 high 才是最接近且未越过 x 的平方根。
*/
