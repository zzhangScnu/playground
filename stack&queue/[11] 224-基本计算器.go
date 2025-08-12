package stack_queue

import "unicode"

// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
//
// 注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
//
// 示例 1：
//
// 输入：s = "1 + 1"
// 输出：2
//
// 示例 2：
//
// 输入：s = " 2-1 + 2 "
// 输出：3
//
// 示例 3：
//
// 输入：s = "(1+(4+5+2)-3)+(6+8)"
// 输出：23
//
// 提示：
//
// 1 <= s.length <= 3 * 10⁵
// s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
// s 表示一个有效的表达式
// '+' 不能用作一元运算(例如， "+1" 和 "+(2 + 3)" 无效)
// '-' 可以用作一元运算(即 "-1" 和 "-(2 + 3)" 是有效的)
// 输入中不存在两个连续的操作符
// 每个数字和运行的计算将适合于一个有符号的 32位 整数
func calculate(s string) int {
	cur, res, operator := 0, 0, 1
	var stack [][]int
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if unicode.IsDigit(rune(ch)) {
			cur = cur*10 + int(s[i]-'0')
		} else if ch == '+' || ch == '-' {
			res += operator * cur
			operator = 1
			if ch == '-' {
				operator = -1
			}
			cur = 0
		} else if ch == '(' {
			stack = append(stack, []int{res, operator})
			res = 0
			operator = 1
			cur = 0
		} else if ch == ')' {
			res += operator * cur
			pre := stack[len(stack)-1]
			preRes, preOperator := pre[0], pre[1]
			stack = stack[:len(stack)-1]
			res *= preOperator
			res += preRes
			cur = 0
		}
	}
	return res + operator*cur
}

/**
思路：
因为引入了括号，需要额外处理优先级的逻辑。实际上涉及到括号，都应该先考虑引入栈实现。
整体思路像是分治：
- 计算括号外子问题的结果；
- 遇到左括号时，将子问题结果和子问题后的运算符一起入栈暂存，同时将之前的变量都清空，便于给括号内的新的子问题使用；
- 计算括号内子问题的结果；
- 遇到右括号时，意味着括号中高优子问题计算完成，此时应出栈，与先前的子问题结果和运算符进行进一步运算；
- 以此类推，直至遍历完成，获得最终结果。
*/

/**
注意：
- 因为使用切片实现栈，所以将operator定义为int类型，便于跟操作数一起入栈；
- 同时因为operator定义为int类型，加法和减法都可以统一表达为：res += operator * cur；
- 因为本解法是遇到操作符时进行运算，所以遇到右括号、到达表达式末尾时，实际上计算是未完成的，需要额外处理：res += operator * cur；
- 遇到右括号时，如原始表达式为 1 - (2 + 3)，
  此时栈内数据为[1, -1]，即preRes == 1, preOperator == -1，
  括号内res = 5，
  此时令res = res * preOperator = -5，
  再计算res = preRes * res = 1 - 5 = -4。
*/
