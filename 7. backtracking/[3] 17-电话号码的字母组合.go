package backtracking

import "strings"

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
// 示例 1：
//
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
// 示例 2：
//
// 输入：digits = ""
// 输出：[]
//
// 示例 3：
//
// 输入：digits = "2"
// 输出：["a","b","c"]
//
// 提示：
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。

var letterPath []string

var letterCombinationRes []string

var mapping = map[uint8][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	letterPath = []string{}
	letterCombinationRes = []string{}
	if digits == "" {
		return letterCombinationRes
	}
	doLetterCombinations(0, digits)
	return letterCombinationRes
}

func doLetterCombinations(digitIdx int, digits string) {
	if len(letterPath) == len(digits) {
		letterCombinationRes = append(letterCombinationRes, strings.Join(letterPath, ""))
		return
	}
	digit := digits[digitIdx]
	characters := mapping[digit]
	for i := 0; i < len(characters); i++ {
		letterPath = append(letterPath, characters[i])
		doLetterCombinations(digitIdx+1, digits)
		letterPath = letterPath[:len(letterPath)-1]
	}
}
