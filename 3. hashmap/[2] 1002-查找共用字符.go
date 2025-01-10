package hashmap

import "math"

// 给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答案
// 。
//
// 示例 1：
//
// 输入：words = ["bella","label","roller"]
// 输出：["e","l","l"]
//
// 示例 2：
//
// 输入：words = ["cool","lock","cook"]
// 输出：["c","o"]
//
// 提示：
//
// 1 <= words.length <= 100
// 1 <= words[i].length <= 100
// words[i] 由小写英文字母组成
func commonChars(words []string) []string {
	baseCounts := [26]int{}
	for _, ch := range words[0] {
		baseCounts[ch-'a']++
	}
	for i := 1; i < len(words); i++ {
		tmpCounts := [26]int{}
		for _, ch := range words[i] {
			tmpCounts[ch-'a']++
		}
		for ch, count := range tmpCounts {
			baseCounts[ch-'a'] = int(math.Min(float64(count), float64(baseCounts[ch-'a'])))
		}
	}
	var res []string
	for ch, count := range baseCounts {
		for i := 0; i < count; i++ {
			res = append(res, string(rune(ch)))
		}
	}
	return res
}
