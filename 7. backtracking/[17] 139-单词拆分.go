package backtracking

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

var memo map[int]bool

func wordBreak(s string, wordDict []string) bool {
	memo = make(map[int]bool)
	wordMap := make(map[string]interface{})
	for _, word := range wordDict {
		wordMap[word] = true
	}
	return doWordBreak(s, wordMap, 0)
}

func doWordBreak(s string, wordMap map[string]interface{}, start int) bool {
	if start >= len(s) {
		return true
	}
	if flag, ok := memo[start]; ok {
		return flag
	}
	for end := start; end < len(s); end++ {
		subs := s[start : end+1]
		if _, ok := wordMap[subs]; ok && doWordBreak(s, wordMap, end+1) {
			memo[start] = true
			return true
		}
	}
	memo[start] = false
	return false
}

/**
思路：
用两个游标不断切割字符串s，用子串跟字典比对。
方法的入参只需要一个起始索引即可，在方法中通过for循环来遍历所有可能的结束索引。
如果找到一个子字符串是符合要求的，就从结束索引的下一个位置再去递归查找。

因为只需要找到任意一种可能性，所以在搜索树上只要遍历到其中一条符合要求的边，就立即返回。
回溯过程逐层将true返回，直至主函数。
base case：
当游标到达s末尾都没有提前返回false，表示到达叶子节点且找到了这条边。

如果当前子字符串在字典内，且递归往后的子字符串们也满足条件，则写入备忘录且返回true。
否则当游标遍历完s[index:]，就直接返回false。

此外，因为会存在大量重复计算，所以需要引入备忘录。

备忘录：
用于记录从某个索引index开始，对该位置后的字符串s[index:]就已经处理过一轮递归&回溯了，意味着其所有子问题都有解了，
从index开始的s[start:end]子字符串已经全部判断过了。
而不是对某个s[start:end]子字符串进行记录，这样递归深度依然很大，超时风险高。

因为需要通过【if flag, ok := memo[start]; ok】来判断是否有写入过备忘录，
所以每次得出结果后都需要显式记录一下结果，包括false。
*/
