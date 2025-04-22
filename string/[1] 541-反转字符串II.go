package string

import "strings"

// 给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
//
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
//
// 示例 1：
//
// 输入：s = "abcdefg", k = 2
// 输出："bacdfeg"
//
// 示例 2：
//
// 输入：s = "abcd", k = 2
// 输出："bacd"
//
// 提示：
//
// 1 <= s.length <= 10⁴
// s 仅由小写英文组成
// 1 <= k <= 10⁴
func reverseStr(s string, k int) string {
	for i := 0; i < len(s); i += 2 * k {
		if i+k > len(s) {
			s = reverse(s, i, len(s)-1)
		} else {
			s = reverse(s, i, i+k-1)
		}
	}
	return s
}

func reverse(s string, beginIdx, endIdx int) string {
	chars := []byte(s)
	for beginIdx < endIdx {
		chars[beginIdx], chars[endIdx] = chars[endIdx], chars[beginIdx]
		beginIdx++
		endIdx--
	}
	var builder strings.Builder
	for _, ch := range chars {
		builder.WriteByte(ch)
	}
	return builder.String()
}

/**
chars := []byte(s)：将原字符串复制到字符数组中，以便通过反转数组来反转字符串；
strings.Builder：代替直接对不可变的string进行修改，提高性能；
合并if-else分支：
	for i := 0; i < len(s); i += 2 * k {
		if i+2*k < len(s) { // 每隔 2k 个字符的前 k 个字符进行反转
			s = reverse(s, i, i+k-1)
		} else if i+k > len(s) { // 剩余字符少于 k 个，则将剩余字符全部反转
			s = reverse(s, i, len(s)-1)
		} else { // 剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符
			s = reverse(s, i, i+k-1)
		}
	}
*/
