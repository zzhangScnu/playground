package dynamicprogramming

// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
//
// 所谓匹配，是要涵盖 整个 字符串 s 的，而不是部分字符串。
//
// 示例 1：
//
// 输入：s = "aa", p = "a"
// 输出：false
// 解释："a" 无法匹配 "aa" 整个字符串。
//
// 示例 2:
//
// 输入：s = "aa", p = "a*"
// 输出：true
// 解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
//
// 示例 3：
//
// 输入：s = "ab", p = ".*"
// 输出：true
// 解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
//
// 提示：
//
// 1 <= s.length <= 20
// 1 <= p.length <= 20
// s 只包含从 a-z 的小写字母。
// p 只包含从 a-z 的小写字母，以及字符 . 和 *。
// 保证每次出现字符 * 时，前面都匹配到有效的字符
func isMatchII(s string, p string) bool {
	m, n := len(s), len(p)
	memo := make(map[[2]int]bool)
	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		if j == n {
			return i == m
		}
		if flag, ok := memo[[2]int{i, j}]; ok {
			return flag
		}
		var flag bool
		firstMatchFlag := i < m && (s[i] == p[j] || p[j] == '.')
		if j+1 < n && p[j+1] == '*' {
			flag = dp(i, j+2) ||
				firstMatchFlag && dp(i+1, j)
		} else {
			flag = firstMatchFlag && dp(i+1, j+1)
		}
		memo[[2]int{i, j}] = flag
		return flag
	}
	return dp(0, 0)
}

/**
迭代：基于先前的计算结果递推当前结果；
递归：基于当前的计算结果+递归计算后续结果得出整体解。


与迭代不同，递归解法存在大量重叠子问题引发的重复计算，故需引入备忘录机制消除冗余。

为什么要使用j是否到达p的终点作为base case？
因为当j == len(p)时，意味着模式串已经遍历完成，如果能够完全匹配字符串s，则s的指针i也应到达终点。
而如果使用i是否到达s的终点作为base case，则会有两种合法的情况：
1. p也走完全程；
2. p仍有剩余模式子串，形如X*，此时可以在s中匹配X出现0次。
所以需要额外判断第二种条件是否满足。
此时需要这样写base case：
if i == m {
	// 处理模式剩余部分可匹配空字符串的情况
	for ; j+1 < n; j += 2 {
		if p[j+1] != '*' {
			return false
		}
	}
	return j == n
}

Go中map支持切片作为key。
定义：make(map[[2]int]bool)
写入：memo[[2]int{i, j}] = flag
读取：memo[[2]int{i, j}]
*/
