package dynamicprogramming

// 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
//
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
//
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
//
// 示例 1：
//
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace" ，它的长度为 3 。
//
// 示例 2：
//
// 输入：text1 = "abc", text2 = "abc"
// 输出：3
// 解释：最长公共子序列是 "abc" ，它的长度为 3 。
//
// 示例 3：
//
// 输入：text1 = "abc", text2 = "def"
// 输出：0
// 解释：两个字符串没有公共子序列，返回 0 。
//
// 提示：
//
// 1 <= text1.length, text2.length <= 1000
// text1 和 text2 仅由小写英文字符组成。
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

/**
DP数组及下标含义：
- i：text1中以下标i-1结尾的子序列；
- j：text2中以下标j-1结尾的子序列；
- dp[i][j]：以上2个子序列中最长公共子序列的长度。

递推公式：
if text1[i-1] == text2[j-1]
dp[i][j] = dp[i-1][j-1]+1
否则应比较：
(text1[0, i-2], text2[0, j-1])
(text1[0, i-1], text2[0, j-2])
这两种组合的最长公共子序列长度，取最大值。

返回值：dp[m][n]
因为一定要将2个字符串都遍历完，才会得到最终结果。

初始化：
第一行、第一列没有实际语义，仅为了累加推导结果。故保持默认零值即可。

遍历方向：从左到右。
*/
