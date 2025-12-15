package slidingwindow

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
//
// 示例 1:
//
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
// 示例 2:
//
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
// 示例 3:
//
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
// 请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
//
// 提示：
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
func lengthOfLongestSubstring(s string) int {
	var i, res int
	cnt := make(map[string]int)
	for j := 0; j < len(s); j++ {
		ch := string(s[j])
		cnt[ch]++
		for cnt[ch] > 1 {
			cnt[string(s[i])]--
			i++
		}
		if j-i+1 > res {
			res = j - i + 1
		}
	}
	return res
}

/**
关于更新结果的时机：
跟前面的几道题一样，核心是要在【满足条件约束的时候更新结果】
所以这道题是在左侧窗口收缩完毕，【没有重复字符】的时候再更新，即在for循环以后。
跟前面的滑动窗口题相比，表现不同，但本质一致。
*/

func lengthOfLongestSubstringII(s string) int {
	l, r, res := 0, 0, 0
	counter := make(map[byte]interface{})
	for ; r < len(s); r++ {
		for _, ok := counter[s[r]]; ok; _, ok = counter[s[r]] { // 基于本字符进入窗口，收缩左边界以至无重复字符
			delete(counter, s[l])
			l++
		}
		res = max(res, r-l+1) // 满足窗口中无重复元素的条件，收割结果
		counter[s[r]] = struct{}{}
	}
	return res
}
