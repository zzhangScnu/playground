package dynamicprogramming

// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
// 示例 1：
//
// 输入：s = "(()"
// 输出：2
// 解释：最长有效括号子串是 "()"
//
// 示例 2：
//
// 输入：s = ")()())"
// 输出：4
// 解释：最长有效括号子串是 "()()"
//
// 示例 3：
//
// 输入：s = ""
// 输出：0
//
// 提示：
//
// 0 <= s.length <= 3 * 10⁴
// s[i] 为 '(' 或 ')'

func longestValidParentheses(s string) int {
	longest, left, right := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			right++
		}
		if left == right {
			longest = max(longest, left+right)
		} else if left < right {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			right++
		}
		if left == right {
			longest = max(longest, left+right)
		} else if right < left {
			left, right = 0, 0
		}
	}
	return longest
}

func longestValidParenthesesII(s string) int {
	var longest int
	n := len(s)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			dp[i] = 2
			if i >= 2 {
				dp[i] += dp[i-2]
			}
		} else if s[i-1] == ')' && i >= dp[i-1]+1 && s[i-dp[i-1]-1] == '(' {
			dp[i] = dp[i-1] + 2
			if i >= dp[i-1]+2 {
				dp[i] += dp[i-dp[i-1]-2]
			}
		}
		longest = max(longest, dp[i])
	}
	return longest
}
