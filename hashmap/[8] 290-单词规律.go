package hashmap

import "strings"

// 给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。
//
// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。
//
// 示例1:
//
// 输入: pattern = "abba", s = "dog cat cat dog"
// 输出: true
//
// 示例 2:
//
// 输入:pattern = "abba", s = "dog cat cat fish"
// 输出: false
//
// 示例 3:
//
// 输入: pattern = "aaaa", s = "dog cat cat dog"
// 输出: false
//
// 提示:
//
// 1 <= pattern.length <= 300
// pattern 只包含小写英文字母
// 1 <= s.length <= 3000
// s 只包含小写英文字母和 ' '
// s 不包含 任何前导或尾随对空格
// s 中每个单词都被 单个空格 分隔
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	char2word, word2char := make(map[uint8]string), make(map[string]uint8)
	for i := 0; i < len(pattern); i++ {
		char, word := pattern[i], words[i]
		if w, ok := char2word[char]; ok && w != word {
			return false
		} else if c, ok := word2char[word]; ok && c != char {
			return false
		}
		char2word[char] = word
		word2char[word] = char
	}
	return true
}

/**
思路：
需做双向映射，如果在遍历过程中一一双向匹配成功，则表示模式和单词遵循相同规律。
注意需做双向映射，否则会导致[a, b, b, c] -> [cat, dog, dog, cat]也被误判为同一规律。
*/
