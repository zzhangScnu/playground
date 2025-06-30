package dynamicprogramming

// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作：
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
// 示例 1：
//
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
//
// 示例 2：
//
// 输入：word1 = "intention", word2 = "execution"
// 输出：5
// 解释：
// intention -> inention (删除 't')
// inention -> enention (将 'i' 替换为 'e')
// enention -> exention (将 'n' 替换为 'x')
// exention -> exection (将 'n' 替换为 'c')
// exection -> execution (插入 'u')
//
// 提示：
//
// 0 <= word1.length, word2.length <= 500
// word1 和 word2 由小写英文字母组成
func minDistance72II(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	var dp func(word1, word2 string, i, j int) int
	dp = func(word1, word2 string, i, j int) int {
		if i == m {
			return n - j
		}
		if j == n {
			return m - i
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if word1[i] == word2[j] {
			memo[i][j] = dp(word1, word2, i+1, j+1)
		} else {
			memo[i][j] = min(
				dp(word1, word2, i+1, j+1),
				dp(word1, word2, i+1, j),
				dp(word1, word2, i, j+1),
			) + 1
		}
		return memo[i][j]
	}
	return dp(word1, word2, 0, 0)
}
