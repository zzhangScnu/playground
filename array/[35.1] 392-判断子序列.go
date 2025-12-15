package array

// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而
// "aec"不是）。
//
// 进阶：
//
// 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代
// 码？
//
// 示例 1：
//
// 输入：s = "abc", t = "ahbgdc"
// 输出：true
//
// 示例 2：
//
// 输入：s = "axc", t = "ahbgdc"
// 输出：false
//
// 提示：
//
// 0 <= s.length <= 100
// 0 <= t.length <= 10^4
// 两个字符串都只由小写字符组成。
func isSubsequence(s string, t string) bool {
	si, ti := 0, 0
	for ; ti < len(t) && si < len(s); ti++ {
		if s[si] == t[ti] {
			si++
		}
	}
	return si == len(s)
}

/**
思路：
双指针
si -> s: abc
ti -> t: ahbgdc

ti不断往前推进；
如果s[si] == t[ti]，说明找到了一个匹配字符，si往前推进。
最终看si是否能到达s的末尾。

就像是消消乐，匹配一个就消去一个，最终看s是否能被完全消去。

注意：
for ; ti < len(t) && si < len(s); ti++
【si < len(s)】必不可少，否则当t结束前就遍历完了s，会出现数组越界。
加了这个控制，能兼容s = " " 和 len(s) > len(t)的场景。
*/
