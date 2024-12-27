package array

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
// 示例 1:
//
// 输入: s = "cbaebabacd", p = "abc"
// 输出: [0,6]
// 解释:
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
//
// 示例 2:
//
// 输入: s = "abab", p = "ab"
// 输出: [0,1,2]
// 解释:
// 起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
// 起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
// 起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
//
// 提示:
//
// 1 <= s.length, p.length <= 3 * 10⁴
// s 和 p 仅包含小写字母
func findAnagrams(s string, p string) []int {
	windowCnt, targetCnt := make(map[string]int), make(map[string]int)
	for _, ch := range p {
		targetCnt[string(ch)]++
	}
	var i, valid int
	var res []int
	for j := 0; j < len(s); j++ {
		rc := string(s[j])
		if targetCnt[rc] > 0 {
			windowCnt[rc]++
			if windowCnt[rc] == targetCnt[rc] {
				valid++
			}
		}
		for valid == len(targetCnt) {
			if j-i+1 == len(p) {
				res = append(res, i)
			}
			lc := string(s[i])
			if targetCnt[lc] > 0 {
				if windowCnt[lc] == targetCnt[lc] {
					valid--
				}
				windowCnt[lc]--
			}
			i++
		}
	}
	return res
}
