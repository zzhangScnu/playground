package stack_queue

import "unicode"

/*
*
Implement a basic calculator to evaluate a simple expression string
The expression string contains only non-negative integers, '+', '-', '*', '\' operators, and
open '(' and closing parentheses ')'. The integer division should truncate toward zero
You may assume that the given expression is always valid. AAll intermediate
results will be in the range of [-231, 231 - 1].
Note: You are not allowed to use any built-in function which evaluates strings
as mathematical expressions, such as eval().
Input: s = "2*(5+5*2)/3+(6/2+8)"
Output: 21
*/
func calculateIII(s string) int {
	pre, cur, res, operator := 0, 0, 0, 1
	var stack []int
	for i := 0; i < len(s); {
		ch := s[i]
		if unicode.IsDigit(rune(ch)) {
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				cur = cur*10 + int(s[i]-'0')
				i++
			}
			i--
			pre, cur, res = doCalculate(pre, cur, res, operator)
		} else {
			if ch != ' ' {
				switch ch {
				case '-':
					operator = -1
				case '+':
					operator = 1
				case '*':
					operator = 2
				case '/':
					operator = 3
				case '(':
					stack = append(stack, pre, res, operator)
					pre, cur, res, operator = 0, 0, 0, 1
				case ')':
					prevPre, prevRes, prevOperator := stack[len(stack)-3], stack[len(stack)-2], stack[len(stack)-1]
					stack = stack[:len(stack)-3]
					cur = res
					pre, res, operator = prevPre, prevRes, prevOperator
					pre, cur, res = doCalculate(pre, cur, res, operator)
				}
			}
			i++
		}
	}
	return res
}

func doCalculate(pre, cur, res, operator int) (int, int, int) {
	switch operator {
	case 1:
		res += cur
		pre = cur
		cur = 0
	case -1:
		res -= cur
		pre = -cur
		cur = 0
	case 2:
		res = res - pre + pre*cur
		pre = pre * cur
		cur = 0
	case 3:
		res = res - pre + pre/cur
		pre = pre / cur
		cur = 0
	}
	return pre, cur, res
}

/**
思路：
结合前两种计算器。

注意事项：
- 因为在括号优先级运算之外，还需处理乘除。所以在入栈暂存先前的子问题结果时，也需将pre入栈。
  为了正确恢复括号外的计算状态，确保括号内的结果能和括号外的前序运算正确衔接。
	- pre：记录"待参与乘除运算的中间结果"；
	- 括号：打断运算流程，需要pre保存现场，后续才能恢复；
	- 栈：保存完整上下文。因为括号内的计算结束后，需要将结果作为一个"整体数值"，继续参与括号外的运算。
*/
