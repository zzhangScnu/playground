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
			if i == j || i+1 == j || dp[i+1][j-1] { // 长度为1（因为没有提前做 dp 值的初始化，所以要额外兼容此分支） / 长度为2，回文的基础 / 长度大于2，看子串是否为回文串
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

/*
看递推公式，i, j 由 i + 1, j - 1 推导出来。
由下方的关系图可知，左斜下方 -> 右斜上方。
所以遍历方向为：i 从下到上，j 从左到右。

——————————————————————————————————————————
｜    i, j - 1 |     i, j |     i, j + 1 ｜
｜i + 1, j - 1 | i + 1, j | i + 1, j + 1 ｜
——————————————————————————————————————————
*/

/*
这个做法，从递推公式就错了。这种写法是求最长回文子序列，而不是子串的。
最长回文子串（连续）
abcba → bcb 是回文串

最长回文子序列（不强制连续）
abca → aba 或 aca 是回文序列
*/
func longestPalindromeI(s string) string {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 1
		}
	}
	var res, begin, end int
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				if i == j-1 {
					dp[i][j] = 2
				} else {
					// 只要两头相等，就把中间最长回文的长度 +2，不管中间是不是连续的。
					// 这是回文子序列的公式，天然会跳过字符，所以不能用来求连续子串。
					dp[i][j] = dp[i+1][j-1] + 2
				}
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
			if res < dp[i][j] {
				res = max(res, dp[i][j])
				begin, end = i, j
			}
		}
	}
	return s[begin : end+1]
}

/*
这才是求最长回文字串的解法
*/
func longestPalindromeIII(s string) string {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ { // 一个字符本身就是一个回文串
			dp[i][j] = 1
		}
	}
	var res, begin, end int
	for i := n - 1; i >= 0; i-- { // i 从下到上
		for j := i + 1; j < n; j++ { // j 从左到右，且由于子串的定义， j >= i + 1
			if s[i] != s[j] { // 如两头不相等，就不是回文串，将 dp 值置零
				dp[i][j] = 0
				continue
			}
			if i == j-1 { // 如两头相等且长度为2，回文的基础
				dp[i][j] = 2
			} else {
				if dp[i+1][j-1] != 0 { // 如子串是回文串，则将当前字符串的回文串长度在子串的基础上+2
					dp[i][j] = dp[i+1][j-1] + 2
				} else { // 如子串不是回文串，则将当前字符串的回文串长度置零
					dp[i][j] = 0
				}
			}
			if res < dp[i][j] { // 处理本轮 dp 后，更新最大回文串长度及对应索引
				res = max(res, dp[i][j])
				begin, end = i, j
			}
		}
	}
	return s[begin : end+1]
}
