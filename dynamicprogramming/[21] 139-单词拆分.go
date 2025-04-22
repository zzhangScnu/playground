package dynamicprogramming

// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
//
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
//
// 示例 1：
//
// 输入: s = "leetcode", wordDict = ["leet", "code"]
// 输出: true
// 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
//
// 示例 2：
//
// 输入: s = "applepenapple", wordDict = ["apple", "pen"]
// 输出: true
// 解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
//
// 注意，你可以重复使用字典中的单词。
//
// 示例 3：
//
// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出: false
//
// 提示：
//
// 1 <= s.length <= 300
// 1 <= wordDict.length <= 1000
// 1 <= wordDict[i].length <= 20
// s 和 wordDict[i] 仅由小写英文字母组成
// wordDict 中的所有字符串 互不相同
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for j := 1; j <= len(s); j++ {
		for i := 0; i < len(wordDict); i++ {
			if j >= len(wordDict[i]) {
				dp[j] = dp[j] || dp[j-len(wordDict[i])] && wordDict[i] == s[j-len(wordDict[i]):j]
			}
		}
	}
	return dp[len(s)]
}

/**
DP数组及下标含义
- j：字符串长度为j，即背包容量为j；
- dp[j]：对物品[0, i]取/不取时，目标字符串长度为j时，是否可以被字典中的子串凑成。

递推公式
dp[j] = dp[j] || dp[j-len(wordDict[i])] && wordDict[i] == s[j-len(wordDict[i]):j]
背包容量为j，即当前遍历到字符串s长度为j的位置，s[:j-1]。
(至于什么截取的索引是[j-len(wordDict[i]), j]，因为Go中字符串截取是前闭后开，且j为背包容量即为字符串长度，所以j-1才是实际的结束索引)
对字典中的每个单词进行遍历，如果单词i能塞进容量为j的背包，且与s末尾的子字符串相等，则表示单词i能作为s的一部分参与拼接；
此时如果前面的字符串也能被字典中的子串凑成，则整个字符串满足条件。

为什么要跟dp[j]取或？
设想一种情况：背包容量为j时，
- 先检查字典中的单词i，能组成目标字符串，此时dp[j] = true；
- i++，继续检查下一个单词i+1，但无法组成目标字符串，此时dp[j] = false；
- 但因为不断地在一维数组上原地与先前结果取或，所以只要有一个单词结果为true，dp[j] = true。
(看起来，只要任意一个单词结果为true，就不再往后遍历单词，而是开始遍历下一个背包容量，也能达到效果，还能节省计算)

初始化
dp[0] = true
实际没有意义，只是为了正确递推。

推导方向
背包和物品都是从小到大遍历；
由于本题是求排列，顺序敏感，所以需要先遍历背包，再遍历物品。
*/
