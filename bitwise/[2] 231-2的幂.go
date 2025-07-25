package bitwise

// 给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。
//
// 如果存在一个整数 x 使得 n == 2ˣ ，则认为 n 是 2 的幂次方。
//
// 示例 1：
//
// 输入：n = 1
// 输出：true
// 解释：2⁰ = 1
//
// 示例 2：
//
// 输入：n = 16
// 输出：true
// 解释：2⁴ = 16
//
// 示例 3：
//
// 输入：n = 3
// 输出：false
//
// 提示：
//
// -2³¹ <= n <= 2³¹ - 1
//
// 进阶：你能够不使用循环/递归解决此问题吗？
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

/**
&：与运算符，两个数相等时运算结果 == 1，相异时运算结果 == 0。
- 若一个数 <= 0，那一定不是2的幂；
- 若一个数 > 0，如果其为2的幂，则二进制表示为有且只有一位是1。
  如1 / 10 / 100 / 1000……
  所以 n & (n - 1)若为0，则为2的幂。
*/
