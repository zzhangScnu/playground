package dynamicprogramming

// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文串。
//
// 返回符合要求的 最少分割次数 。
//
// 示例 1：
//
// 输入：s = "aab"
// 输出：1
// 解释：只需一次分割就可将s 分割成 ["aa","b"] 这样两个回文子串。
//
// 示例 2：
//
// 输入：s = "a"
// 输出：0
//
// 示例 3：
//
// 输入：s = "ab"
// 输出：1
//
// 提示：
//
// 1 <= s.length <= 2000
// s 仅由小写英文字母组成
func minCut(s string) int {
	n := len(s)
	memo := make([][]bool, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]bool, n)
		memo[i][i] = true
		if i > 0 && s[i-1] == s[i] {
			memo[i-1][i] = true
		}
	}
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			memo[i][j] = memo[i+1][j-1] && s[i] == s[j]
		}
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = i
	}
	for i := 0; i < n; i++ {
		if memo[0][i] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if memo[j+1][i] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

/**
DP数组及下标含义：
- i：字符串范围[0...i]
- dp[i]：s[0...i]中回文子串的最少切割次数。


递推公式：
for i [0, n)
for j [0, i)
if s[j + 1 ... i] == 回文子串
	dp[i] = min(dp[i], dp[j] + 1)
即：找到了一个分割位置，使得s[0 ... j]和s[j + 1 ... i]都为回文子串。
此时切割次数：前者 = dp[j]，后者 = 1
因每次都需要判断一个子串是否回文，需要引入备忘录。


初始化：
对于长度为k的字符串，最大的切割次数为k，即每个字符作为一个回文子串。
(简单粗暴初始化为math.MaxInt也可以)
后续在比较中不断择优取最小值。

遍历方向：
从左到右。
*/
