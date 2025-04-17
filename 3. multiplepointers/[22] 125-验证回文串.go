package multiplepointers

import "unicode"

// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
//
// 字母和数字都属于字母数字字符。
//
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入: s = "A man, a plan, a canal: Panama"
// 输出：true
// 解释："amanaplanacanalpanama" 是回文串。
//
// 示例 2：
//
// 输入：s = "race a car"
// 输出：false
// 解释："raceacar" 不是回文串。
//
// 示例 3：
//
// 输入：s = " "
// 输出：true
// 解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
// 由于空字符串正着反着读都一样，所以是回文串。
//
// 提示：
//
// 1 <= s.length <= 2 * 10⁵
// s 仅由可打印的 ASCII 字符组成
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isNumOrChar(s[left]) {
			left++
		}
		for left < right && !isNumOrChar(s[right]) {
			right--
		}
		if unicode.ToLower(rune(s[left])) == unicode.ToLower(rune(s[right])) {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func isNumOrChar(ch byte) bool {
	char := rune(ch)
	return unicode.IsDigit(char) || unicode.IsLetter(char)
}

func isNumOrCharByASCII(char byte) bool {
	if char >= 48 && char <= 57 {
		return true
	}
	if (char >= 65 && char <= 90) || (char >= 97 && char <= 122) {
		return true
	}
	return false
}

/**
- byte -> rune：当是ASCII字符时，可直接强转。内置unicode包中有一系列判断和转换的方法，开箱即用，不用记忆ASCII编码；
- 边界条件控制：for left < right && !isNumOrChar(s[left])。一开始写漏了【left < right】这个条件，会导致找不到结果的场景下数组越界。
还是不够细心啊～
*/
