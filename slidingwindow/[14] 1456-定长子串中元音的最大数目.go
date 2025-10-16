package slidingwindow

// 给你字符串 s 和整数 k 。
//
// 请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数。
//
// 英文中的 元音字母 为（a, e, i, o, u）。
//
// 示例 1：
//
// 输入：s = "abciiidef", k = 3
// 输出：3
// 解释：子字符串 "iii" 包含 3 个元音字母。
//
// 示例 2：
//
// 输入：s = "aeiou", k = 2
// 输出：2
// 解释：任意长度为 2 的子字符串都包含 2 个元音字母。
//
// 示例 3：
//
// 输入：s = "leetcode", k = 3
// 输出：2
// 解释："lee"、"eet" 和 "ode" 都包含 2 个元音字母。
//
// 示例 4：
//
// 输入：s = "rhythms", k = 4
// 输出：0
// 解释：字符串 s 中不含任何元音字母。
//
// 示例 5：
//
// 输入：s = "tryhard", k = 4
// 输出：1
//
// 提示：
//
// 1 <= s.length <= 10^5
// s 由小写英文字母组成
// 1 <= k <= s.length
func maxVowels(s string, k int) int {
	n := len(s)
	if n < k {
		return 0
	}
	vowels := map[rune]interface{}{
		'a': struct{}{},
		'e': struct{}{},
		'i': struct{}{},
		'o': struct{}{},
		'u': struct{}{},
	}
	maxCount, curCount := 0, 0
	left, right := 0, 0
	for right < n {
		if _, ok := vowels[rune(s[right])]; ok {
			curCount++
		}
		if right-left+1 == k {
			maxCount = max(maxCount, curCount)
			if _, ok := vowels[rune(s[left])]; ok {
				curCount--
			}
			left++
		}
		right++
	}
	return maxCount
}

/**
思路：
采用滑动窗口的解法。
在定长的滑动窗口中，不断扩大右边界，
在满足长度的条件下，与全局结果取优，不断缩小左边界。
判断出 / 入的字符是否元音，计算当前窗口中的元音数目。

注意，不能将对右边界的判断放在for循环最开始，这样会导致在窗口大小满足长度要求时，curCount的计数未包括right。
*/
