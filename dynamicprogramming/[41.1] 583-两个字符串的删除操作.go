package dynamicprogramming

// 给定两个单词 word1 和
// word2 ，返回使得
// word1 和
// word2 相同所需的最小步数。
//
// 每步 可以删除任意一个字符串中的一个字符。
//
// 示例 1：
//
// 输入: word1 = "sea", word2 = "eat"
// 输出: 2
// 解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
//
// 示例 2:
//
// 输入：word1 = "leetcode", word2 = "etco"
// 输出：4
//
// 提示：
//
// 1 <= word1.length, word2.length <= 500
// word1 和 word2 只包含小写英文字母
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 0; i <= l1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= l2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1)
			}
		}
	}
	return dp[l1][l2]
}

/**
思路：跟上一题一样，但本次可以同时在两个字符串中进行"删除"即跳过操作。

DP数组及下标含义：
- i：以word1[i - 1]结尾的字符串，即当前游标指向i - 1；
- j：以word2[j - 1]结尾的字符串，即当前游标指向j - 1；
- dp[i][j]：在以word1[i - 1]结尾的字符串和以word2[j - 1]结尾的字符串中，要使得两个字符串相同，所需的最少删除次数。

递推公式：
if word1[i-1] == word2[j-1] { // 如果当前字符相等，说明这一步不用操作删除，直接继承两个字符串前一位的次数
	dp[i][j] = dp[i-1][j-1]
} else {
	dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1) // 否则有3种情况需要考虑
}
dp[i][j] =
1. 删除word1中的字符：dp[i - 1][j] + 1，word1的游标从i - 1 -> i，删除次数+1；
2. 删除word2中的字符：dp[i][j - 1] + 1，word2的游标从j - 1 -> j，删除次数+1；
3. 两个字符串都需要删除字符：dp[i - 1][j - 1] + 2，word1和word2的游标同时推进，删除次数+2。
   而dp[i - 1][j - 1] + 2 == dp[i - 1][j] + 1 == dp[i][j - 1] + 1，即先删一个，再删一个的次数。
取三者次数最少的，作为本轮结果。

初始化：
当word1 == ""时，word2需删除的步数就等于游标本身。如word2 == ""时，则为0步。
当word2 == ""时，word1需删除的步数就等于游标本身。

遍历方向：
从左到右。
*/
