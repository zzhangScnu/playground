package greedy

// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
// 示例 1：
//
// 输入：s = "(()"
// 输出：2
// 解释：最长有效括号子串是 "()"
//
// 示例 2：
//
// 输入：s = ")()())"
// 输出：4
// 解释：最长有效括号子串是 "()()"
//
// 示例 3：
//
// 输入：s = ""
// 输出：0
//
// 提示：
//
// 0 <= s.length <= 3 * 10⁴
// s[i] 为 '(' 或 ')'

func longestValidParentheses(s string) int {
	longest, left, right := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			right++
		}
		if left == right {
			longest = max(longest, left+right)
		} else if left < right {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			right++
		}
		if left == right {
			longest = max(longest, left+right)
		} else if right < left {
			left, right = 0, 0
		}
	}
	return longest
}

/**
思路：
从左到右遍历字符串，计算左括号和右括号的数量left和right。

如果left == right，表示遍历到目前为止，左右括号都是匹配的合法串，此时更新最大长度；

否则如果left < right，由于合法括号串要求顺序性和连续性，即先左后右。
如果左括号数量已经不足了，后续即使在右括号之后再遇到新的左括号，也无法弥补缺失，组成新的一对括号。
此时重置left和right，重新从下一个字符开始计算。

这样遍历下来，会漏掉一种情况：即左括号的数量一直比右括号多，但实际上是有合法括号对的，如 (()。

所以需要我们从右到左再次遍历，重新计算，唯一不同的地方是重置的判断条件变为right < left。
因为逆向遍历时，合法的括号是先右再左的。
*/
