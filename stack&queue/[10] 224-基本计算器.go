package stack_queue

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
	var operands, operators []uint8
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' || s[i] == '(' || s[i] == ')' {
			continue
		}
		if s[i] == '+' || s[i] == '-' {
			operators = append(operators, s[i])
		} else {
			operands = append(operands, s[i])
		}
	}
	var i int
	var res int
	for _, operator := range operators {
		x, y := int(operands[i]-'0'), int(operands[i+1]-'0')
		if operator == '+' {
			res += x + y
		}
		if operator == '-' {
			res += x - y
		}
		i += 2
	}
	return res
}
