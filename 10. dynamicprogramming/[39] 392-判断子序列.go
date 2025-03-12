package dynamicprogramming

// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而
// "aec"不是）。
//
// 进阶：
//
// 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代
// 码？
//
// 示例 1：
//
// 输入：s = "abc", t = "ahbgdc"
// 输出：true
//
// 示例 2：
//
// 输入：s = "axc", t = "ahbgdc"
// 输出：false
//
// 提示：
//
// 0 <= s.length <= 100
// 0 <= t.length <= 10^4
// 两个字符串都只由小写字符组成。
func isSubsequence(s string, t string) bool {
	sl, tl := len(s), len(t)
	dp := make([][]int, sl+1)
	for i := 0; i < sl; i++ {
		dp[i] = make([]int, tl+1)
	}
	for i := 1; i <= sl; i++ {
		for j := 1; j <= tl; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[sl][tl] == len(s)
}
