package array

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
// 注意：
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
// 示例 1：
//
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//
// 示例 2：
//
// 输入：s = "a", t = "a"
// 输出："a"
// 解释：整个字符串 s 是最小覆盖子串。
//
// 示例 3:
//
// 输入: s = "a", t = "aa"
// 输出: ""
// 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
// 因此没有符合条件的子字符串，返回空字符串。
//
// 提示：
//
// m == s.length
// n == t.length
// 1 <= m, n <= 10⁵
// s 和 t 由英文字母组成
//
// 进阶：你能设计一个在
// o(m+n) 时间内解决此问题的算法吗？
func minWindow(s string, t string) string {
	var i, validNum int
	targetCnt, windowCnt := make(map[string]int), make(map[string]int)
	for _, char := range t {
		targetCnt[string(char)]++
	}
	ri, rl := 0, 100_001
	for j := 0; j < len(s); j++ {
		rc := string(s[j])
		if targetCnt[rc] > 0 {
			windowCnt[rc]++
			if windowCnt[rc] == targetCnt[rc] {
				validNum++
			}
			for validNum == len(targetCnt) {
				if j-i+1 < rl {
					ri = i
					rl = j - i + 1
				}
				lc := string(s[i])
				if targetCnt[lc] > 0 {
					if windowCnt[lc] == targetCnt[lc] {
						validNum--
					}
					windowCnt[lc]--
				}
				i++
			}
		}
	}
	if rl == 100_001 {
		return ""
	}
	return s[ri : ri+rl]
}

/**
- 字符串的处理：
	- 遍历string中的每一位：
		- for i：字节切片，可能无法正确处理非ASCII码字符；
 		- for range：Unicode码点。
	- 转换为string：string(a)
- 增大窗口合减小窗口的操作，是对称的；
- validNum其实就是个快捷方式，否则就需要遍历t中的每一个字符，一一去看窗口中是否已经有满足条件的所有字符了。
*/
