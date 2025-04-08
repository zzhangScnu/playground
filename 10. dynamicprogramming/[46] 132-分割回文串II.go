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

同时，在遍历过程中，如果发现[0 ... i]本身就是回文串，则最小分割次数为1，
保持默认值即可，也无需遍历j来试图找到最小切割次数。


初始化：
对于长度为k的字符串，最大的切割次数为k，即每个字符作为一个回文子串。
(简单粗暴初始化为math.MaxInt也可以)
后续在比较中不断择优取最小值。

遍历方向：
从左到右。
*/

/**
如何判断回文串并初始化备忘录：


一开始的做法时间复杂度O(n^3)：

isPalindrome := func(s string, start, end int) bool {
	for i, j := beginIdx, endIdx; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

for i := 0; i < n; i++ {
	for j := 0; j < n; j++ {
		memo[i][j] = isPalindrome(s, i, j)
	}
}

但有一种简单做法，是通过状态转移递推而来，时间复杂度为O(n^2)：
memo := make([][]bool, n)
for i := 0; i < n; i++ {
	memo[i] = make([]bool, n)
	memo[i][i] = true // 单个字符本身构成回文串
	if i > 0 && s[i-1] == s[i] { // 长度为2且字符相同的字符串构成回文串
		memo[i-1][i] = true
	}
}
// 这里先固定字符串长度length
// 再对字符串的每一个起点i，计算出终点j = i + length - 1
// 判断s[i ... j]是否回文串
// 为什么要这么做？因为长度为length的字符串是否回文子串，依赖于长度为length-2的字符串是否回文子串递推而来
// 即长度较小 -> 长度较大，所以需要循环递增length，每次固定length，处理完所有可能的字符串，再对length++，继续处理下一批可能的字符串
for length := 3; length <= n; length++ { // 因为长度为1和2的子字符串已初始化，这里从长度为3开始
	for i := 0; i <= n-length; i++ { // 因为i和j都需要在s的长度范围内，而j = i + length - 1 < n，即i < n - length + 1
		j := i + length - 1
		memo[i][j] = memo[i+1][j-1] && s[i] == s[j] // s[i ... j]是否回文串，可以拆解为子问题：s[i + 1 ... j - 1]是否回文串 && s[i]和s[j]是否相等
	}
}
*/
