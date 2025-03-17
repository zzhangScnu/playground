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
func minDistanceII(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return l1 + l2 - dp[l1][l2]*2
}

/**
思路：
求出两个字符串的最长公共子序列长度，则两个字符总长度减去其两倍就是最少的删除操作。
*/
