package stack_queue

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
// 示例 1：
//
// 输入：s = "()"
//
// 输出：true
//
// 示例 2：
//
// 输入：s = "()[]{}"
//
// 输出：true
//
// 示例 3：
//
// 输入：s = "(]"
//
// 输出：false
//
// 示例 4：
//
// 输入：s = "([])"
//
// 输出：true
//
// 提示：
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
func isValid(s string) bool {
	stack := Stack{}
	mapping := map[int]int{
		')': '(',
		']': '[',
		'}': '{',
	}
	var rightCnt int
	for _, ch := range s {
		if left, ok := mapping[int(ch)]; ok {
			if !stack.IsEmpty() && stack.Peek() == left {
				stack.Pop()
			} else {
				rightCnt++
			}
		} else {
			stack.Push(int(ch))
		}
	}
	return rightCnt == 0 && stack.IsEmpty()
}

/*
*
在遇到左括号时，将对应的有括号入栈；
那么在遇到右括号时，直接跟栈顶比较是否相等即可。
*/
func isValidII(s string) bool {
	stack := Stack{}
	for _, ch := range s {
		if ch == '(' {
			stack.Push(')')
		} else if ch == '[' {
			stack.Push(']')
		} else if ch == '{' {
			stack.Push('}')
		} else if stack.IsEmpty() || int32(stack.Peek()) != ch {
			return false
		} else {
			stack.Pop()
		}
	}
	return stack.IsEmpty()
}

/**
如果只有一种类型的括号，维护一个count，从左到右遍历。
当遇到左括号时，count++；
当遇到右括号时，count--。
遍历完成后，判断count == 0，如果满足，表示左右括号已经匹配完全。

但如果有多种括号，就无法通过简单计算匹配次数来实现，因为不仅需要维护多种括号对应的次数，还无法识别 [(]) 这种非法情况。
所以增加存储的信息量，通过先进后出的栈来实现。
*/
