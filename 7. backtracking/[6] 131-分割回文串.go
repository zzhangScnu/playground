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

/**
beginIdx作用：字符串分割线
- for循环中：在本层控制从哪个字符开始切割，i向右扩展表示切割到哪个字符，以此生成【上层切割到beginIdx-1后】再继续往后切割生成的树层；
- 递归调用中：以i+1作为入参，在下层控制从哪个字符开始切割，以此生成【上层切割到i后】再继续往后切割生成的树枝；
所以在本层中，切割字符的长度为【i-beginIdx+1】。
可以在本层就判断该子字符串是否回文串，是的话就加入单次结果集并进行递归&回溯，否则直接跳过不处理。

结束条件：字符串分割线扫完了整个字符串
即 【分割线】>=【字符串长度】
*/
