package array

// 给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
//
// 回文字符串 是正着读和倒过来读一样的字符串。
//
// 子字符串 是字符串中的由连续字符组成的一个序列。
//
// 示例 1：
//
// 输入：s = "abc"
// 输出：3
// 解释：三个回文子串: "a", "b", "c"
//
// 示例 2：
//
// 输入：s = "aaa"
// 输出：6
// 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
//
// 提示：
//
// 1 <= s.length <= 1000
// s 由小写英文字母组成
func countSubstrings(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		res += doCount(s, i, i)
		res += doCount(s, i, i+1)
	}
	return res
}

func doCount(s string, i, j int) int {
	var res int
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i, j = i-1, j+1
		res++
	}
	return res
}

/**
思路：
对于每一个回文串，有两种情况：
奇数中心；
偶数中心。
对于这两种情况分别向外扩展。
*/
