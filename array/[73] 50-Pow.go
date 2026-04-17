package array

// 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xⁿ ）。
//
// 示例 1：
//
// 输入：x = 2.00000, n = 10
// 输出：1024.00000
//
// 示例 2：
//
// 输入：x = 2.10000, n = 3
// 输出：9.26100
//
// 示例 3：
//
// 输入：x = 2.00000, n = -2
// 输出：0.25000
// 解释：2-2 = 1/22 = 1/4 = 0.25
//
// 提示：
//
// -100.0 < x < 100.0
// -231 <= n <= 231-1
// n 是一个整数
// 要么 x 不为零，要么 n > 0 。
// -104 <= xⁿ <= 104
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var traverse func(x float64, n int) float64
	traverse = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		pow := traverse(x, n/2)
		if n%2 == 0 {
			return pow * pow
		}
		return pow * pow * x
	}
	return traverse(x, n)
}

/**
思路：
一般的思路是逐个相乘，时间复杂度是O(n)
但在 n 为偶数的情况下， x ^ n = (x ^ (n / 2)) * (x ^ (n / 2))，这样时间复杂度就折半了。
如果 n 为奇数，因为 n / 2 是向下取整，所以额外多乘一个 x 即可。
*/
