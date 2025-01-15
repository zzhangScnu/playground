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
