package array

import "strings"

// 编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
// 示例 1：
//
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
//
// 示例 2：
//
// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。
//
// 提示：
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 如果非空，则仅由小写英文字母组成
func longestCommonPrefix(strs []string) string {
	var sb strings.Builder
	var ch byte
	for i := 0; i < len(strs[0]); i++ {
		ch = strs[0][i]
		for _, str := range strs[1:] {
			if i >= len(str) || str[i] != ch {
				return sb.String()
			}
		}
		sb.WriteByte(ch)
	}
	return sb.String()
}
