package array

// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的 排列。如果是，返回 true ；否则，返回 false 。
//
// 换句话说，s1 的排列之一是 s2 的 子串 。
//
// 示例 1：
//
// 输入：s1 = "ab" s2 = "eidbaooo"
// 输出：true
// 解释：s2 包含 s1 的排列之一 ("ba").
//
// 示例 2：
//
// 输入：s1= "ab" s2 = "eidboaoo"
// 输出：false
//
// 提示：
//
// 1 <= s1.length, s2.length <= 10⁴
// s1 和 s2 仅包含小写字母
func checkInclusion(s1 string, s2 string) bool {
	targetCnt, windowCnt := make(map[string]int), make(map[string]int)
	for _, ch := range s1 {
		targetCnt[string(ch)]++
	}
	var i, valid int
	for j, ch := range s2 {
		rc := string(ch)
		if targetCnt[rc] > 0 {
			windowCnt[rc]++
			if windowCnt[rc] == targetCnt[rc] {
				valid++
			}
		}
		for len(targetCnt) == valid {
			if j-i+1 == len(s1) {
				return true
			}
			lc := string(s2[i])
			if targetCnt[lc] > 0 {
				if windowCnt[lc] == targetCnt[lc] {
					valid--
				}
			}
			windowCnt[lc]--
			i++
		}
	}
	return false
}

/**
这题的核心在于，如何判断满足题目中的约束条件：
当窗口中有s1所有字符，且窗口长度也等于s1
*/
