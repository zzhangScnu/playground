package greedy

// 给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。例如，字符串 "ababcc" 能够被分为 ["abab",
// "cc"]，但类似 ["aba", "bcc"] 或 ["ab", "ab", "cc"] 的划分是非法的。
//
// 注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。
//
// 返回一个表示每个字符串片段的长度的列表。
//
// 示例 1：
//
// 输入：s = "ababcbacadefegdehijhklij"
// 输出：[9,7,8]
// 解释：
// 划分结果为 "ababcbaca"、"defegde"、"hijhklij" 。
// 每个字母最多出现在一个片段中。
// 像 "ababcbacadefegde", "hijhklij" 这样的划分是错误的，因为划分的片段数较少。
//
// 示例 2：
//
// 输入：s = "eccbbbbdec"
// 输出：[10]
//
// 提示：
//
// 1 <= s.length <= 500
// s 仅由小写英文字母组成
func partitionLabels(s string) []int {
	pos := make(map[int32]int)
	for i, ch := range s {
		pos[ch] = i
	}
	var res []int
	subStart, subEnd := 0, 0
	for i, ch := range s {
		subEnd = max(subEnd, pos[ch])
		if i == subEnd {
			res = append(res, subEnd-subStart+1)
			subStart = subEnd + 1
		}
	}
	return res
}

/**
思路：
先将每个字符出现过的最远下标记录下来，作为可能的分割边界；
再用一个游标遍历字符串。遇到每一个字符，都判断一下它们的边界是不是出现过的字符中最远的，是的话就更新该子串的分割边界。

什么时候进行切割，开始下一个字符？
当游标遍历到子串的分割边界时，意味着所有字符都圈定成功，且保证同一字符都圈在同一子串中。
*/
