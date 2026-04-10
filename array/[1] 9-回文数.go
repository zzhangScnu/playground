package array

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
// 例如，121 是回文，而 123 不是。
//
// 示例 1：
//
// 输入：x = 121
// 输出：true
//
// 示例 2：
//
// 输入：x = -121
// 输出：false
// 解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//
// 示例 3：
//
// 输入：x = 10
// 输出：false
// 解释：从右向左读, 为 01 。因此它不是一个回文数。
//
// 提示：
//
// -2³¹ <= x <= 2³¹ - 1
//
// 进阶：你能不将整数转为字符串来解决这个问题吗？
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x%10 == 0 && x != 0 {
		return false
	}
	var revertedHalf int
	for revertedHalf < x {
		revertedHalf = revertedHalf*10 + x%10
		x /= 10
	}
	return revertedHalf == x || revertedHalf/10 == x
}

/**
思路：
将前半部分留在x，后半部分翻转后放入revertedHalf
x：每次 ÷ 10（丢弃末位） → 越来越小
revertedHalf：每次 × 10 + 末位 → 越来越大

当 revertedHalf 超过 x 时，说明后半段数字的位数 ≥ 前半段，
即已完成后半段数字的翻转。
因为原 x 的位数可能是奇数或是偶数，所以要考虑 revertedHalf 为奇数/偶数的两种情况。
*/
