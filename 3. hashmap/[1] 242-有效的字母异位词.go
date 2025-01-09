package hashmap

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的 字母异位词。
//
// 示例 1:
//
// 输入: s = "anagram", t = "nagaram"
// 输出: true
//
// 示例 2:
//
// 输入: s = "rat", t = "car"
// 输出: false
//
// 提示:
//
// 1 <= s.length, t.length <= 5 * 10⁴
// s 和 t 仅包含小写字母
//
// 进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
func isAnagram(s string, t string) bool {
	cntMap := make(map[int32]int)
	for _, ch := range s {
		cntMap[ch]++
	}
	for _, ch := range t {
		cntMap[ch]--
	}
	for _, cnt := range cntMap {
		if cnt != 0 {
			return false
		}
	}
	return true
}

func isAnagramII(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := [26]int{}
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}
	for _, cnt := range counts {
		if cnt != 0 {
			return false
		}
	}
	return true
}

/**
【ab, a】不互为字母异位词！
*/
