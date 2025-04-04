package dynamicprogramming

// 给定一个包含大写字母和小写字母的字符串
// s ，返回 通过这些字母构造成的 最长的 回文串 的长度。
//
// 在构造过程中，请注意 区分大小写 。比如 "Aa" 不能当做一个回文字符串。
//
// 示例 1:
//
// 输入:s = "abccccdd"
// 输出:7
// 解释:
// 我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
//
// 示例 2:
//
// 输入:s = "a"
// 输出:1
// 解释：可以构造的最长回文串是"a"，它的长度是 1。
//
// 提示:
//
// 1 <= s.length <= 2000
// s 只由小写 和/或 大写英文字母组成
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	var res, start, end int
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[j] != s[i] {
				continue
			}
			if i == j || i+1 == j || dp[i+1][j-1] {
				dp[i][j] = true
				if j-i+1 > res {
					res, start, end = j-i+1, i, j
				}
			}
		}
	}
	return s[start : end+1]
}

/**
DP数组及其下标含义：
- i：从i开始，
- j：到j结束，
- dp[i][j]：子字符串[i...j]是否回文。
由定义可知，i <= j。

递推公式：
当s[j] == s[i]时，res = max(res, j-i+1)：
	- 若i == j 或 i + 1 == j：dp[i][j] = true；
	- 否则：dp[i][j] = dp[i + 1][j - 1]。
当s[j] != s[i]时，dp[i][j] = false，即保留初始值。

初始化：
默认均不是回文字符串。

遍历方向：
由递推公式可知，从斜下方推导。
故应从下到上，由左到右。
注意，因为j >= i，所以j的遍历从i开始，而不是从0开始。
*/
