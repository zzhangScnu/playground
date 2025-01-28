package backtracking

// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 示例 1：
//
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
//
// 示例 2：
//
// 输入：s = "a"
// 输出：[["a"]]
//
// 提示：
//
// 1 <= s.length <= 16
// s 仅由小写英文字母组成

var substring []string

var partitions [][]string

func partition(s string) [][]string {
	substring, partitions = []string{}, [][]string{}
	doPartition(0, s)
	return partitions
}

func doPartition(beginIdx int, s string) {
	if beginIdx >= len(s) {
		res := make([]string, len(substring))
		copy(res, substring)
		partitions = append(partitions, res)
		return
	}
	for i := beginIdx; i < len(s); i++ {
		if isPalindrome(s, beginIdx, i) {
			substring = append(substring, s[beginIdx:i+1])
			doPartition(i+1, s)
			substring = substring[:len(substring)-1]
		}
	}
}

func isPalindrome(str string, beginIdx, endIdx int) bool {
	for i, j := beginIdx, endIdx; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
