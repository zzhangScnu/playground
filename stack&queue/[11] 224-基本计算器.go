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
