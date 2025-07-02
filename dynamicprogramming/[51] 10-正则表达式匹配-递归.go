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

// todo：那么考虑⼀下，如果加⼊ * 通配符，局⾯就会稍微复杂⼀些，不过只要分情况来分析，也不难理解。
//当 p[j + 1] 为 * 通配符时，我们分情况讨论下：
//1、如果 s[i] == p[j]，那么有两种情况：
//1.1 p[j] 有可能会匹配多个字符，⽐如 s = "aaa", p = "a*"，那么 p[0] 会通过 * 匹配 3 个字符 "a"。
//1.2 p[i] 也有可能匹配 0 个字符，⽐如 s = "aa", p = "a*aa"，由于后⾯的字符可以匹配 s，所以
//p[0] 只能匹配 0 次。
//2、如果 s[i] != p[j]，只有⼀种情况：
//p[j] 只能匹配 0 次，然后看下⼀个字符是否能和 s[i] 匹配。⽐如说 s = "aa", p = "b*aa"，此时
//p[0] 只能匹配 0 次。

// 调整对*的判断先后也可以通过