package stack_queue

// 只有满足下面几点之一，括号字符串才是有效的：
//
// 它是一个空字符串，或者
// 它可以被写成 AB （A 与 B 连接）, 其中 A 和 B 都是有效字符串，或者
// 它可以被写作 (A)，其中 A 是有效字符串。
//
// 给定一个括号字符串 s ，在每一次操作中，你都可以在字符串的任何位置插入一个括号
//
// 例如，如果 s = "()))" ，你可以插入一个开始括号为 "(()))" 或结束括号为 "())))" 。
//
// 返回 为使结果字符串 s 有效而必须添加的最少括号数。
//
// 示例 1：
//
// 输入：s = "())"
// 输出：1
//
// 示例 2：
//
// 输入：s = "((("
// 输出：3
//
// 提示：
//
// 1 <= s.length <= 1000
// s 只包含 '(' 和 ')' 字符。
func minAddToMakeValid(s string) int {
	var left, right int
	for _, ch := range s {
		if ch == '(' {
			right++
		}
		if ch == ')' {
			if right > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return left + right
}

/**
思路：
left、right分别为需补充的左括号、右括号的需求数量。
从左到右遍历字符串。
- 当遇到左括号时，right++，即需要多一个右括号与其组成一对；
- 当遇到右括号时，如果此时right > 0，表示可以抵消掉一个右括号的需求；否则，需要补充一个左括号与其组成一对。
*/

/**
另一种做法：通过栈来维护括号，判断有效性。
func minAddToMakeValid(s string) int {
	stack := Stack{}
	mapping := map[int]int{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, ch := range s {
		ch := int(ch)
		left, ok := mapping[ch]
		if !ok {
			stack.Push(ch)
		} else {
			if !stack.IsEmpty() && stack.Peek() == left {
				stack.Pop()
			}
		}
	}
	return stack.Size()
}
*/
