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

func IsValid(s string) bool {
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

func isValid(s string) bool {
	mapping := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []rune
	for _, ch := range []rune(s) {
		if targetLeft, ok := mapping[ch]; ok { // 如果是右括号(即可以通过 ch 在 mapping 中检索到左括号)，则与栈顶比较看左括号是否一致
			if len(stack) == 0 || stack[len(stack)-1] != targetLeft { // 不一致，意味着左右括号不匹配
				return false
			}
			stack = stack[:len(stack)-1] // 一致，意味着配对成功，消消乐
		} else { // 如果是左括号，直接入栈
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}

// 若栈空时直接入栈所有字符（包括右括号），而非先判断字符类型 —— 虽然部分用例结果正确，但逻辑不够严谨，且没有提前终止无效场景（比如输入 ")" 时，本可直接返回 false，却要等到遍历结束）。

/*
*
在遇到左括号时，将对应的右括号入栈；
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
